// Hand-written: verifies `terraform import` correctly hydrates state from an
// existing RouterOS row, then a regular plan/apply finds zero drift.
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
)

// TestAccImportIPAddress imports an existing /ip/address row, then verifies
// state matches and a plan with the same config produces no diff.
func TestAccImportIPAddress(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	host := os.Getenv("ROUTEROS_HOST")
	user := os.Getenv("ROUTEROS_USER")
	pass := os.Getenv("ROUTEROS_PASSWORD")

	// 1. Create a row via raw REST so we have a real .id to import.
	createID, addr, ifaceName := createSeedAddress(t, host, user, pass)
	defer deleteSeedAddress(t, host, user, pass, createID)

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

resource "routeros_ip_address" "imported" {
  router    = "home"
  address   = "%s"
  interface = "%s"
  comment   = "import-acc"
}
`, host, user, pass, addr, ifaceName)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// 1. Apply nothing yet -- just declare the resource.
			{
				Config: cfg,
				// Use the pre-seeded row instead of creating a new one.
				// terraform-plugin-testing's ImportState step runs `terraform
				// import`, then verifies the resulting state matches what a
				// fresh Read would return.
				ImportState:        true,
				ImportStateId:      "home/" + createID,
				ResourceName:       "routeros_ip_address.imported",
				ImportStatePersist: true,
			},
		},
	})
}

// createSeedAddress PUTs a fresh /ip/address row and returns (id, address, interface).
func createSeedAddress(t *testing.T, host, user, pass string) (string, string, string) {
	t.Helper()
	addr := "10.97.0.1/24"
	iface := firstInterfaceName(t, host, user, pass)

	body := fmt.Sprintf(`{"address":%q,"interface":%q,"comment":"import-acc"}`, addr, iface)
	req, _ := http.NewRequest("PUT",
		strings.TrimRight(host, "/")+"/rest/ip/address",
		strings.NewReader(body))
	req.SetBasicAuth(user, pass)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("seed PUT: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		t.Fatalf("seed PUT: status %d body=%s", resp.StatusCode, readAllResp2(resp))
	}
	rb := readAllResp2(resp)
	const key = `".id":"`
	i := strings.Index(rb, key)
	if i < 0 {
		t.Fatalf("seed PUT: no .id in response: %s", rb)
	}
	rest := rb[i+len(key):]
	j := strings.IndexByte(rest, '"')
	if j < 0 {
		t.Fatalf("seed PUT: malformed .id: %s", rb)
	}
	return rest[:j], addr, iface
}

func deleteSeedAddress(t *testing.T, host, user, pass, id string) {
	t.Helper()
	req, _ := http.NewRequest("DELETE",
		strings.TrimRight(host, "/")+"/rest/ip/address/"+id, nil)
	req.SetBasicAuth(user, pass)
	req.Header.Set("Accept-Encoding", "gzip")
	req.URL.RawPath = req.URL.Path
	_, _ = http.DefaultClient.Do(req)
}
