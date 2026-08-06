//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// forwarding-override names the switch port this port may forward to; RouterOS
// completes it with the port list and answers a boolean write with "input does
// not match any value of port". The menu has no `add` (its rows are created by
// the device, one per switch port), so this reads an existing row via import
// rather than creating one -- see TestAccInterfaceEthernetSwitchPortIsolation
// for the create path.
func TestAccInterfaceEthernetSwitchPortIsolationImport(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	port := os.Getenv("ROUTEROS_SWITCH_PORT")
	if port == "" {
		t.Skip("ROUTEROS_SWITCH_PORT required (name of a switch port, e.g. ether4)")
	}
	cfg := formatProviderCfg(`
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

resource "routeros_interface_ethernet_switch_port_isolation" "imported" {
  router = "home"
  name   = "` + port + `"
}
`)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:             cfg,
				ResourceName:       "routeros_interface_ethernet_switch_port_isolation.imported",
				ImportState:        true,
				ImportStateId:      "home/" + port,
				ImportStateVerify:  false,
				ImportStatePersist: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"routeros_interface_ethernet_switch_port_isolation.imported", "name", port),
				),
			},
		},
	})
}
