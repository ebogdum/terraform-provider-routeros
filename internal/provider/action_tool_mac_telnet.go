package provider

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource = &ToolMACTelnetResource{}
	_                   = path.Root
	_                   = fmt.Sprintf
)

type ToolMACTelnetResource struct {
	reg *client.Registry
}

type ToolMACTelnetModel struct {
	ID       types.String `tfsdk:"id"`
	Trigger  types.String `tfsdk:"trigger"`
	TargetID types.String `tfsdk:"target_id"`
	Params   types.Map    `tfsdk:"params"`
	Router   types.String `tfsdk:"router"`
	Output   types.List   `tfsdk:"output"`
}

func NewToolMACTelnetResource() resource.Resource { return &ToolMACTelnetResource{} }

func (r *ToolMACTelnetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_mac_telnet"
}

func (r *ToolMACTelnetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolMACTelnetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Action: /tool/mac-telnet. Recreating the resource (taint or change trigger) re-runs the command.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Hash of the inputs that produced this run.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"trigger": schema.StringAttribute{
				Optional:      true,
				Description:   "Change to force re-execution.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"target_id": schema.StringAttribute{
				Optional:      true,
				Description:   "RouterOS .id of the row this action targets. Required by per-row actions (e.g. /certificate/sign, /interface/reset-counters, /disk/format).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"params": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Extra parameters forwarded to RouterOS verbatim. Keys with dots are allowed. Example: { ca = \"my-ca\", name = \"new-cert\" }.",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
			"output": schema.ListAttribute{
				Computed:    true,
				ElementType: types.MapType{ElemType: types.StringType},
				Description: "Server response rows.",
			},
		},
	}
}

func (r *ToolMACTelnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolMACTelnetModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !plan.TargetID.IsNull() && !plan.TargetID.IsUnknown() && plan.TargetID.ValueString() != "" {
		body[".id"] = plan.TargetID.ValueString()
	}
	if !plan.Params.IsNull() && !plan.Params.IsUnknown() {
		extra := map[string]types.String{}
		if d := plan.Params.ElementsAs(ctx, &extra, false); !d.HasError() {
			for k, v := range extra {
				if !v.IsNull() && !v.IsUnknown() {
					body[k] = v.ValueString()
				}
			}
		}
	}
	rows, err := c.Exec(ctx, "/tool/mac-telnet", "", body)
	if err != nil {
		resp.Diagnostics.AddError("Run /tool/mac-telnet failed", err.Error())
		return
	}

	h := sha1.New()
	keys := make([]string, 0, len(body))
	for k := range body {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(h, "%s=%s\n", k, body[k])
	}
	plan.ID = types.StringValue(hex.EncodeToString(h.Sum(nil)))
	plan.Output = actionRowsToList(ctx, rows)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACTelnetResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
}
func (r *ToolMACTelnetResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
}
func (r *ToolMACTelnetResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
}
