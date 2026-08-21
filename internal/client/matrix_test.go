package client

import "testing"

// The same concept has three spellings across three boards, all on RouterOS
// 7.23: a hAP ax^3 has both mirror-source and mirror-target, an RB5009 calls it
// mirror-egress-target, a CRS305 has only mirror-target.
func TestMirrorTargetIsNamedDifferentlyPerBoard(t *testing.T) {
	for _, tc := range []struct{ board, prop, wantWire, wantAbsent string }{
		{"hAP ax^3", "mirror-source", "mirror-source", ""},
		{"RB5009UPr+S+", "mirror-source", "mirror-egress-target", ""},
		{"CRS305-1G-4S+", "mirror-source", "mirror-source", "mirror-target"},
		{"hAP ax^3", "mirror-target", "mirror-target", ""},
		{"CRS305-1G-4S+", "mirror-target", "mirror-target", ""},
		{"RB5009UPr+S+", "mirror-target", "mirror-target", "mirror-egress-target"},
	} {
		wire, absent := resolveProperty("/interface/ethernet/switch", tc.prop, tc.board)
		if wire != tc.wantWire {
			t.Errorf("%s %s -> wire %q, want %q", tc.board, tc.prop, wire, tc.wantWire)
		}
		if (tc.wantAbsent == "") != (absent == "") {
			t.Errorf("%s %s -> absent %q, want absent=%v", tc.board, tc.prop, absent, tc.wantAbsent != "")
		}
	}
}

func TestL3HWOffloadingOnlyOnSwitchChipsThatHaveIt(t *testing.T) {
	for board, want := range map[string]bool{
		"CRS305-1G-4S+": false, "CRS326-24G-2S+": false,
		"RB5009UPr+S+": true, "hAP ax^3": true,
	} {
		absent := KnownAbsent("/interface/ethernet/switch", "l3-hw-offloading", board)
		if (absent != "") != want {
			t.Errorf("%s: absent=%q, want absent=%v", board, absent, want)
		}
	}
}

// A name on no board at all needs no board list to reject it.
func TestFasttrackHWIsOnNoBoard(t *testing.T) {
	for _, board := range []string{"hAP ax^3", "RB5009UPr+S+", "CRS305-1G-4S+", "Anything"} {
		if KnownAbsent("/interface/ethernet/switch", "fasttrack-hw", board) == "" {
			t.Errorf("%s: fasttrack-hw was not rejected", board)
		}
	}
}

func TestApplyMatrixRenamesAndRejects(t *testing.T) {
	body := Object{"mirror-source": "ether1", "name": "switch1"}
	out, absent := applyMatrix("/interface/ethernet/switch", "RB5009UPr+S+", body)
	if _, ok := out["mirror-egress-target"]; !ok {
		t.Errorf("mirror-source was not renamed: %v", out)
	}
	if out["name"] != "switch1" {
		t.Errorf("an unaffected property was disturbed: %v", out)
	}
	if len(absent) != 0 {
		t.Errorf("nothing should be absent here: %v", absent)
	}

	out, absent = applyMatrix("/interface/ethernet/switch", "CRS305-1G-4S+",
		Object{"cpu-flow-control": "yes", "name": "switch1"})
	if _, ok := absent["cpu-flow-control"]; !ok {
		t.Errorf("cpu-flow-control should be absent on a CRS305: %v", absent)
	}
	if _, ok := out["cpu-flow-control"]; ok {
		t.Error("an absent property was still put on the wire")
	}
}

// An unknown board must not be second-guessed: the matrix says nothing, so the
// property goes out as written and the device decides.
func TestAnUnknownBoardIsLeftAlone(t *testing.T) {
	out, absent := applyMatrix("/interface/ethernet/switch", "Some Future Board",
		Object{"mirror-source": "ether1"})
	if len(absent) != 0 {
		t.Errorf("absent = %v, want none", absent)
	}
	if _, ok := out["mirror-source"]; !ok {
		t.Errorf("out = %v, want mirror-source untouched", out)
	}
	if out, _ := applyMatrix("/some/unlisted/menu", "hAP ax^3", Object{"x": "1"}); out["x"] != "1" {
		t.Error("an unlisted menu was disturbed")
	}
	if out, _ := applyMatrix("/interface/ethernet/switch", "", Object{"mirror-source": "e"}); out["mirror-source"] != "e" {
		t.Error("an unknown board name was not left alone")
	}
}

// The matrix runs before the device is asked, so a board known not to have a
// property is refused with the reason rather than a bare "not accepted".
func TestMatrixRejectionExplainsItself(t *testing.T) {
	e := &UnsupportedArgs{
		Menu: "/interface/ethernet/switch", Board: "hAP ax^3", Version: "7.23.2",
		Rejected: []string{"fasttrack-hw"},
		Reasons:  []string{`fasttrack_hw ("fasttrack-hw") -- not a property of this menu on RouterOS 7.23`},
	}
	msg := e.Error()
	for _, want := range []string{"hAP ax^3", "7.23.2", "/interface/ethernet/switch",
		"fasttrack_hw", "not a property of this menu"} {
		if !contains(msg, want) {
			t.Errorf("message is missing %q: %s", want, msg)
		}
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	}())
}
