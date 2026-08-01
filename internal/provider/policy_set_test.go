package provider

import (
	"context"
	"sort"
	"testing"
)

func TestDecodePolicySet(t *testing.T) {
	ctx := context.Background()
	// RouterOS returns granted permissions plus everything else negated, in its
	// own order. decodePolicySet must keep only the granted ones.
	wire := "read,test,winbox,!local,!telnet,!ssh,!write,!policy"
	set := decodePolicySet(ctx, wire)
	var got []string
	if d := set.ElementsAs(ctx, &got, false); d.HasError() {
		t.Fatalf("ElementsAs: %v", d)
	}
	sort.Strings(got)
	want := []string{"read", "test", "winbox"}
	if len(got) != len(want) {
		t.Fatalf("got %v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v want %v", got, want)
		}
	}
	// Empty / all-negated -> empty set, never null-panics.
	if s := decodePolicySet(ctx, "!ftp,!telnet"); s.IsNull() {
		t.Error("all-negated should yield an empty (non-null) set")
	}
}
