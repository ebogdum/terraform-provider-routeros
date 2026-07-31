//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIPService(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	// telnet is off by default on a stock RouterOS install, so toggling it does
	// not touch the session this test runs over.
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

resource "routeros_ip_service" "acc" {
  router   = "home"
  name     = "telnet"
  disabled = true
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_ip_service.acc", "id"),
				resource.TestCheckResourceAttr("routeros_ip_service.acc", "name", "telnet"),
				resource.TestCheckResourceAttr("routeros_ip_service.acc", "disabled", "true"),
				resource.TestCheckResourceAttrSet("routeros_ip_service.acc", "port"),
			)},
			{
				ResourceName:      "routeros_ip_service.acc",
				ImportState:       true,
				ImportStateId:     "home/telnet",
				ImportStateVerify: false,
			},
		},
	})
}
