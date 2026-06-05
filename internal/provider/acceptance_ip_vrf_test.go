//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIPVrf(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	cfg := `
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

resource "routeros_ip_vrf" "acc" {
  router = "home"
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, PlanOnly: true, ExpectNonEmptyPlan: true},
		},
	})
	_ = "routeros_ip_vrf" // silence unused
}
