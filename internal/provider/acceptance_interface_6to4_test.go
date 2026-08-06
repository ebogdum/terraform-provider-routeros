//go:build acceptance

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterface6to4(t *testing.T) {
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

resource "routeros_interface_6to4" "acc" {
  router         = "home"
  name           = "acc-6to4"
  local_address  = "0.0.0.0"
  remote_address = "192.0.2.1"
  dont_fragment  = %q
}
`
	// dont-fragment is a RouterOS enum (no|inherit), not a bool: the menu
	// rejects yes outright ("must be either inherit or no"), so the old bool
	// schema could not express inherit and turned true into a hard 400.
	cfgInherit := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s", "inherit"))
	cfgNo := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s", "no"))

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfgInherit, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_interface_6to4.acc", "id"),
				resource.TestCheckResourceAttr("routeros_interface_6to4.acc", "dont_fragment", "inherit"),
			)},
			{Config: cfgNo, Check: resource.TestCheckResourceAttr(
				"routeros_interface_6to4.acc", "dont_fragment", "no")},
		},
	})
}
