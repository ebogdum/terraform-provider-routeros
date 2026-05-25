package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// handwrittenResources holds resources implemented by hand (rather than
// generated). These exist either to validate the runtime before code-gen lands,
// or for menus whose semantics resist generation. The generator is the long-term
// home of nearly everything; this list shrinks over time.
func handwrittenResources() []func() resource.Resource {
	return []func() resource.Resource{}
}

func handwrittenDataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
