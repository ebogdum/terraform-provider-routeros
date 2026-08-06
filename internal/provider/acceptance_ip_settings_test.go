//go:build acceptance

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIPSettings(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	const tmpl = `
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

resource "routeros_ip_settings" "this" {
  router = "home"
%s
}
`
	// rp-filter is a RouterOS enum (no|loose|strict), not a bool.
	cfgEnum := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s", `  rp_filter = "loose"`))
	cfgOff := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s", `  rp_filter = "no"`))

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfgEnum, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_ip_settings.this", "id"),
				resource.TestCheckResourceAttr("routeros_ip_settings.this", "rp_filter", "loose"),
			)},
			// Restore the factory default.
			{Config: cfgOff, Check: resource.TestCheckResourceAttr(
				"routeros_ip_settings.this", "rp_filter", "no")},
		},
	})
}
