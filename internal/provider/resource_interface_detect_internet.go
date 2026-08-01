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
	_ resource.Resource                = &InterfaceDetectInternetResource{}
	_ resource.ResourceWithImportState = &InterfaceDetectInternetResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type InterfaceDetectInternetResource struct {
	reg *client.Registry
}

type InterfaceDetectInternetModel struct {
	ID                    types.String `tfsdk:"id"`
	DetectInterfaceList   types.String `tfsdk:"detect_interface_list"`
	InternetInterfaceList types.String `tfsdk:"internet_interface_list"`
	LanInterfaceList      types.String `tfsdk:"lan_interface_list"`
	RequestInterval       types.String `tfsdk:"request_interval"`
	WanInterfaceList      types.String `tfsdk:"wan_interface_list"`
	Router                types.String `tfsdk:"router"`
}

func NewInterfaceDetectInternetResource() resource.Resource {
	return &InterfaceDetectInternetResource{}
}

func (r *InterfaceDetectInternetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_detect_internet"
}

func (r *InterfaceDetectInternetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceDetectInternetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/detect-internet`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"detect_interface_list": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"internet_interface_list": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"lan_interface_list": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"request_interval": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"wan_interface_list": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceDetectInternetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceDetectInternetModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceDetectInternetUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceDetectInternetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceDetectInternetModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state InterfaceDetectInternetModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceDetectInternetUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceDetectInternetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceDetectInternetModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/detect-internet")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/detect-internet failed", err.Error())
		return
	}
	interfaceDetectInternetApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/detect-internet", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceDetectInternetResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfaceDetectInternetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/detect-internet" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/detect-internet", types.StringValue(routerName))))...)
}

func interfaceDetectInternetUpsert(ctx context.Context, reg *client.Registry, plan, state *InterfaceDetectInternetModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DetectInterfaceList.IsNull() || plan.DetectInterfaceList.IsUnknown()) && (state == nil || !plan.DetectInterfaceList.Equal(state.DetectInterfaceList)) {
		body["detect-interface-list"] = plan.DetectInterfaceList.ValueString()
	}
	if !(plan.InternetInterfaceList.IsNull() || plan.InternetInterfaceList.IsUnknown()) && (state == nil || !plan.InternetInterfaceList.Equal(state.InternetInterfaceList)) {
		body["internet-interface-list"] = plan.InternetInterfaceList.ValueString()
	}
	if !(plan.LanInterfaceList.IsNull() || plan.LanInterfaceList.IsUnknown()) && (state == nil || !plan.LanInterfaceList.Equal(state.LanInterfaceList)) {
		body["lan-interface-list"] = plan.LanInterfaceList.ValueString()
	}
	if !(plan.RequestInterval.IsNull() || plan.RequestInterval.IsUnknown()) && (state == nil || !plan.RequestInterval.Equal(state.RequestInterval)) {
		body["request-interval"] = plan.RequestInterval.ValueString()
	}
	if !(plan.WanInterfaceList.IsNull() || plan.WanInterfaceList.IsUnknown()) && (state == nil || !plan.WanInterfaceList.Equal(state.WanInterfaceList)) {
		body["wan-interface-list"] = plan.WanInterfaceList.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/interface/detect-internet", body)
	if err != nil {
		diags.AddError("Upsert /interface/detect-internet failed", err.Error())
		return
	}
	interfaceDetectInternetApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/detect-internet", plan.Router))
}

func interfaceDetectInternetApply(ctx context.Context, obj client.Object, m *InterfaceDetectInternetModel) {
	_ = ctx
	if v, ok := obj["detect-interface-list"]; ok {
		_ = v
		if v != "" {
			m.DetectInterfaceList = types.StringValue(v)
		} else {
			m.DetectInterfaceList = types.StringNull()
		}
	}
	if v, ok := obj["internet-interface-list"]; ok {
		_ = v
		if v != "" {
			m.InternetInterfaceList = types.StringValue(v)
		} else {
			m.InternetInterfaceList = types.StringNull()
		}
	}
	if v, ok := obj["lan-interface-list"]; ok {
		_ = v
		if v != "" {
			m.LanInterfaceList = types.StringValue(v)
		} else {
			m.LanInterfaceList = types.StringNull()
		}
	}
	if v, ok := obj["request-interval"]; ok {
		_ = v
		if v != "" {
			m.RequestInterval = types.StringValue(v)
		} else {
			m.RequestInterval = types.StringNull()
		}
	}
	if v, ok := obj["wan-interface-list"]; ok {
		_ = v
		if v != "" {
			m.WanInterfaceList = types.StringValue(v)
		} else {
			m.WanInterfaceList = types.StringNull()
		}
	}
}
