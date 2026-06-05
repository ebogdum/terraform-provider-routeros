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
	_ datasource.DataSource = &MPLSLdpAcceptFilterDataSource{}
	_                       = fmt.Sprintf
)

type MPLSLdpAcceptFilterDataSource struct{ reg *client.Registry }

type MPLSLdpAcceptFilterDSModel struct {
	Router  types.String `tfsdk:"router"`
	Filter  types.Map    `tfsdk:"filter"`
	Records types.List   `tfsdk:"records"`
}

func NewMPLSLdpAcceptFilterDataSource() datasource.DataSource {
	return &MPLSLdpAcceptFilterDataSource{}
}

func (d *MPLSLdpAcceptFilterDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_ldp_accept_filter"
}

func (d *MPLSLdpAcceptFilterDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		d.reg = reg
	}
}

func (d *MPLSLdpAcceptFilterDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists records from /mpls/ldp/accept-filter.",
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

func (d *MPLSLdpAcceptFilterDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m MPLSLdpAcceptFilterDSModel
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
	rows, err := c.List(ctx, "/mpls/ldp/accept-filter", opts...)
	if err != nil {
		resp.Diagnostics.AddError("Read /mpls/ldp/accept-filter failed", err.Error())
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
