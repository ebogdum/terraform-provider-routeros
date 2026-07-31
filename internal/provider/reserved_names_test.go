package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Terraform meta-arguments and reserved block names. A resource/data-source
// attribute with any of these names makes Terraform core reject the entire
// provider at schema-load time — every plan against the provider fails, not
// just that resource. The plugin-framework does NOT flag this when Schema() is
// built, so it must be asserted here.
var reservedAttributeNames = map[string]bool{
	"count":       true,
	"for_each":    true,
	"depends_on":  true,
	"lifecycle":   true,
	"provider":    true,
	"provisioner": true,
	"connection":  true,
}

func TestNoReservedAttributeNames(t *testing.T) {
	for _, f := range registryResources() {
		r := f()
		mr := &resource.MetadataResponse{}
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, mr)
		sr := &resource.SchemaResponse{}
		r.Schema(context.Background(), resource.SchemaRequest{}, sr)
		res := strings.TrimPrefix(mr.TypeName, "routeros_")
		for name := range sr.Schema.Attributes {
			if reservedAttributeNames[name] {
				t.Errorf("%s declares reserved attribute %q — Terraform rejects the whole provider", res, name)
			}
		}
	}
}
