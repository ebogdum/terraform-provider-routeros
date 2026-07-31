//go:build acceptance

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIPDNSAdlist(t *testing.T) {
	if os.Getenv("TF_ACC") == "" || os.Getenv("ROUTEROS_HOST") == "" {
		t.Skip("TF_ACC and ROUTEROS_HOST required")
	}
	// Created disabled so the router never downloads the list during the test.
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

resource "routeros_ip_dns_adlist" "acc" {
  router   = "home"
  url      = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
  disabled = true
  comment  = "tf-acc-adlist"
}
`
	cfg = formatProviderCfg(cfg)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_ip_dns_adlist.acc", "id"),
				resource.TestCheckResourceAttr("routeros_ip_dns_adlist.acc", "disabled", "true"),
				resource.TestCheckResourceAttr("routeros_ip_dns_adlist.acc", "comment", "tf-acc-adlist"),
			)},
		},
	})
}
