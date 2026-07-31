package schemautil

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// checkString runs a string validator and reports whether it accepted the value.
func checkString(v validator.String, s string) bool {
	resp := &validator.StringResponse{}
	v.ValidateString(context.Background(), validator.StringRequest{
		Path:        path.Root("test"),
		ConfigValue: types.StringValue(s),
	}, resp)
	return !resp.Diagnostics.HasError()
}

func TestOneOfFold(t *testing.T) {
	v := OneOfFold("MD5", "SHA1")
	// RouterOS reports MD5; the CLI documents md5. Both must be accepted.
	for _, ok := range []string{"MD5", "md5", "Md5", "SHA1", "sha1"} {
		if !checkString(v, ok) {
			t.Errorf("OneOfFold rejected %q, want accepted", ok)
		}
	}
	for _, bad := range []string{"sha256", "", "md-5"} {
		if checkString(v, bad) {
			t.Errorf("OneOfFold accepted %q, want rejected", bad)
		}
	}
	// The case-sensitive OneOf must keep its old strictness.
	if checkString(OneOf("md5", "sha1"), "MD5") {
		t.Error("OneOf became case-insensitive; that is OneOfFold's job")
	}
}

func TestIsDurationOrKeyword(t *testing.T) {
	arp := IsDurationOrKeyword("auto")
	// The factory default on bridge/ethernet interfaces.
	for _, ok := range []string{"auto", "AUTO", " auto ", "30s", "1w2d3h", "120"} {
		if !checkString(arp, ok) {
			t.Errorf("arp-timeout validator rejected %q, want accepted", ok)
		}
	}
	for _, bad := range []string{"sometimes", "auto-ish", "5 furlongs"} {
		if checkString(arp, bad) {
			t.Errorf("arp-timeout validator accepted %q, want rejected", bad)
		}
	}

	dpd := IsDurationOrKeyword("disable-dpd")
	// 8s is the /ip/ipsec/profile default; the old OneOf allowed only the keyword.
	for _, ok := range []string{"8s", "disable-dpd", "2m"} {
		if !checkString(dpd, ok) {
			t.Errorf("dpd-interval validator rejected %q, want accepted", ok)
		}
	}
	if checkString(dpd, "auto") {
		t.Error("dpd-interval validator accepted an unrelated keyword")
	}
}

// planString runs a string plan modifier and returns the resulting plan value.
func planString(m planmodifier.String, config, state string) string {
	resp := &planmodifier.StringResponse{PlanValue: types.StringValue(config)}
	req := planmodifier.StringRequest{
		Path:        path.Root("test"),
		ConfigValue: types.StringValue(config),
		StateValue:  types.StringValue(state),
	}
	m.PlanModifyString(context.Background(), req, resp)
	if resp.Diagnostics.HasError() {
		return "ERROR"
	}
	return resp.PlanValue.ValueString()
}

func TestNormalizeDurationExceptPassesKeywordThrough(t *testing.T) {
	m := NormalizeDurationExcept("auto")
	// The bug: the plain duration modifier raised "Invalid value" on the
	// router's own default, so a bridge at defaults could not even be planned.
	if got := planString(m, "auto", "auto"); got == "ERROR" {
		t.Fatal("NormalizeDurationExcept errored on the keyword it was told to pass through")
	}
	if got := planString(NormalizeDuration(), "auto", "auto"); got != "ERROR" {
		t.Fatalf("plain NormalizeDuration accepted %q; expected it to still reject", "auto")
	}
	// Real durations still normalise against state.
	if got := planString(m, "60s", "1m"); got != "1m" {
		t.Errorf("duration normalization = %q, want %q", got, "1m")
	}
}

func TestNormalizeCase(t *testing.T) {
	m := NormalizeCase("MD5", "SHA1")
	// A config written lower-case must reach state in the device's spelling,
	// otherwise every plan shows a diff against a router reporting MD5.
	if got := planString(m, "md5", "MD5"); got != "MD5" {
		t.Errorf("NormalizeCase(md5) = %q, want MD5", got)
	}
	// Unlisted values pass through untouched; rejecting them is the validator's job.
	if got := planString(m, "whatever", "whatever"); got != "whatever" {
		t.Errorf("NormalizeCase mangled an unlisted value: %q", got)
	}
}
