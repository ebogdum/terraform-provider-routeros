// Hand-written lifecycle + drift acceptance test. NAMED with the
// "_acc_test.go" suffix (NOT "acceptance_..._test.go") so the regen
// wildcard `find ... -name 'acceptance_*.go'` does not delete it.
//
//go:build acceptance

package provider

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// TestAccLifecycleIPAddress runs the full lifecycle against /ip/address:
//
//	step 1: Create address=10.99.0.1/24 comment="lc-1"
//	step 2: Re-apply same config; expect empty plan (no drift)
//	step 3: Change comment to "lc-2"; expect Update
//	step 4: Re-apply; expect empty plan again
func TestAccLifecycleIPAddress(t *testing.T) {
	skipIfNoAccLifecycle(t)
	cfg := func(comment string) string {
		return fmt.Sprintf(`
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

data "routeros_interface" "all" { router = "home" }

resource "routeros_ip_address" "lc" {
  router    = "home"
  address   = "10.99.0.1/24"
  interface = data.routeros_interface.all.records[0].name
  comment   = "%s"
}
`,
			os.Getenv("ROUTEROS_HOST"),
			os.Getenv("ROUTEROS_USER"),
			os.Getenv("ROUTEROS_PASSWORD"),
			comment)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg("lc-1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("routeros_ip_address.lc", "comment", "lc-1"),
					resource.TestCheckResourceAttr("routeros_ip_address.lc", "address", "10.99.0.1/24"),
				),
			},
			{
				Config: cfg("lc-1"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				Config: cfg("lc-2"),
				Check:  resource.TestCheckResourceAttr("routeros_ip_address.lc", "comment", "lc-2"),
			},
			{
				Config: cfg("lc-2"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}

// TestAccDriftCorrectionIPAddress verifies out-of-band changes are detected
// and reconciled on the next apply.
//
//	step 1: Create comment="drift-1"
//	PreConfig of step 2: PATCH the row's comment to "drift-EXTERNAL" via REST
//	step 2: Re-apply with comment="drift-1" -- provider should plan + apply
//	        an Update that resets comment back to "drift-1"
func TestAccDriftCorrectionIPAddress(t *testing.T) {
	skipIfNoAccLifecycle(t)
	host := os.Getenv("ROUTEROS_HOST")
	user := os.Getenv("ROUTEROS_USER")
	pass := os.Getenv("ROUTEROS_PASSWORD")

	cfg := fmt.Sprintf(`
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

data "routeros_interface" "all" { router = "home" }

resource "routeros_ip_address" "drift" {
  router    = "home"
  address   = "10.98.0.1/24"
  interface = data.routeros_interface.all.records[0].name
  comment   = "drift-1"
}
`, host, user, pass)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check:  resource.TestCheckResourceAttr("routeros_ip_address.drift", "comment", "drift-1"),
			},
			{
				PreConfig: func() { mustExternalPatch(t, host, user, pass, "drift-1", "drift-EXTERNAL") },
				Config:    cfg,
				Check:     resource.TestCheckResourceAttr("routeros_ip_address.drift", "comment", "drift-1"),
			},
		},
	})
}

func mustExternalPatch(t *testing.T, host, user, pass, markerFrom, markerTo string) {
	t.Helper()
	req, err := http.NewRequest("GET", strings.TrimRight(host, "/")+"/rest/ip/address", nil)
	if err != nil {
		t.Fatalf("external probe: %v", err)
	}
	req.SetBasicAuth(user, pass)
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("external GET: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("external GET: status %d", resp.StatusCode)
	}
	body := readAllResp(t, resp)
	id := extractIDForComment(body, markerFrom)
	if id == "" {
		t.Fatalf("external GET: no row with comment %q in %s", markerFrom, body)
	}
	patchURL := strings.TrimRight(host, "/") + "/rest/ip/address/" + id
	patch, err := http.NewRequest("PATCH", patchURL,
		strings.NewReader(fmt.Sprintf(`{"comment":%q}`, markerTo)))
	if err != nil {
		t.Fatalf("external PATCH build: %v", err)
	}
	patch.Header.Set("Content-Type", "application/json")
	patch.Header.Set("Accept-Encoding", "gzip")
	patch.SetBasicAuth(user, pass)
	patch.URL.RawPath = patch.URL.Path
	pResp, err := http.DefaultClient.Do(patch)
	if err != nil {
		t.Fatalf("external PATCH: %v", err)
	}
	defer pResp.Body.Close()
	if pResp.StatusCode != 200 {
		body := readAllResp(t, pResp)
		t.Fatalf("external PATCH: status %d body=%s", pResp.StatusCode, body)
	}
}

func readAllResp(t *testing.T, r *http.Response) string {
	t.Helper()
	var b strings.Builder
	buf := make([]byte, 4096)
	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			b.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}
	return b.String()
}

func extractIDForComment(body, marker string) string {
	rows := strings.Split(body, "},{")
	for _, row := range rows {
		if !strings.Contains(row, `"comment":"`+marker+`"`) {
			continue
		}
		i := strings.Index(row, `".id":"`)
		if i < 0 {
			continue
		}
		rest := row[i+len(`".id":"`):]
		j := strings.IndexByte(rest, '"')
		if j < 0 {
			continue
		}
		return rest[:j]
	}
	return ""
}

func skipIfNoAccLifecycle(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
}
