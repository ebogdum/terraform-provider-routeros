package provider

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// An Update body write of the form `if !plan.X.Equal(state.X) { body["k"]=... }`
// without a `!plan.X.IsUnknown()` guard writes a blank string when the attribute
// is a Computed value that went Unknown on change — RouterOS then rejects it
// ("invalid value"). Every Update diff write must guard against Unknown.
func TestNoBlankOrUnknownWrites(t *testing.T) {
	files, _ := filepath.Glob("resource_*.go")
	unguarded := regexp.MustCompile(`if !plan\.(\w+)\.Equal\(state\.\w+\) \{\n\s*body\[`)
	var offenders []string
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		src, _ := os.ReadFile(f)
		for _, m := range unguarded.FindAllStringSubmatch(string(src), -1) {
			offenders = append(offenders, f+": "+m[1])
		}
	}
	if len(offenders) > 0 {
		t.Errorf("%d Update writes lack the !IsUnknown guard (blank-write risk):", len(offenders))
		for _, o := range offenders[:min(len(offenders), 12)] {
			t.Errorf("  %s", o)
		}
	}
}
