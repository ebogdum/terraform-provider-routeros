//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSystemRouterboardButtons(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	// Only the mode button is exercised, and it is left bound to the built-in
	// dark-mode event rather than a script, so the test does not depend on any
	// script existing on the device.
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

resource "routeros_system_routerboard_mode_button" "acc" {
  router    = "home"
  enabled   = true
  hold_time = "0s..1m"
  on_event  = "dark-mode"
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_system_routerboard_mode_button.acc", "id"),
				resource.TestCheckResourceAttr("routeros_system_routerboard_mode_button.acc", "enabled", "true"),
				resource.TestCheckResourceAttr("routeros_system_routerboard_mode_button.acc", "on_event", "dark-mode"),
				resource.TestCheckResourceAttr("routeros_system_routerboard_mode_button.acc", "hold_time", "0s..1m"),
			)},
			{
				ResourceName:      "routeros_system_routerboard_mode_button.acc",
				ImportState:       true,
				ImportStateId:     "home",
				ImportStateVerify: false,
			},
		},
	})
}
