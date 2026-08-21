package schemautil

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
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

// These attributes are Optional+Computed and "" was accepted before they gained
// a validator; failing the plan on it would break existing configurations.
func TestEmptyStringMeansUnsetNotInvalid(t *testing.T) {
	for name, v := range map[string]validator.String{
		"IsMAC":                IsMAC(),
		"IsDurationRouterOS":   IsDurationRouterOS(),
		"IsCIDR":               IsCIDR(),
		"IsTimeOfDayOrStartup": IsTimeOfDayOrStartup(),
	} {
		if !checkString(v, "") {
			t.Errorf("%s rejected the empty string", name)
		}
	}
	// Enumerations are not format validators: "" is a value, not an absence.
	if checkString(OneOfFold("MD5", "SHA1"), "") {
		t.Error("OneOfFold accepted the empty string")
	}
}
