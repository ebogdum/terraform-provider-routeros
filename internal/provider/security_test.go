package provider

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Attribute names that look secret-bearing but are not actually secrets:
// format selectors, group-name references, and numeric policy knobs.
var notReallySecret = map[string]bool{
	"password_format":         true, // how to format a RADIUS password, not the value
	"multi_passphrase_group":  true, // a group name reference
	"minimum_password_length": true, // a length policy number
	"password_authentication": true, // an enable/disable toggle
}

var secretNameRe = regexp.MustCompile(`(?i)passphrase|password|(^|_)secret($|_)|private_key|preshared|(^|_)psk($|_)|otp_secret`)

// Every attribute whose name denotes a secret must be marked Sensitive so
// Terraform redacts it in plan output, state diffs and CLI display. A cleartext
// private key or pre-shared key in the plan is a real credential leak.
func TestSecretsAreSensitive(t *testing.T) {
	for _, f := range registryResources() {
		r := f()
		mr := &resource.MetadataResponse{}
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, mr)
		sr := &resource.SchemaResponse{}
		r.Schema(context.Background(), resource.SchemaRequest{}, sr)
		res := strings.TrimPrefix(mr.TypeName, "routeros_")
		for name, a := range sr.Schema.Attributes {
			if !secretNameRe.MatchString(name) || notReallySecret[name] {
				continue
			}
			if !a.IsSensitive() {
				t.Errorf("%s.%s is secret-bearing but not marked Sensitive", res, name)
			}
		}
	}
}
