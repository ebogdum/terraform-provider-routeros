package schemautil

import (
	"testing"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// RouterOS writes must use yes/no for bools — several menus reject true/false.
func TestBoolYesNo(t *testing.T) {
	if got := client.FormatBool(true); got != "yes" {
		t.Errorf("FormatBool(true) = %q, want yes", got)
	}
	if got := client.FormatBool(false); got != "no" {
		t.Errorf("FormatBool(false) = %q, want no", got)
	}
	// Reads must still accept both forms RouterOS may return.
	for _, s := range []string{"yes", "true", "no", "false", "YES", "False"} {
		if _, err := client.ParseBool(s); err != nil {
			t.Errorf("ParseBool(%q) failed: %v", s, err)
		}
	}
}
