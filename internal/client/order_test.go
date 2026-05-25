package client

import "testing"

func TestEncodeDecodeOrderedComment(t *testing.T) {
	cases := []struct {
		set, user string
		pos       int64
	}{
		{"default", "drop invalid", 100},
		{"input", "allow LAN", 250},
		{"default", "", 0},
		{"my-set", "with multiple words", 9999},
	}
	for _, c := range cases {
		encoded := EncodeOrderedComment(c.set, c.pos, c.user)
		gotSet, gotPos, gotUser, has := DecodeOrderedComment(encoded)
		if !has {
			t.Errorf("Decode(%q) returned has=false", encoded)
			continue
		}
		if gotSet != c.set || gotPos != c.pos || gotUser != c.user {
			t.Errorf("round trip mismatch: set=%q→%q pos=%d→%d user=%q→%q (encoded=%q)",
				c.set, gotSet, c.pos, gotPos, c.user, gotUser, encoded)
		}
	}
}

func TestDecodeNoMarker(t *testing.T) {
	for _, c := range []string{"", "regular user comment", "[tf:pos=10] missing set"} {
		_, _, user, has := DecodeOrderedComment(c)
		if has && c != "[tf:pos=10] missing set" {
			t.Errorf("Decode(%q) found marker where none exists", c)
		}
		// pos-only marker form is also accepted per regex
		if c == "" || c == "regular user comment" {
			if user != c {
				t.Errorf("user mismatch on %q: got %q", c, user)
			}
		}
	}
}
