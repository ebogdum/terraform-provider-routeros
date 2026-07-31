package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type resSchema struct {
	slug  string
	attrs map[string]bool
}

func allResourceSchemas(t *testing.T) []resSchema {
	t.Helper()
	var out []resSchema
	for _, f := range registryResources() {
		r := f()
		mr := &resource.MetadataResponse{}
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, mr)
		sr := &resource.SchemaResponse{}
		r.Schema(context.Background(), resource.SchemaRequest{}, sr)
		if sr.Diagnostics.HasError() {
			t.Fatalf("%s schema error", mr.TypeName)
		}
		attrs := map[string]bool{}
		for name := range sr.Schema.Attributes {
			attrs[name] = true
		}
		out = append(out, resSchema{slug: strings.TrimPrefix(mr.TypeName, "routeros_"), attrs: attrs})
	}
	return out
}
