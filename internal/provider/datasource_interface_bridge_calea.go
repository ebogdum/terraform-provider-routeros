package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ datasource.DataSource = &InterfaceBridgeCaleaDataSource{}
	_                       = fmt.Sprintf
)

type InterfaceBridgeCaleaDataSource struct{ reg *client.Registry }

type InterfaceBridgeCaleaDSModel struct {
	Router  types.String `tfsdk:"router"`
	Filter  types.Map    `tfsdk:"filter"`
	Records types.List   `tfsdk:"records"`
}

func NewInterfaceBridgeCaleaDataSource() datasource.DataSource {
	return &InterfaceBridgeCaleaDataSource{}
}

func (d *InterfaceBridgeCaleaDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge_calea"
}

func (d *InterfaceBridgeCaleaDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		d.reg = reg
	}
}

func (d *InterfaceBridgeCaleaDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists records from /interface/bridge/calea.",
		Attributes: map[string]schema.Attribute{
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router. Omit to use the default.",
			},
			"filter": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Equality filters (?k=v).",
			},
			"records": schema.ListAttribute{
				Computed:    true,
				ElementType: types.MapType{ElemType: types.StringType},
				Description: "Matching records as flat string maps (RouterOS wire form).",
			},
		},
	}
}

func (d *InterfaceBridgeCaleaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m InterfaceBridgeCaleaDSModel
	if diags := req.Config.Get(ctx, &m); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(d.reg, m.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	opts := []client.QueryOption{}
	if !m.Filter.IsNull() && !m.Filter.IsUnknown() {
		filt := map[string]string{}
		m.Filter.ElementsAs(ctx, &filt, false)
		for k, v := range filt {
			opts = append(opts, client.WithFilter(k, v))
		}
	}
	rows, err := c.List(ctx, "/interface/bridge/calea", opts...)
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/bridge/calea failed", err.Error())
		return
	}
	recs, diags := dsRowsToList(ctx, rows)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}
	m.Records = recs
	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}
