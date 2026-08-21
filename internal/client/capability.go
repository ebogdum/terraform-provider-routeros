package client

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// UnsupportedArgs reports property names a menu does not accept on this device.
type UnsupportedArgs struct {
	Menu     string
	Verb     string
	Rejected []string
	Accepted []string
	Board    string
	Version  string
}

func (e *UnsupportedArgs) Error() string {
	attrs := make([]string, len(e.Rejected))
	for i, r := range e.Rejected {
		attrs[i] = fmt.Sprintf("%s (%q)", strings.ReplaceAll(r, "-", "_"), r)
	}
	where := "this device"
	if e.Board != "" {
		where = e.Board
		if e.Version != "" {
			where += ", RouterOS " + e.Version
		}
	}
	return fmt.Sprintf("%s does not accept %s on %s. The menu accepts: %s",
		where, strings.Join(attrs, ", "), e.Menu, strings.Join(e.Accepted, ", "))
}

// MenuArgs lists the property names menuPath accepts for verb ("set"/"add").
// nil means the device could not be asked, not that it accepts nothing.
func (c *Client) MenuArgs(ctx context.Context, menuPath, verb string) map[string]bool {
	key := menuPath + "|" + verb
	c.capMu.Lock()
	if c.capCache == nil {
		c.capCache = map[string]map[string]bool{}
	}
	got, ok := c.capCache[key]
	c.capMu.Unlock()
	if ok {
		return got
	}

	args := c.fetchMenuArgs(ctx, menuPath, verb)

	c.capMu.Lock()
	c.capCache[key] = args
	c.capMu.Unlock()
	return args
}

func (c *Client) fetchMenuArgs(ctx context.Context, menuPath, verb string) map[string]bool {
	consolePath := strings.Join(append(
		strings.Split(strings.Trim(menuPath, "/"), "/"), verb), ",")
	raw, err := c.request(ctx, http.MethodPost, c.rel("console/inspect").String(),
		Object{"request": "child", "path": consolePath, "as-value": "yes"})
	if err != nil {
		return nil
	}
	rows, err := decodeList(raw)
	if err != nil {
		return nil
	}
	args := map[string]bool{}
	for _, r := range rows {
		if r["node-type"] == "arg" && r["type"] == "child" && r["name"] != "numbers" {
			args[r["name"]] = true
		}
	}
	if len(args) == 0 {
		return nil
	}
	return args
}

func (c *Client) identity(ctx context.Context) (board, version string) {
	c.capMu.Lock()
	done, board, version := c.idDone, c.idBoard, c.idVersion
	c.capMu.Unlock()
	if done {
		return board, version
	}
	obj, err := c.SystemResource(ctx)
	c.capMu.Lock()
	defer c.capMu.Unlock()
	c.idDone = true
	if err == nil {
		c.idBoard, c.idVersion = obj["board-name"], obj["version"]
	}
	return c.idBoard, c.idVersion
}

// CheckWritable reports property names in body that menuPath does not accept.
// A router that does not expose /console/inspect behaves exactly as before.
func (c *Client) CheckWritable(ctx context.Context, menuPath, verb string, body Object) error {
	if len(body) == 0 {
		return nil
	}
	args := c.MenuArgs(ctx, menuPath, verb)
	if args == nil {
		return nil
	}
	var rejected []string
	for k := range body {
		if !args[k] {
			rejected = append(rejected, k)
		}
	}
	if len(rejected) == 0 {
		return nil
	}
	accepted := make([]string, 0, len(args))
	for a := range args {
		accepted = append(accepted, a)
	}
	sort.Strings(rejected)
	sort.Strings(accepted)
	board, version := c.identity(ctx)
	return &UnsupportedArgs{
		Menu: menuPath, Verb: verb,
		Rejected: rejected, Accepted: accepted,
		Board: board, Version: version,
	}
}
