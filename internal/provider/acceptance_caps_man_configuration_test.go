//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCapsManConfiguration(t *testing.T) {
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

resource "routeros_caps_man_configuration" "acc" {
  router = "home"
  name = "tf-example"
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_caps_man_configuration.acc", "id"),
			)},
		},
	})
}
