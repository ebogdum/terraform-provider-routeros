package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// capFake serves /console/inspect for one menu and counts how often it is asked.
type capFake struct {
	args      []string
	inspects  int
	writes    int
	noInspect bool
}

func (f *capFake) client(t *testing.T) *Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/console/inspect") {
			f.inspects++
			if f.noInspect {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(`{"error":404,"message":"no such command"}`))
				return
			}
			rows := []Object{{"name": "set", "node-type": "cmd", "type": "self"}}
			for _, a := range f.args {
				rows = append(rows, Object{"name": a, "node-type": "arg", "type": "child"})
			}
			_ = json.NewEncoder(w).Encode(rows)
			return
		}
		if strings.HasSuffix(r.URL.Path, "system/resource") {
			_ = json.NewEncoder(w).Encode(Object{"board-name": "TestBoard", "version": "7.23.2"})
			return
		}
		f.writes++
		_ = json.NewEncoder(w).Encode(Object{".id": "*1"})
	}))
	t.Cleanup(srv.Close)
	c, err := New(Config{Host: srv.URL, Username: "u", Password: "p"})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c
}

func TestAddRefusesAnArgumentTheMenuDoesNotAccept(t *testing.T) {
	f := &capFake{args: []string{"cpu-flow-control", "name", "switch-all-ports"}}
	c := f.client(t)

	_, err := c.Add(context.Background(), "/interface/ethernet/switch",
		Object{"name": "switch1", "l3-hw-offloading": "yes"})
	if err == nil {
		t.Fatal("Add accepted an argument the menu does not have")
	}
	ua, ok := err.(*UnsupportedArgs)
	if !ok {
		t.Fatalf("err is %T, want *UnsupportedArgs", err)
	}
	if len(ua.Rejected) != 1 || ua.Rejected[0] != "l3-hw-offloading" {
		t.Errorf("rejected = %v, want [l3-hw-offloading]", ua.Rejected)
	}
	for _, want := range []string{"TestBoard", "7.23.2", "l3_hw_offloading",
		"l3-hw-offloading", "cpu-flow-control"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("message is missing %q: %s", want, err)
		}
	}
	if f.writes != 0 {
		t.Errorf("the write went out anyway (%d)", f.writes)
	}
}

func TestWritesTheMenuAcceptsGoThrough(t *testing.T) {
	f := &capFake{args: []string{"name", "comment"}}
	c := f.client(t)
	if _, err := c.Add(context.Background(), "/x", Object{"name": "a", "comment": "b"}); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if _, err := c.Set(context.Background(), "/x", "*1", Object{"comment": "c"}); err != nil {
		t.Fatalf("Set: %v", err)
	}
	if f.writes != 2 {
		t.Errorf("writes = %d, want 2", f.writes)
	}
}

// A router that does not expose /console/inspect, or an account that cannot
// reach it, must behave exactly as it did before the check existed.
func TestAWriteStillWorksWhenTheDeviceCannotBeAsked(t *testing.T) {
	f := &capFake{noInspect: true}
	c := f.client(t)
	if _, err := c.Add(context.Background(), "/x", Object{"anything": "at-all"}); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if f.writes != 1 {
		t.Errorf("writes = %d, want 1", f.writes)
	}
}

func TestMenuArgsAreFetchedOncePerMenu(t *testing.T) {
	f := &capFake{args: []string{"name"}}
	c := f.client(t)
	for i := 0; i < 5; i++ {
		if _, err := c.Add(context.Background(), "/x", Object{"name": "a"}); err != nil {
			t.Fatalf("Add: %v", err)
		}
	}
	if f.inspects != 1 {
		t.Errorf("inspects = %d, want 1 -- the cache is not holding", f.inspects)
	}
}

func TestAnEmptyBodyIsNeverChecked(t *testing.T) {
	f := &capFake{args: []string{"name"}}
	c := f.client(t)
	if err := c.CheckWritable(context.Background(), "/x", "set", Object{}); err != nil {
		t.Fatalf("CheckWritable on an empty body: %v", err)
	}
	if f.inspects != 0 {
		t.Errorf("inspects = %d, want 0", f.inspects)
	}
}
