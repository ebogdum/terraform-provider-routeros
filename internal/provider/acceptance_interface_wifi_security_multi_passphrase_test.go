//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfaceWifiSecurityMultiPassphrase(t *testing.T) {
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

resource "routeros_interface_wifi_security_multi_passphrase" "acc" {
  router     = "home"
  group      = "tf_acc_multi_passphrase"
  passphrase = "tf_acc_passphrase"
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_interface_wifi_security_multi_passphrase.acc", "id"),
				resource.TestCheckResourceAttr("routeros_interface_wifi_security_multi_passphrase.acc", "group", "tf_acc_multi_passphrase"),
			)},
		},
	})
}
