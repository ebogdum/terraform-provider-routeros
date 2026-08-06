//go:build acceptance

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfaceSSTPServerServer(t *testing.T) {
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

resource "routeros_interface_sstp_server_server" "this" {
  router = "home"
%s
}
`
	// pfs is a RouterOS enum (no|yes|required), not a bool.
	cfgEnum := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s", `  pfs = "required"`))
	cfgOff := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s", `  pfs = "no"`))

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfgEnum, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_interface_sstp_server_server.this", "id"),
				resource.TestCheckResourceAttr("routeros_interface_sstp_server_server.this", "pfs", "required"),
			)},
			// Restore the factory default.
			{Config: cfgOff, Check: resource.TestCheckResourceAttr(
				"routeros_interface_sstp_server_server.this", "pfs", "no")},
		},
	})
}
