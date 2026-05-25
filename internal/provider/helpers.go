package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// diagBuf is the diagnostics type used by generated Apply/Upsert helpers.
type diagBuf = diag.Diagnostics

// configureRegistry is called from every generated resource's and data source's
// Configure(). It pulls a *client.Registry out of ProviderData and stores it.
// Returns the registry plus a diagnostic if the provider isn't configured yet.
func configureRegistry(providerData any) (*client.Registry, diag.Diagnostics) {
	var d diag.Diagnostics
	if providerData == nil {
		return nil, d
	}
	reg, ok := providerData.(*client.Registry)
	if !ok {
		d.AddError("Unexpected provider data", fmt.Sprintf("got %T, want *client.Registry", providerData))
		return nil, d
	}
	return reg, d
}

// pickClient resolves a `router` attribute value to a *Client, falling back to
// the registry's default.
func pickClient(reg *client.Registry, routerAttr types.String, diags *diag.Diagnostics) *client.Client {
	if reg == nil {
		diags.AddError("Provider not configured", "the provider was not fully initialised before this resource was used")
		return nil
	}
	name := ""
	if !routerAttr.IsNull() && !routerAttr.IsUnknown() {
		name = routerAttr.ValueString()
	}
	c, err := reg.Get(name)
	if err != nil {
		diags.AddError("Unknown router", err.Error())
		return nil
	}
	return c
}

// boolFromObj decodes a RouterOS bool string. Missing/empty -> Null.
func boolFromObj(obj client.Object, key string) types.Bool {
	v, ok := obj[key]
	if !ok || v == "" {
		return types.BoolNull()
	}
	b, err := client.ParseBool(v)
	if err != nil {
		return types.BoolNull()
	}
	return types.BoolValue(b)
}

func listStringValidators(vs ...validator.String) []validator.String { return vs }
