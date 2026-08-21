package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// orderFake serves a fixed chain and records every move it is asked to perform.
type orderFake struct {
	rows  []Object
	moves []Object
}

func (f *orderFake) client(t *testing.T) (*Client, func()) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/move") {
			var body Object
			_ = json.NewDecoder(r.Body).Decode(&body)
			f.moves = append(f.moves, body)
			_, _ = w.Write([]byte(`[]`))
			return
		}
		_ = json.NewEncoder(w).Encode(f.rows)
	}))
	c, err := New(Config{Host: srv.URL, Username: "u", Password: "p"})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return c, srv.Close
}

func chain(ids ...string) []Object {
	rows := make([]Object, len(ids))
	for i, id := range ids {
		rows[i] = Object{".id": id}
	}
	return rows
}

func TestPlaceOrderedNeverSendsARulePastUnmanagedRules(t *testing.T) {
	f := &orderFake{rows: chain("*1", "*2", "*9")} // *9 unmanaged, e.g. a defconf drop
	c, done := f.client(t)
	defer done()

	snap := map[string]int64{"*1": 20, "*2": 10}
	if err := c.PlaceOrdered(context.Background(), "", "/ip/firewall/filter", "*1", 20, snap); err != nil {
		t.Fatalf("PlaceOrdered: %v", err)
	}
	if len(f.moves) != 1 {
		t.Fatalf("moves = %v, want exactly one", f.moves)
	}
	if got := f.moves[0]["destination"]; got != "*9" {
		t.Errorf("destination = %q, want %q (the row after the last managed peer)", got, "*9")
	}
}

func TestPlaceOrderedMovesBeforeTheNextManagedRule(t *testing.T) {
	f := &orderFake{rows: chain("*1", "*2")}
	c, done := f.client(t)
	defer done()

	snap := map[string]int64{"*1": 20, "*2": 10}
	if err := c.PlaceOrdered(context.Background(), "", "/ip/firewall/filter", "*2", 10, snap); err != nil {
		t.Fatalf("PlaceOrdered: %v", err)
	}
	if len(f.moves) != 1 || f.moves[0]["destination"] != "*1" {
		t.Fatalf("moves = %v, want one move before *1", f.moves)
	}
}

func TestPlaceOrderedSkipsWhenAlreadyOrdered(t *testing.T) {
	f := &orderFake{rows: chain("*1", "*2")}
	c, done := f.client(t)
	defer done()

	snap := map[string]int64{"*1": 10, "*2": 20}
	if err := c.PlaceOrdered(context.Background(), "", "/ip/firewall/filter", "*1", 10, snap); err != nil {
		t.Fatalf("PlaceOrdered: %v", err)
	}
	if len(f.moves) != 0 {
		t.Errorf("moves = %v, want none", f.moves)
	}
}

func TestPlaceOrderedIsANoOpForASingleManagedRule(t *testing.T) {
	f := &orderFake{rows: chain("*1", "*9")}
	c, done := f.client(t)
	defer done()

	if err := c.PlaceOrdered(context.Background(), "", "/ip/firewall/filter", "*1", 200, map[string]int64{"*1": 200}); err != nil {
		t.Fatalf("PlaceOrdered: %v", err)
	}
	if len(f.moves) != 0 {
		t.Errorf("moves = %v, want none", f.moves)
	}
}

func TestMoveRefusesAnEmptyDestination(t *testing.T) {
	f := &orderFake{rows: chain("*1")}
	c, done := f.client(t)
	defer done()

	if err := c.Move(context.Background(), "/ip/firewall/filter", "*1", ""); err == nil {
		t.Fatal("Move accepted an empty destination")
	}
	if len(f.moves) != 0 {
		t.Errorf("Move issued a request anyway: %v", f.moves)
	}
	if err := c.MoveToEnd(context.Background(), "/ip/firewall/filter", "*1"); err != nil {
		t.Fatalf("MoveToEnd: %v", err)
	}
	if _, ok := f.moves[0]["destination"]; len(f.moves) != 1 || ok {
		t.Errorf("MoveToEnd sent %v, want numbers only", f.moves)
	}
}
