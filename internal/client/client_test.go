package client

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestServer(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	c, err := New(Config{Host: srv.URL, Username: "u", Password: "p"})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c
}

func TestListAddSetRemove(t *testing.T) {
	store := map[string]Object{
		"*1": {".id": "*1", "address": "10.0.0.1/24", "interface": "ether1", "disabled": "false"},
	}
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/rest/")
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			if path == "ip/address" {
				out := []Object{}
				for _, v := range store {
					out = append(out, v)
				}
				json.NewEncoder(w).Encode(out)
				return
			}
			id := strings.TrimPrefix(path, "ip/address/")
			if obj, ok := store[id]; ok {
				json.NewEncoder(w).Encode(obj)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIError{Code: 404, Message: "Not Found"})
		case http.MethodPut:
			body, _ := io.ReadAll(r.Body)
			var in Object
			json.Unmarshal(body, &in)
			in[".id"] = "*2"
			store["*2"] = in
			json.NewEncoder(w).Encode(in)
		case http.MethodPatch:
			id := strings.TrimPrefix(path, "ip/address/")
			body, _ := io.ReadAll(r.Body)
			var in Object
			json.Unmarshal(body, &in)
			obj := store[id]
			for k, v := range in {
				obj[k] = v
			}
			store[id] = obj
			json.NewEncoder(w).Encode(obj)
		case http.MethodDelete:
			id := strings.TrimPrefix(path, "ip/address/")
			if _, ok := store[id]; !ok {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(APIError{Code: 404, Message: "Not Found"})
				return
			}
			delete(store, id)
			w.WriteHeader(http.StatusOK)
		}
	})
	ctx := context.Background()

	list, err := c.List(ctx, "/ip/address")
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("want 1 record, got %d", len(list))
	}

	obj, err := c.Add(ctx, "/ip/address", Object{"address": "10.0.0.2/24", "interface": "ether2"})
	if err != nil {
		t.Fatalf("Add: %v", err)
	}
	if obj[".id"] != "*2" {
		t.Fatalf("want .id=*2, got %q", obj[".id"])
	}

	got, err := c.GetByID(ctx, "/ip/address", "*2")
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if got["address"] != "10.0.0.2/24" {
		t.Fatalf("read mismatch: %+v", got)
	}

	upd, err := c.Set(ctx, "/ip/address", "*2", Object{"comment": "updated"})
	if err != nil {
		t.Fatalf("Set: %v", err)
	}
	if upd["comment"] != "updated" {
		t.Fatalf("set mismatch: %+v", upd)
	}

	if err := c.Remove(ctx, "/ip/address", "*2"); err != nil {
		t.Fatalf("Remove: %v", err)
	}
	// Remove on missing record is idempotent (returns nil).
	if err := c.Remove(ctx, "/ip/address", "*2"); err != nil {
		t.Fatalf("Remove(missing): %v", err)
	}
	if _, err := c.GetByID(ctx, "/ip/address", "*2"); !IsNotFound(err) {
		t.Fatalf("want IsNotFound, got %v", err)
	}
}

func TestAPIError(t *testing.T) {
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(APIError{Code: 406, Message: "Not Acceptable", Detail: "no such argument"})
	})
	_, err := c.Add(context.Background(), "/ip/address", Object{"bogus": "x"})
	if err == nil {
		t.Fatal("want error")
	}
	if !strings.Contains(err.Error(), "Not Acceptable") || !strings.Contains(err.Error(), "no such argument") {
		t.Fatalf("error format wrong: %v", err)
	}
}
