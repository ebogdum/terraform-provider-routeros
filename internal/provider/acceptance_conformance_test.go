//go:build acceptance

package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Conformance sweep: every reachable resource, every safely settable attribute,
// driven through create -> update -> destroy with two *different* values.
//
// This exists because the hand-written suite could not catch a wrong RouterOS
// property name. 296 of its 364 files set no attribute at all, and a name that
// is never set is never put on the wire; a wrong read key is a skipped branch
// that leaves the attribute null, which is legal for Optional+Computed. So 202
// broken names sat behind a green suite.
//
// The manifest is generated from the device's own schema by
// tools/conformance/gen_manifest.go's Python counterpart, so the values are ones
// RouterOS actually accepts rather than ones a human guessed. Regenerate it
// whenever the provider or the RouterOS version changes:
//
//	make conformance-manifest
//
// Each resource is its own subtest, so one broken menu reports as one failure
// instead of stopping the sweep. Run the dead-man switch alongside it:
//
//	make dms-backup dms-install
//	make dms-arm &            # keeps the watchdog fed
//	make conformance
//	make dms-disarm

type conformanceAttr struct {
	Kind string `json:"kind"`
	A    string `json:"a"`
	B    string `json:"b"`
}

type conformanceRequired struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

type conformanceResource struct {
	Menu       string                         `json:"menu"`
	Addable    bool                           `json:"addable"`
	Key        string                         `json:"key"`
	Required   map[string]conformanceRequired `json:"required"`
	Attributes map[string]conformanceAttr     `json:"attributes"`
}

type conformanceManifest struct {
	Resources map[string]conformanceResource `json:"resources"`
	Skipped   map[string]string              `json:"skipped"`
}

// hclValue renders a manifest value as HCL of the right type. Bools and numbers
// must not be quoted or Terraform rejects the config before the provider ever
// sees it.
func hclValue(a conformanceAttr, which string) string {
	v := a.A
	if which == "b" {
		v = a.B
	}
	switch a.Kind {
	case "Bool", "Int64", "Float64", "Number":
		return v
	case "Set", "List":
		// RouterOS comma-joins multi-valued properties; Terraform wants a
		// collection, so each token becomes an element.
		parts := strings.Split(v, ",")
		for i, p := range parts {
			parts[i] = fmt.Sprintf("%q", strings.TrimSpace(p))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	default:
		return fmt.Sprintf("%q", strings.ReplaceAll(v, `"`, `\"`))
	}
}

// conformanceConfig renders a resource carrying only `set` -- one attribute at
// a time. Setting every attribute at once looks efficient but produces
// combinations RouterOS rejects on semantic grounds (offer chacha20poly1305
// together with an auth algorithm and /ip/ipsec/proposal answers "AEAD already
// provides authentication"). That is a property of the protocol, not a provider
// defect, and it would mask the defects this sweep exists to find. Isolating
// each attribute also means a failure names the exact attribute rather than a
// combination.
func conformanceConfig(typeName string, r conformanceResource, set map[string]string) string {
	var b strings.Builder
	b.WriteString(formatProviderCfg(`
provider "routeros" {
  routers = {
    home = {
      host     = "%s"
      username = "%s"
      password = "%s"
      insecure = true
    }
  }
}
`))
	fmt.Fprintf(&b, "\nresource %q \"conf\" {\n  router = \"home\"\n", typeName)
	if r.Addable && r.Key != "" {
		fmt.Fprintf(&b, "  %s = \"tfacc-conf\"\n", r.Key)
	}
	// Required attributes are pinned into every config: Terraform rejects the
	// config outright if one is absent, so the provider is never reached and
	// the attribute under test is never exercised.
	reqNames := make([]string, 0, len(r.Required))
	for n := range r.Required {
		if _, dup := set[n]; !dup && n != r.Key {
			reqNames = append(reqNames, n)
		}
	}
	sort.Strings(reqNames)
	for _, n := range reqNames {
		fmt.Fprintf(&b, "  %s = %s\n", n,
			hclValue(conformanceAttr{Kind: r.Required[n].Kind, A: r.Required[n].Value}, "a"))
	}
	names := make([]string, 0, len(set))
	for n := range set {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		fmt.Fprintf(&b, "  %s = %s\n", n, set[n])
	}
	b.WriteString("}\n")
	return b.String()
}

func TestAccConformanceSweep(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	raw, err := os.ReadFile("testdata/conformance.json")
	if err != nil {
		t.Fatalf("read manifest (run `make conformance-manifest` first): %v", err)
	}
	var m conformanceManifest
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("parse manifest: %v", err)
	}

	only := os.Getenv("CONFORMANCE_ONLY") // substring filter, for triage
	names := make([]string, 0, len(m.Resources))
	for n := range m.Resources {
		if only == "" || strings.Contains(n, only) {
			names = append(names, n)
		}
	}
	sort.Strings(names)
	t.Logf("sweeping %d resources, %d attributes", len(names), func() int {
		n := 0
		for _, r := range names {
			n += len(m.Resources[r].Attributes)
		}
		return n
	}())

	for _, name := range names {
		r := m.Resources[name]
		t.Run(name, func(t *testing.T) {
			addr := name + ".conf"

			// One check per attribute. Collections are addressed by index, so
			// assert the element count rather than a scalar equality.
			check := func(attrName string, a conformanceAttr, which string) resource.TestCheckFunc {
				want := a.A
				if which == "b" {
					want = a.B
				}
				switch a.Kind {
				case "Set", "List":
					return resource.TestCheckResourceAttr(addr, attrName+".#",
						fmt.Sprint(len(strings.Split(want, ","))))
				default:
					return resource.TestCheckResourceAttr(addr, attrName, want)
				}
			}

			attrNames := make([]string, 0, len(r.Attributes))
			for n := range r.Attributes {
				attrNames = append(attrNames, n)
			}
			sort.Strings(attrNames)

			// Each attribute gets a fresh resource. Reusing one instance across
			// attributes does not isolate them: dropping an attribute from the
			// config stops Terraform sending it, but the value stays on the
			// device, so a later attribute still collides with it. A new
			// instance per attribute is the only way a failure means "this
			// attribute is broken" rather than "something earlier set a
			// conflicting value".
			for _, an := range attrNames {
				a := r.Attributes[an]
				t.Run(an, func(t *testing.T) {
					resource.Test(t, resource.TestCase{
						ProtoV6ProviderFactories: testAccProviderFactories,
						Steps: []resource.TestStep{
							{ // create carrying this attribute
								Config: conformanceConfig(name, r,
									map[string]string{an: hclValue(a, "a")}),
								Check: resource.ComposeAggregateTestCheckFunc(
									resource.TestCheckResourceAttrSet(addr, "id"),
									check(an, a, "a"),
								),
							},
							{ // update it to a different legal value
								Config: conformanceConfig(name, r,
									map[string]string{an: hclValue(a, "b")}),
								Check: resource.ComposeAggregateTestCheckFunc(
									resource.TestCheckResourceAttrSet(addr, "id"),
									check(an, a, "b"),
								),
							},
						},
					})
				})
			}
		})
	}
}
