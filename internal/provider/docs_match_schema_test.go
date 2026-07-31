package provider

import (
	"os"
	"regexp"
	"testing"
)

var docBulletRe = regexp.MustCompile(`(?m)^\*\s+` + "`" + `([a-z0-9_]+)` + "`")

// Docs must match the code exactly: every schema attribute is documented, and
// every documented attribute exists in the schema. No missing, no invented.
// Verified as a build-time invariant so drift cannot creep back in.
func TestDocsMatchSchema(t *testing.T) {
	var missing, invented int
	for _, tc := range allResourceSchemas(t) {
		docPath := "../../docs/resources/" + tc.slug + ".md"
		raw, err := os.ReadFile(docPath)
		if err != nil {
			t.Errorf("%s: no doc file", tc.slug)
			continue
		}
		documented := map[string]bool{}
		for _, m := range docBulletRe.FindAllStringSubmatch(string(raw), -1) {
			documented[m[1]] = true
		}
		for a := range tc.attrs {
			if a == "id" || a == "router" {
				continue
			}
			if !documented[a] {
				t.Errorf("%s.%s is in the schema but not documented", tc.slug, a)
				missing++
			}
		}
		for a := range documented {
			if a == "id" || a == "router" {
				continue
			}
			if !tc.attrs[a] {
				t.Errorf("%s doc references %q which is not a schema attribute", tc.slug, a)
				invented++
			}
		}
	}
	if missing+invented > 0 {
		t.Fatalf("doc/schema drift: %d missing, %d invented", missing, invented)
	}
}
