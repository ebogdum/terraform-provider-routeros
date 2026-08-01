package provider

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Singleton resources must send only CHANGED fields on Update, never the full
// body. A full-body write re-pushes device values the user never touched — that
// is what activated the CAPsMAN controller and took the AP's wifi down. Every
// `body["k"] = ...` in a SetSingleton resource must be guarded by a
// `state == nil || !plan.X.Equal(state.X)` diff check.
func TestSingletonUpdateDiffs(t *testing.T) {
	files, _ := filepath.Glob("resource_*.go")
	bodyWrite := regexp.MustCompile(`\n(\s*)if ([^\n]*?) \{\n\s*body\["[^"]+"\] =`)
	var offenders []string
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		src, _ := os.ReadFile(f)
		s := string(src)
		if !strings.Contains(s, "SetSingleton(ctx") {
			continue
		}
		for _, m := range bodyWrite.FindAllStringSubmatch(s, -1) {
			cond := m[2]
			if !strings.Contains(cond, "state == nil || !plan.") {
				offenders = append(offenders, f+": "+strings.TrimSpace(cond))
			}
		}
	}
	if len(offenders) > 0 {
		t.Errorf("%d singleton body writes are not diff-guarded (full-body write risk):", len(offenders))
		for _, o := range offenders[:min(len(offenders), 15)] {
			t.Errorf("  %s", o)
		}
	}
}
