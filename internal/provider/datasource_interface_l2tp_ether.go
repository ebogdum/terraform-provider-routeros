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
	_ datasource.DataSource = &InterfaceL2TPEtherDataSource{}
	_                       = fmt.Sprintf
)

type InterfaceL2TPEtherDataSource struct{ reg *client.Registry }

type InterfaceL2TPEtherDSModel struct {
	Router  types.String `tfsdk:"router"`
	Filter  types.Map    `tfsdk:"filter"`
	Records types.List   `tfsdk:"records"`
}

func NewInterfaceL2TPEtherDataSource() datasource.DataSource { return &InterfaceL2TPEtherDataSource{} }

func (d *InterfaceL2TPEtherDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_l2tp_ether"
}

func (d *InterfaceL2TPEtherDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		d.reg = reg
	}
}

func (d *InterfaceL2TPEtherDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists records from /interface/l2tp-ether.",
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

func (d *InterfaceL2TPEtherDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m InterfaceL2TPEtherDSModel
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
	rows, err := c.List(ctx, "/interface/l2tp-ether", opts...)
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/l2tp-ether failed", err.Error())
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
