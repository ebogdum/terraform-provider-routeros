//go:build acceptance

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccToolEMail(t *testing.T) {
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

resource "routeros_tool_e_mail" "this" {
  router = "home"
%s
}
`
	// tls and certificate-verification are RouterOS enums, not bools: the menu
	// also accepts starttls and yes-without-crl, and rejects true/false outright
	// ("input does not match any value of tls"). Exercise the non-boolean values
	// so a regression back to a bool schema fails here.
	cfgEnum := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s",
		"  tls = \"starttls\"\n  certificate_verification = \"yes-without-crl\""))
	cfgOff := formatProviderCfg(fmt.Sprintf(tmpl, "%s", "%s", "%s",
		"  tls = \"no\"\n  certificate_verification = \"no\""))

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{Config: cfgEnum, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("routeros_tool_e_mail.this", "id"),
				resource.TestCheckResourceAttr("routeros_tool_e_mail.this", "tls", "starttls"),
				resource.TestCheckResourceAttr("routeros_tool_e_mail.this", "certificate_verification", "yes-without-crl"),
			)},
			// Leave the menu on its factory-default setting.
			{Config: cfgOff, Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("routeros_tool_e_mail.this", "tls", "no"),
				resource.TestCheckResourceAttr("routeros_tool_e_mail.this", "certificate_verification", "no"),
			)},
		},
	})
}
