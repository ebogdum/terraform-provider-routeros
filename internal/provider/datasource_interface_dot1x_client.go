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
	_ datasource.DataSource = &InterfaceDot1xClientDataSource{}
	_                       = fmt.Sprintf
)

type InterfaceDot1xClientDataSource struct{ reg *client.Registry }

type InterfaceDot1xClientDSModel struct {
	Router   types.String `tfsdk:"router"`
	Filter   types.Map    `tfsdk:"filter"`
	Proplist types.List   `tfsdk:"proplist"`
	Records  types.List   `tfsdk:"records"`
}

func NewInterfaceDot1xClientDataSource() datasource.DataSource {
	return &InterfaceDot1xClientDataSource{}
}

func (d *InterfaceDot1xClientDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_dot1x_client"
}

func (d *InterfaceDot1xClientDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		d.reg = reg
	}
}

func (d *InterfaceDot1xClientDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists records from /interface/dot1x/client.",
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
			"proplist": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Columns to return. Omit to return every column, including any secret this menu holds.",
			},
			"records": schema.ListAttribute{
				Computed:    true,
				Sensitive:   true,
				ElementType: types.MapType{ElemType: types.StringType},
				Description: "Matching records as flat string maps (RouterOS wire form).",
			},
		},
	}
}

func (d *InterfaceDot1xClientDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m InterfaceDot1xClientDSModel
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
		resp.Diagnostics.Append(m.Filter.ElementsAs(ctx, &filt, false)...)
		for k, v := range filt {
			opts = append(opts, client.WithFilter(k, v))
		}
	}
	if !m.Proplist.IsNull() && !m.Proplist.IsUnknown() {
		var props []string
		resp.Diagnostics.Append(m.Proplist.ElementsAs(ctx, &props, false)...)
		if len(props) > 0 {
			opts = append(opts, client.WithProplist(props...))
		}
	}
	rows, err := c.List(ctx, "/interface/dot1x/client", opts...)
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/dot1x/client failed", err.Error())
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
