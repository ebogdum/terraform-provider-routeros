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
	_ datasource.DataSource = &IPReverseProxyDataSource{}
	_                       = fmt.Sprintf
)

type IPReverseProxyDataSource struct{ reg *client.Registry }

type IPReverseProxyDSModel struct {
	Router  types.String `tfsdk:"router"`
	Filter  types.Map    `tfsdk:"filter"`
	Records types.List   `tfsdk:"records"`
}

func NewIPReverseProxyDataSource() datasource.DataSource { return &IPReverseProxyDataSource{} }

func (d *IPReverseProxyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_reverse_proxy"
}

func (d *IPReverseProxyDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		d.reg = reg
	}
}

func (d *IPReverseProxyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists records from /ip/reverse-proxy.",
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

func (d *IPReverseProxyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m IPReverseProxyDSModel
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
	rows, err := c.List(ctx, "/ip/reverse-proxy", opts...)
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/reverse-proxy failed", err.Error())
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
