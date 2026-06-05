//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacePPTPClient(t *testing.T) {
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

resource "routeros_interface_pptp_client" "acc" {
  router = "home"
  connect_to = "127.0.0.1"
  name = "tf-example"
  user = "myuser"
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_interface_pptp_client.acc", "id"),
			)},
		},
	})
}
