package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// InspectRequest mirrors POST /rest/console/inspect, RouterOS's CLI completion
// engine. It is the closest thing to a schema introspection endpoint and is
// used by the audit tool to enumerate menus and per-argument metadata.
//
// Common shapes:
//
//	{"path":"ip address","request":"child"}      → list child menus / items
//	{"path":"ip address","request":"argument"}   → list arguments and their types
//	{"path":"ip address","request":"command"}    → list commands (add/set/print/remove/move/...)
type InspectRequest struct {
	Path    string `json:"path"`
	Request string `json:"request"` // "child" | "argument" | "command" | "completion"
	Word    string `json:"word,omitempty"`
}

// InspectItem is one row from an inspect response. Keys vary by request type.
type InspectItem map[string]string

// Inspect calls /rest/console/inspect.
func (c *Client) Inspect(ctx context.Context, r InspectRequest) ([]InspectItem, error) {
	r.Path = strings.TrimSpace(strings.Trim(r.Path, "/"))
	u := c.rel("console/inspect")
	raw, err := c.request(ctx, http.MethodPost, u.String(), r)
	if err != nil {
		return nil, err
	}
	var items []InspectItem
	if err := json.Unmarshal(raw, &items); err != nil {
		// Some firmwares return a single object.
		var one InspectItem
		if jerr := json.Unmarshal(raw, &one); jerr != nil {
			return nil, fmt.Errorf("routeros: inspect: %w", err)
		}
		return []InspectItem{one}, nil
	}
	return items, nil
}
