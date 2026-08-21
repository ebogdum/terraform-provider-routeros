package client

import "testing"

// RouterOS reads "+" in a query value literally, so a form-encoded space matches
// nothing. Values confirmed against ROS 7.23.2 via /ip/firewall/address-list.
func TestQueryEncodeUsesPercentTwentyForSpaces(t *testing.T) {
	for _, tc := range []struct {
		name, key, value, want string
	}{
		{"space in value", "comment", "drop all not from LAN", "comment=drop%20all%20not%20from%20LAN"},
		{"literal plus stays escaped", "comment", "a+b", "comment=a%2Bb"},
		{"plus and space together", "comment", "a b+c", "comment=a%20b%2Bc"},
		{"space in key", "we ird", "v", "we%20ird=v"},
		{"ampersand and equals", "comment", "a&b=c", "comment=a%26b%3Dc"},
		{"non-ascii", "comment", "café au lait", "comment=caf%C3%A9%20au%20lait"},
		{"no special characters", "chain", "input", "chain=input"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := NewQuery(WithFilter(tc.key, tc.value)).Encode()
			if got != tc.want {
				t.Errorf("Encode() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestQueryEncodeProplist(t *testing.T) {
	got := NewQuery(WithProplist("name", "comment")).Encode()
	if want := ".proplist=name%2Ccomment"; got != want {
		t.Errorf("Encode() = %q, want %q", got, want)
	}
}

func TestQueryEncodeNilAndEmpty(t *testing.T) {
	var q *Query
	if got := q.Encode(); got != "" {
		t.Errorf("nil Query.Encode() = %q, want empty", got)
	}
	if got := NewQuery().Encode(); got != "" {
		t.Errorf("empty Query.Encode() = %q, want empty", got)
	}
}
