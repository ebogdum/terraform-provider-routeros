package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// TestNullifyUnknownAttrs proves the post-apply sweep resolves every Unknown
// attribute to Null while leaving known values (user-set Optional, prior state)
// untouched. This is the guard against "Provider returned invalid result object
// after apply", which fires when a Computed attribute the device never returns
// is left at the plan's Unknown after a fresh Create.
func TestNullifyUnknownAttrs(t *testing.T) {
	type model struct {
		Unset    types.String // unknown -> must become null
		UserSet  types.String // known -> must stay
		FromRead types.String // null -> must stay null
		ABool    types.Bool   // unknown -> null
		AnInt    types.Int64  // unknown -> null
	}
	m := &model{
		Unset:    types.StringUnknown(),
		UserSet:  types.StringValue("keep-me"),
		FromRead: types.StringNull(),
		ABool:    types.BoolUnknown(),
		AnInt:    types.Int64Unknown(),
	}
	nullifyUnknownAttrs(m)

	if !m.Unset.IsNull() {
		t.Errorf("Unset: want null, got %v", m.Unset)
	}
	if m.UserSet.ValueString() != "keep-me" {
		t.Errorf("UserSet: want keep-me, got %v", m.UserSet)
	}
	if !m.FromRead.IsNull() {
		t.Errorf("FromRead: want null, got %v", m.FromRead)
	}
	if !m.ABool.IsNull() {
		t.Errorf("ABool: want null, got %v", m.ABool)
	}
	if !m.AnInt.IsNull() {
		t.Errorf("AnInt: want null, got %v", m.AnInt)
	}

	// A non-pointer / non-struct argument must be a harmless no-op.
	nullifyUnknownAttrs(nil)
	nullifyUnknownAttrs(model{})
	nullifyUnknownAttrs("not a struct")
}

// The switch named only the basetypes, so a custom-typed attribute stayed
// Unknown and reached Terraform as an invalid result. Every ordered firewall
// resource carries `time` as a csvSetValue, which is how it surfaced.
func TestNullifyUnknownAttrsResolvesCustomAndCollectionTypes(t *testing.T) {
	type model struct {
		Time     csvSetValue
		AList    types.List
		ASet     types.Set
		AMap     types.Map
		KeptTime csvSetValue
	}
	m := &model{
		Time:     csvSetValue{StringValue: basetypes.NewStringUnknown()},
		AList:    types.ListUnknown(types.StringType),
		ASet:     types.SetUnknown(types.StringType),
		AMap:     types.MapUnknown(types.StringType),
		KeptTime: csvSetValue{StringValue: basetypes.NewStringValue("mon,tue")},
	}
	nullifyUnknownAttrs(m)

	for name, v := range map[string]attr.Value{
		"Time": m.Time, "AList": m.AList, "ASet": m.ASet, "AMap": m.AMap,
	} {
		if v.IsUnknown() {
			t.Errorf("%s: still unknown", name)
		}
		if !v.IsNull() {
			t.Errorf("%s: want null, got %v", name, v)
		}
	}
	if m.KeptTime.ValueString() != "mon,tue" {
		t.Errorf("KeptTime: want mon,tue, got %v", m.KeptTime)
	}
}
