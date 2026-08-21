package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &CapsManManagerResource{}
	_ resource.ResourceWithImportState = &CapsManManagerResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type CapsManManagerResource struct {
	reg *client.Registry
}

type CapsManManagerModel struct {
	ID                     types.String `tfsdk:"id"`
	CACertificate          types.String `tfsdk:"ca_certificate"`
	Certificate            types.String `tfsdk:"certificate"`
	Enabled                types.Bool   `tfsdk:"enabled"`
	PackagePath            types.String `tfsdk:"package_path"`
	RequirePeerCertificate types.Bool   `tfsdk:"require_peer_certificate"`
	UpgradePolicy          types.String `tfsdk:"upgrade_policy"`
	Router                 types.String `tfsdk:"router"`
}

func NewCapsManManagerResource() resource.Resource { return &CapsManManagerResource{} }

func (r *CapsManManagerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_caps_man_manager"
}

func (r *CapsManManagerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CapsManManagerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/caps-man/manager`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ca_certificate": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"certificate": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"package_path": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"require_peer_certificate": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"upgrade_policy": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *CapsManManagerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CapsManManagerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	capsManManagerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManManagerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CapsManManagerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state CapsManManagerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	capsManManagerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManManagerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CapsManManagerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/caps-man/manager")
	if err != nil {
		resp.Diagnostics.AddError("Read /caps-man/manager failed", err.Error())
		return
	}
	capsManManagerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/caps-man/manager", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CapsManManagerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *CapsManManagerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/caps-man/manager" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/caps-man/manager", types.StringValue(routerName))))...)
}

func capsManManagerUpsert(ctx context.Context, reg *client.Registry, plan, state *CapsManManagerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CACertificate.IsNull() || plan.CACertificate.IsUnknown()) && (state == nil || !plan.CACertificate.Equal(state.CACertificate)) {
		body["ca-certificate"] = plan.CACertificate.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) && (state == nil || !plan.Certificate.Equal(state.Certificate)) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.PackagePath.IsNull() || plan.PackagePath.IsUnknown()) && (state == nil || !plan.PackagePath.Equal(state.PackagePath)) {
		body["package-path"] = plan.PackagePath.ValueString()
	}
	if !(plan.RequirePeerCertificate.IsNull() || plan.RequirePeerCertificate.IsUnknown()) && (state == nil || !plan.RequirePeerCertificate.Equal(state.RequirePeerCertificate)) {
		body["require-peer-certificate"] = client.FormatBool(plan.RequirePeerCertificate.ValueBool())
	}
	if !(plan.UpgradePolicy.IsNull() || plan.UpgradePolicy.IsUnknown()) && (state == nil || !plan.UpgradePolicy.Equal(state.UpgradePolicy)) {
		body["upgrade-policy"] = plan.UpgradePolicy.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/caps-man/manager", body)
	if err != nil {
		diags.AddError("Upsert /caps-man/manager failed", err.Error())
		return
	}
	capsManManagerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/caps-man/manager", plan.Router))
}

func capsManManagerApply(ctx context.Context, obj client.Object, m *CapsManManagerModel) {
	_ = ctx
	if v, ok := obj["ca-certificate"]; ok {
		_ = v
		if v != "" {
			m.CACertificate = types.StringValue(v)
		} else {
			m.CACertificate = types.StringNull()
		}
	}
	if v, ok := obj["certificate"]; ok {
		_ = v
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["package-path"]; ok {
		_ = v
		if v != "" {
			m.PackagePath = types.StringValue(v)
		} else {
			m.PackagePath = types.StringNull()
		}
	}
	if v, ok := obj["require-peer-certificate"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RequirePeerCertificate = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.RequirePeerCertificate = types.BoolValue(true)
		} else {
			m.RequirePeerCertificate = types.BoolNull()
		}
	}
	if v, ok := obj["upgrade-policy"]; ok {
		_ = v
		if v != "" {
			m.UpgradePolicy = types.StringValue(v)
		} else {
			m.UpgradePolicy = types.StringNull()
		}
	}
}
