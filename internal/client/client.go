// Package client implements a typed REST client for the MikroTik RouterOS 7.x
// /rest API. It is consumed by the Terraform provider runtime and by the
// schema audit tool.
//
// Method mapping (per help.mikrotik.com/.../REST+API):
//
//	GET    /rest/<path>           print (list or by id)
//	PUT    /rest/<path>           add (returns full new object with .id)
//	PATCH  /rest/<path>/<id>      set (returns full updated object)
//	DELETE /rest/<path>/<id>      remove
//	POST   /rest/<path>/<cmd>     arbitrary command (print, monitor once, move, action)
//
// All response field values come back as JSON strings; typed coercion lives in
// coerce.go.
package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// Config carries provider-block options. Empty values are treated as unset.
type Config struct {
	Host       string // e.g. "https://192.0.2.1" or "192.0.2.1" (https assumed)
	Username   string
	Password   string
	CACertPEM  string
	Insecure   bool
	ROSVersion string

	// HTTPTimeout is the per-request timeout. Defaults to 90s.
	// RouterOS itself terminates requests at 60s, so 90s leaves headroom for TCP.
	HTTPTimeout time.Duration

	// MaxRetries on transport / 5xx errors. Defaults to 3.
	MaxRetries int

	// RetryBackoff is the initial backoff; doubles per retry. Defaults to 500ms.
	RetryBackoff time.Duration
}

// Client is a RouterOS REST client. Safe for concurrent use.
type Client struct {
	cfg  Config
	base *url.URL // always ends in /rest/
	http *http.Client

	orderLocksMu sync.Mutex
	orderLocks   map[string]*sync.Mutex
}

// New constructs a Client. The router's reachability is not checked here.
func New(cfg Config) (*Client, error) {
	if cfg.Host == "" {
		return nil, errors.New("routeros: empty host")
	}
	if cfg.HTTPTimeout == 0 {
		cfg.HTTPTimeout = 90 * time.Second
	}
	if cfg.MaxRetries == 0 {
		cfg.MaxRetries = 3
	}
	if cfg.RetryBackoff == 0 {
		cfg.RetryBackoff = 500 * time.Millisecond
	}

	host := cfg.Host
	if !strings.Contains(host, "://") {
		host = "https://" + host
	}
	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("routeros: invalid host %q: %w", cfg.Host, err)
	}
	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}
	if !strings.HasSuffix(u.Path, "/rest/") {
		u.Path += "rest/"
	}

	tlsCfg := &tls.Config{
		// Skipping verification is a user-opt feature for self-signed RouterOS
		// certs. Documented and surfaced in the provider schema as `insecure`.
		InsecureSkipVerify: cfg.Insecure, //nolint:gosec // G402: user-opt
		MinVersion:         tls.VersionTLS12,
	}
	if cfg.CACertPEM != "" {
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(cfg.CACertPEM)) {
			return nil, errors.New("routeros: ca_cert: no certificates parsed from PEM")
		}
		tlsCfg.RootCAs = pool
	}

	// Cap header timeout at half of overall timeout (min 10s) so a slow server
	// can be detected before the overall ctx is exhausted.
	headerTimeout := min(max(cfg.HTTPTimeout/2, 10*time.Second), cfg.HTTPTimeout)

	tr := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		TLSClientConfig:       tlsCfg,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: headerTimeout,
		ExpectContinueTimeout: 1 * time.Second,
		IdleConnTimeout:       90 * time.Second,
		MaxIdleConns:          20,
		MaxIdleConnsPerHost:   8,
		ForceAttemptHTTP2:     false,
		// RouterOS REST has historical issues with HTTP/2. An empty
		// TLSNextProto map prevents the runtime from auto-configuring h2
		// even if ALPN would negotiate it.
		TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
	}

	return &Client{
		cfg:  cfg,
		base: u,
		http: &http.Client{
			Timeout:   cfg.HTTPTimeout,
			Transport: tr,
		},
	}, nil
}

// APIError is the structured error body RouterOS returns on 4xx/5xx.
type APIError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"error"`
	Message    string `json:"message"`
	Detail     string `json:"detail,omitempty"`
}

func (e *APIError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("routeros: %d %s: %s", e.Code, e.Message, e.Detail)
	}
	return fmt.Sprintf("routeros: %d %s", e.Code, e.Message)
}

// IsNotFound reports whether err is a 404 from RouterOS.
func IsNotFound(err error) bool {
	var ae *APIError
	if errors.As(err, &ae) {
		return ae.StatusCode == http.StatusNotFound || ae.Code == http.StatusNotFound
	}
	return false
}

// Object is a flat RouterOS record. Every value is a string in JSON, mirroring
// the wire format; typed access lives in coerce.go.
type Object map[string]string

// menuPath normalises a CLI-style menu path ("/ip/address", "ip/address",
// "/ip/address/") to the form REST expects ("ip/address").
func menuPath(p string) string {
	p = strings.TrimSpace(p)
	p = strings.Trim(p, "/")
	return p
}

// rel builds the URL for path[+suffix]. RouterOS REST requires literal `*` in
// .id values (e.g. /rest/interface/wireguard/*1) and rejects the URL-encoded
// form %2A. We let net/url compute the proper escaped form for everything else
// (spaces, '#', '?', non-ASCII), then substitute %2A back to '*' in RawPath.
func (c *Client) rel(path string, suffix ...string) *url.URL {
	u := *c.base
	parts := []string{menuPath(path)}
	for _, s := range suffix {
		s = strings.Trim(s, "/")
		if s != "" {
			parts = append(parts, s)
		}
	}
	joined := strings.Join(parts, "/")
	u.Path += joined
	// EscapedPath returns the canonical encoding of Path. Patch %2A -> '*' so
	// RouterOS accepts the .id segment; every other escape is preserved.
	u.RawPath = strings.ReplaceAll(u.EscapedPath(), "%2A", "*")
	return &u
}

// do issues req, retrying on transport errors and 5xx for safe-to-retry
// requests only. Returns the response status code, fully drained body, and
// any error. The underlying response body is always closed before return.
//
//	GET, DELETE                -- always retry-safe
//	POST <path>/set            -- RouterOS singleton upsert; idempotent (same
//	                             body produces the same final state). Retry-safe.
//	POST <path>/move           -- idempotent (Move(id, dest) twice = once).
//	                             Retry-safe.
//	PUT, other POST            -- NOT retry-safe (PUT=add creates duplicates;
//	                             POST <action> may run twice).
func (c *Client) do(ctx context.Context, req *http.Request) (int, []byte, error) {
	retrySafe := req.Method == http.MethodGet || req.Method == http.MethodDelete
	if req.Method == http.MethodPost {
		p := req.URL.Path
		if strings.HasSuffix(p, "/set") || strings.HasSuffix(p, "/move") {
			retrySafe = true
		}
	}
	backoff := c.cfg.RetryBackoff
	var lastErr error
	for attempt := 0; attempt <= c.cfg.MaxRetries; attempt++ {
		if attempt > 0 {
			t := time.NewTimer(backoff)
			select {
			case <-ctx.Done():
				t.Stop()
				return 0, nil, ctx.Err()
			case <-t.C:
			}
			backoff *= 2
		}
		// http.Request is single-shot; if there's a body we need to clone it.
		r := req.Clone(ctx)
		if req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return 0, nil, err
			}
			r.Body = body
		}
		resp, err := c.http.Do(r)
		if err != nil {
			// Caller-cancelled or deadline-exceeded: no point retrying;
			// the next iteration would just see the same ctx.Err().
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return 0, nil, err
			}
			lastErr = err
			if !retrySafe {
				return 0, nil, err
			}
			continue
		}
		body, readErr := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if readErr != nil {
			if errors.Is(readErr, context.Canceled) || errors.Is(readErr, context.DeadlineExceeded) {
				return 0, nil, readErr
			}
			lastErr = readErr
			if !retrySafe {
				return 0, nil, readErr
			}
			continue
		}
		if resp.StatusCode >= 500 && retrySafe {
			lastErr = fmt.Errorf("routeros: %s %s -> %d (will retry)", req.Method, req.URL.Path, resp.StatusCode)
			continue
		}
		return resp.StatusCode, body, nil
	}
	return 0, nil, lastErr
}

// request builds and executes; on non-2xx it returns *APIError.
func (c *Client) request(ctx context.Context, method, urlStr string, body any) ([]byte, error) {
	var bodyReader io.Reader
	var getBody func() (io.ReadCloser, error)
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("routeros: marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(buf)
		getBody = func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewReader(buf)), nil }
	}
	req, err := http.NewRequestWithContext(ctx, method, urlStr, bodyReader)
	if err != nil {
		return nil, err
	}
	req.GetBody = getBody
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.cfg.Username != "" {
		req.SetBasicAuth(c.cfg.Username, c.cfg.Password)
	}

	status, raw, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}
	if status >= 200 && status < 300 {
		return raw, nil
	}
	apiErr := &APIError{StatusCode: status}
	if jerr := json.Unmarshal(raw, apiErr); jerr != nil || apiErr.Message == "" {
		apiErr.Message = strings.TrimSpace(string(raw))
		if apiErr.Message == "" {
			apiErr.Message = http.StatusText(status)
		}
	}
	return nil, apiErr
}

// List fetches every record at menuPath. opts is optional.
func (c *Client) List(ctx context.Context, path string, opts ...QueryOption) ([]Object, error) {
	u := c.rel(path)
	q := NewQuery(opts...)
	if qs := q.Encode(); qs != "" {
		u.RawQuery = qs
	}
	raw, err := c.request(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return decodeList(raw)
}

// GetByID fetches a single record by .id ("*1"). Returns IsNotFound on 404
// or on 200-with-empty-body (some intermediaries swallow the body).
func (c *Client) GetByID(ctx context.Context, path, id string) (Object, error) {
	u := c.rel(path, id)
	raw, err := c.request(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	if len(bytes.TrimSpace(raw)) == 0 {
		return nil, &APIError{StatusCode: http.StatusNotFound, Code: http.StatusNotFound, Message: "empty response body"}
	}
	return decodeOne(raw)
}

// Add creates a record. Returns the full created object (including .id).
func (c *Client) Add(ctx context.Context, path string, body Object) (Object, error) {
	u := c.rel(path)
	raw, err := c.request(ctx, http.MethodPut, u.String(), body)
	if err != nil {
		return nil, err
	}
	return decodeOne(raw)
}

// Set patches fields on a record by id, returning the full updated object.
func (c *Client) Set(ctx context.Context, path, id string, body Object) (Object, error) {
	u := c.rel(path, id)
	raw, err := c.request(ctx, http.MethodPatch, u.String(), body)
	if err != nil {
		return nil, err
	}
	return decodeOne(raw)
}

// SetSingleton sets fields on a singleton menu (no id, e.g. /ip/dns,
// /system/identity, /caps-man/aaa). RouterOS REST rejects PATCH for these --
// "missing or invalid resource identifier" -- because PATCH requires an .id
// segment. The accepted form is POST /rest/<menu>/set with the same body.
// After the set we GET the menu to return the fresh state.
func (c *Client) SetSingleton(ctx context.Context, path string, body Object) (Object, error) {
	u := c.rel(path, "set")
	if _, err := c.request(ctx, http.MethodPost, u.String(), body); err != nil {
		return nil, err
	}
	return c.GetSingleton(ctx, path)
}

// GetSingleton fetches a singleton menu's current state via GET. RouterOS
// returns either a bare object or a one-element array depending on the menu.
func (c *Client) GetSingleton(ctx context.Context, path string) (Object, error) {
	u := c.rel(path)
	raw, err := c.request(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	return decodeOne(raw)
}

// Remove deletes a record. Returns nil on success and on already-missing (404).
func (c *Client) Remove(ctx context.Context, path, id string) error {
	u := c.rel(path, id)
	_, err := c.request(ctx, http.MethodDelete, u.String(), nil)
	if err != nil && IsNotFound(err) {
		return nil
	}
	return err
}

// Exec issues POST /rest/<path>/<command> with the given JSON body. Use this
// for arbitrary commands (move, monitor once, system reboot, tool fetch, etc).
// The result is decoded as either a list or a single object depending on the
// server's reply shape.
func (c *Client) Exec(ctx context.Context, path, command string, body any) ([]Object, error) {
	u := c.rel(path, command)
	raw, err := c.request(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
	}
	return decodeAny(raw)
}

// Move issues POST /rest/<path>/move with {numbers:<id>, destination:<dest>}.
// Used for ordered menus (firewall filter/nat/mangle, queue tree, etc.).
func (c *Client) Move(ctx context.Context, path, id, destination string) error {
	body := Object{"numbers": id}
	if destination != "" {
		body["destination"] = destination
	}
	_, err := c.Exec(ctx, path, "move", body)
	return err
}

// SystemResource fetches /system/resource (used to discover RouterOS version).
func (c *Client) SystemResource(ctx context.Context) (Object, error) {
	raw, err := c.request(ctx, http.MethodGet, c.rel("system/resource").String(), nil)
	if err != nil {
		return nil, err
	}
	// /system/resource is a singleton: server returns a bare object.
	return decodeOne(raw)
}

func decodeOne(raw []byte) (Object, error) {
	if len(bytes.TrimSpace(raw)) == 0 {
		return Object{}, nil
	}
	// Some endpoints return [obj], others {obj}.
	trim := bytes.TrimSpace(raw)
	if trim[0] == '[' {
		var arr []Object
		if err := json.Unmarshal(raw, &arr); err != nil {
			return nil, fmt.Errorf("routeros: decode response: %w", err)
		}
		if len(arr) == 0 {
			return Object{}, nil
		}
		return arr[0], nil
	}
	var obj Object
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil, fmt.Errorf("routeros: decode response: %w", err)
	}
	return obj, nil
}

func decodeList(raw []byte) ([]Object, error) {
	trim := bytes.TrimSpace(raw)
	if len(trim) == 0 {
		return nil, nil
	}
	if trim[0] == '[' {
		var arr []Object
		if err := json.Unmarshal(raw, &arr); err != nil {
			return nil, fmt.Errorf("routeros: decode list: %w", err)
		}
		return arr, nil
	}
	var obj Object
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil, fmt.Errorf("routeros: decode list (object fallback): %w", err)
	}
	return []Object{obj}, nil
}

func decodeAny(raw []byte) ([]Object, error) {
	return decodeList(raw)
}
