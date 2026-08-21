package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &ToolEMailResource{}
	_ resource.ResourceWithImportState = &ToolEMailResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolEMailResource struct {
	reg *client.Registry
}

type ToolEMailModel struct {
	ID                      types.String `tfsdk:"id"`
	CertificateVerification types.String `tfsdk:"certificate_verification"`
	From                    types.String `tfsdk:"from"`
	Password                types.String `tfsdk:"password"`
	Port                    types.Int64  `tfsdk:"port"`
	Server                  types.String `tfsdk:"server"`
	TLS                     types.String `tfsdk:"tls"`
	User                    types.String `tfsdk:"user"`
	Vrf                     types.String `tfsdk:"vrf"`
	Router                  types.String `tfsdk:"router"`
}

func NewToolEMailResource() resource.Resource { return &ToolEMailResource{} }

func (r *ToolEMailResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_e_mail"
}

func (r *ToolEMailResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolEMailResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/e-mail`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"certificate_verification": schema.StringAttribute{Optional: true, Computed: true,
				Description: "How the SMTP server's certificate is checked: `no`, `yes`, or `yes-without-crl` (verify the chain but skip the CRL check).",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "yes-without-crl"}...)},
			},
			"from": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"password": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true,
				Description: "",
			},
			"port": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"server": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"tls": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Transport security: `no`, `starttls` (upgrade a plain connection via STARTTLS), or `yes` (implicit TLS).",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "starttls", "yes"}...)},
			},
			"user": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"vrf": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *ToolEMailResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolEMailModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolEMailUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolEMailResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolEMailModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolEMailModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolEMailUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolEMailResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolEMailModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/e-mail")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/e-mail failed", err.Error())
		return
	}
	toolEMailApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/e-mail", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolEMailResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolEMailResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/e-mail" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/e-mail", types.StringValue(routerName))))...)
}

func toolEMailUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolEMailModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CertificateVerification.IsNull() || plan.CertificateVerification.IsUnknown()) && (state == nil || !plan.CertificateVerification.Equal(state.CertificateVerification)) {
		body["certificate-verification"] = plan.CertificateVerification.ValueString()
	}
	if !(plan.From.IsNull() || plan.From.IsUnknown()) && (state == nil || !plan.From.Equal(state.From)) {
		body["from"] = plan.From.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) && (state == nil || !plan.Password.Equal(state.Password)) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) && (state == nil || !plan.Port.Equal(state.Port)) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.Server.IsNull() || plan.Server.IsUnknown()) && (state == nil || !plan.Server.Equal(state.Server)) {
		body["server"] = plan.Server.ValueString()
	}
	if !(plan.TLS.IsNull() || plan.TLS.IsUnknown()) && (state == nil || !plan.TLS.Equal(state.TLS)) {
		body["tls"] = plan.TLS.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) && (state == nil || !plan.User.Equal(state.User)) {
		body["user"] = plan.User.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) && (state == nil || !plan.Vrf.Equal(state.Vrf)) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/tool/e-mail", body)
	if err != nil {
		diags.AddError("Upsert /tool/e-mail failed", err.Error())
		return
	}
	toolEMailApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/e-mail", plan.Router))
}

func toolEMailApply(ctx context.Context, obj client.Object, m *ToolEMailModel) {
	_ = ctx
	if v, ok := obj["certificate-verification"]; ok {
		_ = v
		if v != "" {
			m.CertificateVerification = types.StringValue(v)
		} else {
			m.CertificateVerification = types.StringNull()
		}
	}
	if v, ok := obj["from"]; ok {
		_ = v
		if v != "" {
			m.From = types.StringValue(v)
		} else {
			m.From = types.StringNull()
		}
	}
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	}
	if v, ok := obj["server"]; ok {
		_ = v
		if v != "" {
			m.Server = types.StringValue(v)
		} else {
			m.Server = types.StringNull()
		}
	}
	if v, ok := obj["tls"]; ok {
		_ = v
		if v != "" {
			m.TLS = types.StringValue(v)
		} else {
			m.TLS = types.StringNull()
		}
	}
	if v, ok := obj["user"]; ok {
		_ = v
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
