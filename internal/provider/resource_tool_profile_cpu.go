package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &ToolProfileCpuResource{}
	_ resource.ResourceWithImportState = &ToolProfileCpuResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolProfileCpuResource struct {
	reg *client.Registry
}

type ToolProfileCpuModel struct {
	ID                 types.String `tfsdk:"id"`
	Backup             types.String `tfsdk:"backup"`
	Bfd                types.String `tfsdk:"bfd"`
	BGP                types.String `tfsdk:"bgp"`
	Bridging           types.String `tfsdk:"bridging"`
	Btest              types.String `tfsdk:"btest"`
	Certificate        types.String `tfsdk:"certificate"`
	Console            types.String `tfsdk:"console"`
	Container          types.String `tfsdk:"container"`
	DHCP               types.String `tfsdk:"dhcp"`
	Disk               types.String `tfsdk:"disk"`
	DNS                types.String `tfsdk:"dns"`
	Dude               types.String `tfsdk:"dude"`
	EMail              types.String `tfsdk:"e_mail"`
	Encrypting         types.String `tfsdk:"encrypting"`
	Eoip               types.String `tfsdk:"eoip"`
	Ethernet           types.String `tfsdk:"ethernet"`
	Fetcher            types.String `tfsdk:"fetcher"`
	Fileman            types.String `tfsdk:"fileman"`
	Firewall           types.String `tfsdk:"firewall"`
	FirewallMgmt       types.String `tfsdk:"firewall_mgmt"`
	Flash              types.String `tfsdk:"flash"`
	Ftp                types.String `tfsdk:"ftp"`
	Gps                types.String `tfsdk:"gps"`
	Graphing           types.String `tfsdk:"graphing"`
	Gre                types.String `tfsdk:"gre"`
	Health             types.String `tfsdk:"health"`
	Hotspot            types.String `tfsdk:"hotspot"`
	Idle               types.String `tfsdk:"idle"`
	IgmpProxy          types.String `tfsdk:"igmp_proxy"`
	InternetDetect     types.String `tfsdk:"internet_detect"`
	IPPool             types.String `tfsdk:"ip_pool"`
	Ipsec              types.String `tfsdk:"ipsec"`
	Kvm                types.String `tfsdk:"kvm"`
	L7Matcher          types.String `tfsdk:"l7_matcher"`
	Lcd                types.String `tfsdk:"lcd"`
	Ldp                types.String `tfsdk:"ldp"`
	Logging            types.String `tfsdk:"logging"`
	Management         types.String `tfsdk:"management"`
	MPLS               types.String `tfsdk:"mpls"`
	NeighbourDiscovery types.String `tfsdk:"neighbour_discovery"`
	Networking         types.String `tfsdk:"networking"`
	NTP                types.String `tfsdk:"ntp"`
	OSPF               types.String `tfsdk:"ospf"`
	OVPN               types.String `tfsdk:"ovpn"`
	Pim                types.String `tfsdk:"pim"`
	Profiling          types.String `tfsdk:"profiling"`
	QueueMgmt          types.String `tfsdk:"queue_mgmt"`
	Queuing            types.String `tfsdk:"queuing"`
	RADIUS             types.String `tfsdk:"radius"`
	Radv               types.String `tfsdk:"radv"`
	RemoteAccess       types.String `tfsdk:"remote_access"`
	Rip                types.String `tfsdk:"rip"`
	Routing            types.String `tfsdk:"routing"`
	Serial             types.String `tfsdk:"serial"`
	Sniffing           types.String `tfsdk:"sniffing"`
	SNMP               types.String `tfsdk:"snmp"`
	Socks              types.String `tfsdk:"socks"`
	Spi                types.String `tfsdk:"spi"`
	SSH                types.String `tfsdk:"ssh"`
	SSL                types.String `tfsdk:"ssl"`
	Telnet             types.String `tfsdk:"telnet"`
	Tftp               types.String `tfsdk:"tftp"`
	TrafficAccounting  types.String `tfsdk:"traffic_accounting"`
	TrafficFlow        types.String `tfsdk:"traffic_flow"`
	Unclassified       types.String `tfsdk:"unclassified"`
	Upnp               types.String `tfsdk:"upnp"`
	Usb                types.String `tfsdk:"usb"`
	UserManager        types.String `tfsdk:"user_manager"`
	VRRP               types.String `tfsdk:"vrrp"`
	WebProxy           types.String `tfsdk:"web_proxy"`
	Winbox             types.String `tfsdk:"winbox"`
	Wireguard          types.String `tfsdk:"wireguard"`
	Wireless           types.String `tfsdk:"wireless"`
	Www                types.String `tfsdk:"www"`
	Zerotier           types.String `tfsdk:"zerotier"`
	Router             types.String `tfsdk:"router"`
}

func NewToolProfileCpuResource() resource.Resource { return &ToolProfileCpuResource{} }

func (r *ToolProfileCpuResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_profile_cpu"
}

func (r *ToolProfileCpuResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolProfileCpuResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Long-running monitor command, not CRUD",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"backup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Backup service",
			},
			"bfd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "BFD service",
			},
			"bgp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "BGP service",
			},
			"bridging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bridging service",
			},
			"btest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bandwidth test.",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate service",
			},
			"console": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Console",
			},
			"container": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "combined container usage",
			},
			"dhcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DHCP-Server and DHCP-Client services",
			},
			"disk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "storage-related services",
			},
			"dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS-related services",
			},
			"dude": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The Dude package services",
			},
			"e_mail": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "e-mail tool",
			},
			"encrypting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "encrypting processes",
			},
			"eoip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "EoIP",
			},
			"ethernet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ethernet-related properties like link speed, auto-negotiation, duplex mode, monitor a transceiver diagnostic information, etc.",
			},
			"fetcher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fetch tool",
			},
			"fileman": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File manager",
			},
			"firewall": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firewall-related processes",
			},
			"firewall_mgmt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Firewall Management: Filtering, NAT, Mangle",
			},
			"flash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "storage-related services",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FTP Service",
			},
			"gps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "GPS Service",
			},
			"graphing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Graphing tool",
			},
			"gre": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "GRE",
			},
			"health": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "system monitoring, workd health",
			},
			"hotspot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hotspot service",
			},
			"idle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free CPU resources",
			},
			"igmp_proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IGMP Proxy service",
			},
			"internet_detect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Detect Internet tool",
			},
			"ip_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP Pool service",
			},
			"ipsec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPsec service: xfrm -\u00a0 set of statistics showing numbers of packets dropped by the transformation code and why.\u00a0 drivers/crypto - drivers that provide access to the hardware cryptographic accelerators. ipsec - processes that relate to the Internet Key Exchange (IKE) protocols, Authentication Header (AH),\u00a0Encapsulating Security Payload (ESP).",
			},
			"kvm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "KVM virtual machine functionality",
			},
			"l7_matcher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 matcher",
			},
			"lcd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "LCD Interfaces system",
			},
			"ldp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label Distribution Protocol (LDP)",
			},
			"logging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Logging system",
			},
			"management": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "different subsystems: scheduler, networking, file management, etc.",
			},
			"mpls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MPLS-related features",
			},
			"neighbour_discovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Neighbour discovery service",
			},
			"networking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "common set of services included in the networking",
			},
			"ntp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "NTP service",
			},
			"ospf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "OSPF service",
			},
			"ovpn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "OVPN service",
			},
			"pim": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol Independent Multicast",
			},
			"profiling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profiler service",
			},
			"queue_mgmt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Queues: Simple queues, Queue tree, Queue types",
			},
			"queuing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Intermediate Queuing",
			},
			"radius": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS service",
			},
			"radv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 radv daemon log messages service",
			},
			"remote_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "accessing the device directly without logging into RouterOS",
			},
			"rip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Routing Information Protocol",
			},
			"routing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Routing-related services",
			},
			"serial": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "serial console and terminal tool",
			},
			"sniffing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "packet Sniffer tool",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SNMP",
			},
			"socks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Socket Secure",
			},
			"spi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "storage-related services",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSH Server",
			},
			"ssl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Telnet service",
			},
			"tftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TFTP service",
			},
			"traffic_accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Traffic-Flow log system",
			},
			"traffic_flow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Traffic-Flow system",
			},
			"unclassified": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "processes or services that are not defined by this classifier",
			},
			"upnp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "UPnP protocol",
			},
			"usb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "USB features",
			},
			"user_manager": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User Manager service",
			},
			"vrrp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VRRP",
			},
			"web_proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web Proxy",
			},
			"winbox": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Winbox",
			},
			"wireguard": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wireguard",
			},
			"wireless": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "common set of services using Wireless systems",
			},
			"www": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Webfig HTTP service",
			},
			"zerotier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ZeroTier",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolProfileCpuResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolProfileCpuModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Backup.IsNull() || plan.Backup.IsUnknown()) {
		body["backup"] = plan.Backup.ValueString()
	}
	if !(plan.Bfd.IsNull() || plan.Bfd.IsUnknown()) {
		body["bfd"] = plan.Bfd.ValueString()
	}
	if !(plan.BGP.IsNull() || plan.BGP.IsUnknown()) {
		body["bgp"] = plan.BGP.ValueString()
	}
	if !(plan.Bridging.IsNull() || plan.Bridging.IsUnknown()) {
		body["bridging"] = plan.Bridging.ValueString()
	}
	if !(plan.Btest.IsNull() || plan.Btest.IsUnknown()) {
		body["btest"] = plan.Btest.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Console.IsNull() || plan.Console.IsUnknown()) {
		body["console"] = plan.Console.ValueString()
	}
	if !(plan.Container.IsNull() || plan.Container.IsUnknown()) {
		body["container"] = plan.Container.ValueString()
	}
	if !(plan.DHCP.IsNull() || plan.DHCP.IsUnknown()) {
		body["dhcp"] = plan.DHCP.ValueString()
	}
	if !(plan.Disk.IsNull() || plan.Disk.IsUnknown()) {
		body["disk"] = plan.Disk.ValueString()
	}
	if !(plan.DNS.IsNull() || plan.DNS.IsUnknown()) {
		body["dns"] = plan.DNS.ValueString()
	}
	if !(plan.Dude.IsNull() || plan.Dude.IsUnknown()) {
		body["dude"] = plan.Dude.ValueString()
	}
	if !(plan.EMail.IsNull() || plan.EMail.IsUnknown()) {
		body["e-mail"] = plan.EMail.ValueString()
	}
	if !(plan.Encrypting.IsNull() || plan.Encrypting.IsUnknown()) {
		body["encrypting"] = plan.Encrypting.ValueString()
	}
	if !(plan.Eoip.IsNull() || plan.Eoip.IsUnknown()) {
		body["eoip"] = plan.Eoip.ValueString()
	}
	if !(plan.Ethernet.IsNull() || plan.Ethernet.IsUnknown()) {
		body["ethernet"] = plan.Ethernet.ValueString()
	}
	if !(plan.Fetcher.IsNull() || plan.Fetcher.IsUnknown()) {
		body["fetcher"] = plan.Fetcher.ValueString()
	}
	if !(plan.Fileman.IsNull() || plan.Fileman.IsUnknown()) {
		body["fileman"] = plan.Fileman.ValueString()
	}
	if !(plan.Firewall.IsNull() || plan.Firewall.IsUnknown()) {
		body["firewall"] = plan.Firewall.ValueString()
	}
	if !(plan.FirewallMgmt.IsNull() || plan.FirewallMgmt.IsUnknown()) {
		body["firewall-mgmt"] = plan.FirewallMgmt.ValueString()
	}
	if !(plan.Flash.IsNull() || plan.Flash.IsUnknown()) {
		body["flash"] = plan.Flash.ValueString()
	}
	if !(plan.Ftp.IsNull() || plan.Ftp.IsUnknown()) {
		body["ftp"] = plan.Ftp.ValueString()
	}
	if !(plan.Gps.IsNull() || plan.Gps.IsUnknown()) {
		body["gps"] = plan.Gps.ValueString()
	}
	if !(plan.Graphing.IsNull() || plan.Graphing.IsUnknown()) {
		body["graphing"] = plan.Graphing.ValueString()
	}
	if !(plan.Gre.IsNull() || plan.Gre.IsUnknown()) {
		body["gre"] = plan.Gre.ValueString()
	}
	if !(plan.Health.IsNull() || plan.Health.IsUnknown()) {
		body["health"] = plan.Health.ValueString()
	}
	if !(plan.Hotspot.IsNull() || plan.Hotspot.IsUnknown()) {
		body["hotspot"] = plan.Hotspot.ValueString()
	}
	if !(plan.Idle.IsNull() || plan.Idle.IsUnknown()) {
		body["idle"] = plan.Idle.ValueString()
	}
	if !(plan.IgmpProxy.IsNull() || plan.IgmpProxy.IsUnknown()) {
		body["igmp-proxy"] = plan.IgmpProxy.ValueString()
	}
	if !(plan.InternetDetect.IsNull() || plan.InternetDetect.IsUnknown()) {
		body["internet-detect"] = plan.InternetDetect.ValueString()
	}
	if !(plan.IPPool.IsNull() || plan.IPPool.IsUnknown()) {
		body["ip-pool"] = plan.IPPool.ValueString()
	}
	if !(plan.Ipsec.IsNull() || plan.Ipsec.IsUnknown()) {
		body["ipsec"] = plan.Ipsec.ValueString()
	}
	if !(plan.Kvm.IsNull() || plan.Kvm.IsUnknown()) {
		body["kvm"] = plan.Kvm.ValueString()
	}
	if !(plan.L7Matcher.IsNull() || plan.L7Matcher.IsUnknown()) {
		body["l7-matcher"] = plan.L7Matcher.ValueString()
	}
	if !(plan.Lcd.IsNull() || plan.Lcd.IsUnknown()) {
		body["lcd"] = plan.Lcd.ValueString()
	}
	if !(plan.Ldp.IsNull() || plan.Ldp.IsUnknown()) {
		body["ldp"] = plan.Ldp.ValueString()
	}
	if !(plan.Logging.IsNull() || plan.Logging.IsUnknown()) {
		body["logging"] = plan.Logging.ValueString()
	}
	if !(plan.Management.IsNull() || plan.Management.IsUnknown()) {
		body["management"] = plan.Management.ValueString()
	}
	if !(plan.MPLS.IsNull() || plan.MPLS.IsUnknown()) {
		body["mpls"] = plan.MPLS.ValueString()
	}
	if !(plan.NeighbourDiscovery.IsNull() || plan.NeighbourDiscovery.IsUnknown()) {
		body["neighbour-discovery"] = plan.NeighbourDiscovery.ValueString()
	}
	if !(plan.Networking.IsNull() || plan.Networking.IsUnknown()) {
		body["networking"] = plan.Networking.ValueString()
	}
	if !(plan.NTP.IsNull() || plan.NTP.IsUnknown()) {
		body["ntp"] = plan.NTP.ValueString()
	}
	if !(plan.OSPF.IsNull() || plan.OSPF.IsUnknown()) {
		body["ospf"] = plan.OSPF.ValueString()
	}
	if !(plan.OVPN.IsNull() || plan.OVPN.IsUnknown()) {
		body["ovpn"] = plan.OVPN.ValueString()
	}
	if !(plan.Pim.IsNull() || plan.Pim.IsUnknown()) {
		body["pim"] = plan.Pim.ValueString()
	}
	if !(plan.Profiling.IsNull() || plan.Profiling.IsUnknown()) {
		body["profiling"] = plan.Profiling.ValueString()
	}
	if !(plan.QueueMgmt.IsNull() || plan.QueueMgmt.IsUnknown()) {
		body["queue-mgmt"] = plan.QueueMgmt.ValueString()
	}
	if !(plan.Queuing.IsNull() || plan.Queuing.IsUnknown()) {
		body["queuing"] = plan.Queuing.ValueString()
	}
	if !(plan.RADIUS.IsNull() || plan.RADIUS.IsUnknown()) {
		body["radius"] = plan.RADIUS.ValueString()
	}
	if !(plan.Radv.IsNull() || plan.Radv.IsUnknown()) {
		body["radv"] = plan.Radv.ValueString()
	}
	if !(plan.RemoteAccess.IsNull() || plan.RemoteAccess.IsUnknown()) {
		body["remote-access"] = plan.RemoteAccess.ValueString()
	}
	if !(plan.Rip.IsNull() || plan.Rip.IsUnknown()) {
		body["rip"] = plan.Rip.ValueString()
	}
	if !(plan.Routing.IsNull() || plan.Routing.IsUnknown()) {
		body["routing"] = plan.Routing.ValueString()
	}
	if !(plan.Serial.IsNull() || plan.Serial.IsUnknown()) {
		body["serial"] = plan.Serial.ValueString()
	}
	if !(plan.Sniffing.IsNull() || plan.Sniffing.IsUnknown()) {
		body["sniffing"] = plan.Sniffing.ValueString()
	}
	if !(plan.SNMP.IsNull() || plan.SNMP.IsUnknown()) {
		body["snmp"] = plan.SNMP.ValueString()
	}
	if !(plan.Socks.IsNull() || plan.Socks.IsUnknown()) {
		body["socks"] = plan.Socks.ValueString()
	}
	if !(plan.Spi.IsNull() || plan.Spi.IsUnknown()) {
		body["spi"] = plan.Spi.ValueString()
	}
	if !(plan.SSH.IsNull() || plan.SSH.IsUnknown()) {
		body["ssh"] = plan.SSH.ValueString()
	}
	if !(plan.SSL.IsNull() || plan.SSL.IsUnknown()) {
		body["ssl"] = plan.SSL.ValueString()
	}
	if !(plan.Telnet.IsNull() || plan.Telnet.IsUnknown()) {
		body["telnet"] = plan.Telnet.ValueString()
	}
	if !(plan.Tftp.IsNull() || plan.Tftp.IsUnknown()) {
		body["tftp"] = plan.Tftp.ValueString()
	}
	if !(plan.TrafficAccounting.IsNull() || plan.TrafficAccounting.IsUnknown()) {
		body["traffic-accounting"] = plan.TrafficAccounting.ValueString()
	}
	if !(plan.TrafficFlow.IsNull() || plan.TrafficFlow.IsUnknown()) {
		body["traffic-flow"] = plan.TrafficFlow.ValueString()
	}
	if !(plan.Unclassified.IsNull() || plan.Unclassified.IsUnknown()) {
		body["unclassified"] = plan.Unclassified.ValueString()
	}
	if !(plan.Upnp.IsNull() || plan.Upnp.IsUnknown()) {
		body["upnp"] = plan.Upnp.ValueString()
	}
	if !(plan.Usb.IsNull() || plan.Usb.IsUnknown()) {
		body["usb"] = plan.Usb.ValueString()
	}
	if !(plan.UserManager.IsNull() || plan.UserManager.IsUnknown()) {
		body["user-manager"] = plan.UserManager.ValueString()
	}
	if !(plan.VRRP.IsNull() || plan.VRRP.IsUnknown()) {
		body["vrrp"] = plan.VRRP.ValueString()
	}
	if !(plan.WebProxy.IsNull() || plan.WebProxy.IsUnknown()) {
		body["web-proxy"] = plan.WebProxy.ValueString()
	}
	if !(plan.Winbox.IsNull() || plan.Winbox.IsUnknown()) {
		body["winbox"] = plan.Winbox.ValueString()
	}
	if !(plan.Wireguard.IsNull() || plan.Wireguard.IsUnknown()) {
		body["wireguard"] = plan.Wireguard.ValueString()
	}
	if !(plan.Wireless.IsNull() || plan.Wireless.IsUnknown()) {
		body["wireless"] = plan.Wireless.ValueString()
	}
	if !(plan.Www.IsNull() || plan.Www.IsUnknown()) {
		body["www"] = plan.Www.ValueString()
	}
	if !(plan.Zerotier.IsNull() || plan.Zerotier.IsUnknown()) {
		body["zerotier"] = plan.Zerotier.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/profile/cpu", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/profile/cpu failed", err.Error())
		return
	}
	toolProfileCpuApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolProfileCpuResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolProfileCpuModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/profile/cpu", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/profile/cpu failed", err.Error())
		return
	}
	toolProfileCpuApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolProfileCpuResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolProfileCpuModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !plan.Backup.Equal(state.Backup) {
		body["backup"] = plan.Backup.ValueString()
	}
	if !plan.Bfd.Equal(state.Bfd) {
		body["bfd"] = plan.Bfd.ValueString()
	}
	if !plan.BGP.Equal(state.BGP) {
		body["bgp"] = plan.BGP.ValueString()
	}
	if !plan.Bridging.Equal(state.Bridging) {
		body["bridging"] = plan.Bridging.ValueString()
	}
	if !plan.Btest.Equal(state.Btest) {
		body["btest"] = plan.Btest.ValueString()
	}
	if !plan.Certificate.Equal(state.Certificate) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Console.Equal(state.Console) {
		body["console"] = plan.Console.ValueString()
	}
	if !plan.Container.Equal(state.Container) {
		body["container"] = plan.Container.ValueString()
	}
	if !plan.DHCP.Equal(state.DHCP) {
		body["dhcp"] = plan.DHCP.ValueString()
	}
	if !plan.Disk.Equal(state.Disk) {
		body["disk"] = plan.Disk.ValueString()
	}
	if !plan.DNS.Equal(state.DNS) {
		body["dns"] = plan.DNS.ValueString()
	}
	if !plan.Dude.Equal(state.Dude) {
		body["dude"] = plan.Dude.ValueString()
	}
	if !plan.EMail.Equal(state.EMail) {
		body["e-mail"] = plan.EMail.ValueString()
	}
	if !plan.Encrypting.Equal(state.Encrypting) {
		body["encrypting"] = plan.Encrypting.ValueString()
	}
	if !plan.Eoip.Equal(state.Eoip) {
		body["eoip"] = plan.Eoip.ValueString()
	}
	if !plan.Ethernet.Equal(state.Ethernet) {
		body["ethernet"] = plan.Ethernet.ValueString()
	}
	if !plan.Fetcher.Equal(state.Fetcher) {
		body["fetcher"] = plan.Fetcher.ValueString()
	}
	if !plan.Fileman.Equal(state.Fileman) {
		body["fileman"] = plan.Fileman.ValueString()
	}
	if !plan.Firewall.Equal(state.Firewall) {
		body["firewall"] = plan.Firewall.ValueString()
	}
	if !plan.FirewallMgmt.Equal(state.FirewallMgmt) {
		body["firewall-mgmt"] = plan.FirewallMgmt.ValueString()
	}
	if !plan.Flash.Equal(state.Flash) {
		body["flash"] = plan.Flash.ValueString()
	}
	if !plan.Ftp.Equal(state.Ftp) {
		body["ftp"] = plan.Ftp.ValueString()
	}
	if !plan.Gps.Equal(state.Gps) {
		body["gps"] = plan.Gps.ValueString()
	}
	if !plan.Graphing.Equal(state.Graphing) {
		body["graphing"] = plan.Graphing.ValueString()
	}
	if !plan.Gre.Equal(state.Gre) {
		body["gre"] = plan.Gre.ValueString()
	}
	if !plan.Health.Equal(state.Health) {
		body["health"] = plan.Health.ValueString()
	}
	if !plan.Hotspot.Equal(state.Hotspot) {
		body["hotspot"] = plan.Hotspot.ValueString()
	}
	if !plan.Idle.Equal(state.Idle) {
		body["idle"] = plan.Idle.ValueString()
	}
	if !plan.IgmpProxy.Equal(state.IgmpProxy) {
		body["igmp-proxy"] = plan.IgmpProxy.ValueString()
	}
	if !plan.InternetDetect.Equal(state.InternetDetect) {
		body["internet-detect"] = plan.InternetDetect.ValueString()
	}
	if !plan.IPPool.Equal(state.IPPool) {
		body["ip-pool"] = plan.IPPool.ValueString()
	}
	if !plan.Ipsec.Equal(state.Ipsec) {
		body["ipsec"] = plan.Ipsec.ValueString()
	}
	if !plan.Kvm.Equal(state.Kvm) {
		body["kvm"] = plan.Kvm.ValueString()
	}
	if !plan.L7Matcher.Equal(state.L7Matcher) {
		body["l7-matcher"] = plan.L7Matcher.ValueString()
	}
	if !plan.Lcd.Equal(state.Lcd) {
		body["lcd"] = plan.Lcd.ValueString()
	}
	if !plan.Ldp.Equal(state.Ldp) {
		body["ldp"] = plan.Ldp.ValueString()
	}
	if !plan.Logging.Equal(state.Logging) {
		body["logging"] = plan.Logging.ValueString()
	}
	if !plan.Management.Equal(state.Management) {
		body["management"] = plan.Management.ValueString()
	}
	if !plan.MPLS.Equal(state.MPLS) {
		body["mpls"] = plan.MPLS.ValueString()
	}
	if !plan.NeighbourDiscovery.Equal(state.NeighbourDiscovery) {
		body["neighbour-discovery"] = plan.NeighbourDiscovery.ValueString()
	}
	if !plan.Networking.Equal(state.Networking) {
		body["networking"] = plan.Networking.ValueString()
	}
	if !plan.NTP.Equal(state.NTP) {
		body["ntp"] = plan.NTP.ValueString()
	}
	if !plan.OSPF.Equal(state.OSPF) {
		body["ospf"] = plan.OSPF.ValueString()
	}
	if !plan.OVPN.Equal(state.OVPN) {
		body["ovpn"] = plan.OVPN.ValueString()
	}
	if !plan.Pim.Equal(state.Pim) {
		body["pim"] = plan.Pim.ValueString()
	}
	if !plan.Profiling.Equal(state.Profiling) {
		body["profiling"] = plan.Profiling.ValueString()
	}
	if !plan.QueueMgmt.Equal(state.QueueMgmt) {
		body["queue-mgmt"] = plan.QueueMgmt.ValueString()
	}
	if !plan.Queuing.Equal(state.Queuing) {
		body["queuing"] = plan.Queuing.ValueString()
	}
	if !plan.RADIUS.Equal(state.RADIUS) {
		body["radius"] = plan.RADIUS.ValueString()
	}
	if !plan.Radv.Equal(state.Radv) {
		body["radv"] = plan.Radv.ValueString()
	}
	if !plan.RemoteAccess.Equal(state.RemoteAccess) {
		body["remote-access"] = plan.RemoteAccess.ValueString()
	}
	if !plan.Rip.Equal(state.Rip) {
		body["rip"] = plan.Rip.ValueString()
	}
	if !plan.Routing.Equal(state.Routing) {
		body["routing"] = plan.Routing.ValueString()
	}
	if !plan.Serial.Equal(state.Serial) {
		body["serial"] = plan.Serial.ValueString()
	}
	if !plan.Sniffing.Equal(state.Sniffing) {
		body["sniffing"] = plan.Sniffing.ValueString()
	}
	if !plan.SNMP.Equal(state.SNMP) {
		body["snmp"] = plan.SNMP.ValueString()
	}
	if !plan.Socks.Equal(state.Socks) {
		body["socks"] = plan.Socks.ValueString()
	}
	if !plan.Spi.Equal(state.Spi) {
		body["spi"] = plan.Spi.ValueString()
	}
	if !plan.SSH.Equal(state.SSH) {
		body["ssh"] = plan.SSH.ValueString()
	}
	if !plan.SSL.Equal(state.SSL) {
		body["ssl"] = plan.SSL.ValueString()
	}
	if !plan.Telnet.Equal(state.Telnet) {
		body["telnet"] = plan.Telnet.ValueString()
	}
	if !plan.Tftp.Equal(state.Tftp) {
		body["tftp"] = plan.Tftp.ValueString()
	}
	if !plan.TrafficAccounting.Equal(state.TrafficAccounting) {
		body["traffic-accounting"] = plan.TrafficAccounting.ValueString()
	}
	if !plan.TrafficFlow.Equal(state.TrafficFlow) {
		body["traffic-flow"] = plan.TrafficFlow.ValueString()
	}
	if !plan.Unclassified.Equal(state.Unclassified) {
		body["unclassified"] = plan.Unclassified.ValueString()
	}
	if !plan.Upnp.Equal(state.Upnp) {
		body["upnp"] = plan.Upnp.ValueString()
	}
	if !plan.Usb.Equal(state.Usb) {
		body["usb"] = plan.Usb.ValueString()
	}
	if !plan.UserManager.Equal(state.UserManager) {
		body["user-manager"] = plan.UserManager.ValueString()
	}
	if !plan.VRRP.Equal(state.VRRP) {
		body["vrrp"] = plan.VRRP.ValueString()
	}
	if !plan.WebProxy.Equal(state.WebProxy) {
		body["web-proxy"] = plan.WebProxy.ValueString()
	}
	if !plan.Winbox.Equal(state.Winbox) {
		body["winbox"] = plan.Winbox.ValueString()
	}
	if !plan.Wireguard.Equal(state.Wireguard) {
		body["wireguard"] = plan.Wireguard.ValueString()
	}
	if !plan.Wireless.Equal(state.Wireless) {
		body["wireless"] = plan.Wireless.ValueString()
	}
	if !plan.Www.Equal(state.Www) {
		body["www"] = plan.Www.ValueString()
	}
	if !plan.Zerotier.Equal(state.Zerotier) {
		body["zerotier"] = plan.Zerotier.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/profile/cpu", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/profile/cpu failed", err.Error())
			return
		}
		toolProfileCpuApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolProfileCpuResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolProfileCpuModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/profile/cpu", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/profile/cpu failed", err.Error())
	}
}

func (r *ToolProfileCpuResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	routerName, id := parseImportID(r.reg, req.ID)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := toolProfileCpuLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/profile/cpu matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolProfileCpuLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolProfileCpuLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/tool/profile/cpu", id)
}

func toolProfileCpuApply(ctx context.Context, obj client.Object, m *ToolProfileCpuModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["backup"]; ok {
		_ = v
		if v != "" {
			m.Backup = types.StringValue(v)
		} else {
			m.Backup = types.StringNull()
		}
	} else {
		m.Backup = types.StringNull()
	}
	if v, ok := obj["bfd"]; ok {
		_ = v
		if v != "" {
			m.Bfd = types.StringValue(v)
		} else {
			m.Bfd = types.StringNull()
		}
	} else {
		m.Bfd = types.StringNull()
	}
	if v, ok := obj["bgp"]; ok {
		_ = v
		if v != "" {
			m.BGP = types.StringValue(v)
		} else {
			m.BGP = types.StringNull()
		}
	} else {
		m.BGP = types.StringNull()
	}
	if v, ok := obj["bridging"]; ok {
		_ = v
		if v != "" {
			m.Bridging = types.StringValue(v)
		} else {
			m.Bridging = types.StringNull()
		}
	} else {
		m.Bridging = types.StringNull()
	}
	if v, ok := obj["btest"]; ok {
		_ = v
		if v != "" {
			m.Btest = types.StringValue(v)
		} else {
			m.Btest = types.StringNull()
		}
	} else {
		m.Btest = types.StringNull()
	}
	if v, ok := obj["certificate"]; ok {
		_ = v
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	} else {
		m.Certificate = types.StringNull()
	}
	if v, ok := obj["console"]; ok {
		_ = v
		if v != "" {
			m.Console = types.StringValue(v)
		} else {
			m.Console = types.StringNull()
		}
	} else {
		m.Console = types.StringNull()
	}
	if v, ok := obj["container"]; ok {
		_ = v
		if v != "" {
			m.Container = types.StringValue(v)
		} else {
			m.Container = types.StringNull()
		}
	} else {
		m.Container = types.StringNull()
	}
	if v, ok := obj["dhcp"]; ok {
		_ = v
		if v != "" {
			m.DHCP = types.StringValue(v)
		} else {
			m.DHCP = types.StringNull()
		}
	} else {
		m.DHCP = types.StringNull()
	}
	if v, ok := obj["disk"]; ok {
		_ = v
		if v != "" {
			m.Disk = types.StringValue(v)
		} else {
			m.Disk = types.StringNull()
		}
	} else {
		m.Disk = types.StringNull()
	}
	if v, ok := obj["dns"]; ok {
		_ = v
		if v != "" {
			m.DNS = types.StringValue(v)
		} else {
			m.DNS = types.StringNull()
		}
	} else {
		m.DNS = types.StringNull()
	}
	if v, ok := obj["dude"]; ok {
		_ = v
		if v != "" {
			m.Dude = types.StringValue(v)
		} else {
			m.Dude = types.StringNull()
		}
	} else {
		m.Dude = types.StringNull()
	}
	if v, ok := obj["e-mail"]; ok {
		_ = v
		if v != "" {
			m.EMail = types.StringValue(v)
		} else {
			m.EMail = types.StringNull()
		}
	} else {
		m.EMail = types.StringNull()
	}
	if v, ok := obj["encrypting"]; ok {
		_ = v
		if v != "" {
			m.Encrypting = types.StringValue(v)
		} else {
			m.Encrypting = types.StringNull()
		}
	} else {
		m.Encrypting = types.StringNull()
	}
	if v, ok := obj["eoip"]; ok {
		_ = v
		if v != "" {
			m.Eoip = types.StringValue(v)
		} else {
			m.Eoip = types.StringNull()
		}
	} else {
		m.Eoip = types.StringNull()
	}
	if v, ok := obj["ethernet"]; ok {
		_ = v
		if v != "" {
			m.Ethernet = types.StringValue(v)
		} else {
			m.Ethernet = types.StringNull()
		}
	} else {
		m.Ethernet = types.StringNull()
	}
	if v, ok := obj["fetcher"]; ok {
		_ = v
		if v != "" {
			m.Fetcher = types.StringValue(v)
		} else {
			m.Fetcher = types.StringNull()
		}
	} else {
		m.Fetcher = types.StringNull()
	}
	if v, ok := obj["fileman"]; ok {
		_ = v
		if v != "" {
			m.Fileman = types.StringValue(v)
		} else {
			m.Fileman = types.StringNull()
		}
	} else {
		m.Fileman = types.StringNull()
	}
	if v, ok := obj["firewall"]; ok {
		_ = v
		if v != "" {
			m.Firewall = types.StringValue(v)
		} else {
			m.Firewall = types.StringNull()
		}
	} else {
		m.Firewall = types.StringNull()
	}
	if v, ok := obj["firewall-mgmt"]; ok {
		_ = v
		if v != "" {
			m.FirewallMgmt = types.StringValue(v)
		} else {
			m.FirewallMgmt = types.StringNull()
		}
	} else {
		m.FirewallMgmt = types.StringNull()
	}
	if v, ok := obj["flash"]; ok {
		_ = v
		if v != "" {
			m.Flash = types.StringValue(v)
		} else {
			m.Flash = types.StringNull()
		}
	} else {
		m.Flash = types.StringNull()
	}
	if v, ok := obj["ftp"]; ok {
		_ = v
		if v != "" {
			m.Ftp = types.StringValue(v)
		} else {
			m.Ftp = types.StringNull()
		}
	} else {
		m.Ftp = types.StringNull()
	}
	if v, ok := obj["gps"]; ok {
		_ = v
		if v != "" {
			m.Gps = types.StringValue(v)
		} else {
			m.Gps = types.StringNull()
		}
	} else {
		m.Gps = types.StringNull()
	}
	if v, ok := obj["graphing"]; ok {
		_ = v
		if v != "" {
			m.Graphing = types.StringValue(v)
		} else {
			m.Graphing = types.StringNull()
		}
	} else {
		m.Graphing = types.StringNull()
	}
	if v, ok := obj["gre"]; ok {
		_ = v
		if v != "" {
			m.Gre = types.StringValue(v)
		} else {
			m.Gre = types.StringNull()
		}
	} else {
		m.Gre = types.StringNull()
	}
	if v, ok := obj["health"]; ok {
		_ = v
		if v != "" {
			m.Health = types.StringValue(v)
		} else {
			m.Health = types.StringNull()
		}
	} else {
		m.Health = types.StringNull()
	}
	if v, ok := obj["hotspot"]; ok {
		_ = v
		if v != "" {
			m.Hotspot = types.StringValue(v)
		} else {
			m.Hotspot = types.StringNull()
		}
	} else {
		m.Hotspot = types.StringNull()
	}
	if v, ok := obj["idle"]; ok {
		_ = v
		if v != "" {
			m.Idle = types.StringValue(v)
		} else {
			m.Idle = types.StringNull()
		}
	} else {
		m.Idle = types.StringNull()
	}
	if v, ok := obj["igmp-proxy"]; ok {
		_ = v
		if v != "" {
			m.IgmpProxy = types.StringValue(v)
		} else {
			m.IgmpProxy = types.StringNull()
		}
	} else {
		m.IgmpProxy = types.StringNull()
	}
	if v, ok := obj["internet-detect"]; ok {
		_ = v
		if v != "" {
			m.InternetDetect = types.StringValue(v)
		} else {
			m.InternetDetect = types.StringNull()
		}
	} else {
		m.InternetDetect = types.StringNull()
	}
	if v, ok := obj["ip-pool"]; ok {
		_ = v
		if v != "" {
			m.IPPool = types.StringValue(v)
		} else {
			m.IPPool = types.StringNull()
		}
	} else {
		m.IPPool = types.StringNull()
	}
	if v, ok := obj["ipsec"]; ok {
		_ = v
		if v != "" {
			m.Ipsec = types.StringValue(v)
		} else {
			m.Ipsec = types.StringNull()
		}
	} else {
		m.Ipsec = types.StringNull()
	}
	if v, ok := obj["kvm"]; ok {
		_ = v
		if v != "" {
			m.Kvm = types.StringValue(v)
		} else {
			m.Kvm = types.StringNull()
		}
	} else {
		m.Kvm = types.StringNull()
	}
	if v, ok := obj["l7-matcher"]; ok {
		_ = v
		if v != "" {
			m.L7Matcher = types.StringValue(v)
		} else {
			m.L7Matcher = types.StringNull()
		}
	} else {
		m.L7Matcher = types.StringNull()
	}
	if v, ok := obj["lcd"]; ok {
		_ = v
		if v != "" {
			m.Lcd = types.StringValue(v)
		} else {
			m.Lcd = types.StringNull()
		}
	} else {
		m.Lcd = types.StringNull()
	}
	if v, ok := obj["ldp"]; ok {
		_ = v
		if v != "" {
			m.Ldp = types.StringValue(v)
		} else {
			m.Ldp = types.StringNull()
		}
	} else {
		m.Ldp = types.StringNull()
	}
	if v, ok := obj["logging"]; ok {
		_ = v
		if v != "" {
			m.Logging = types.StringValue(v)
		} else {
			m.Logging = types.StringNull()
		}
	} else {
		m.Logging = types.StringNull()
	}
	if v, ok := obj["management"]; ok {
		_ = v
		if v != "" {
			m.Management = types.StringValue(v)
		} else {
			m.Management = types.StringNull()
		}
	} else {
		m.Management = types.StringNull()
	}
	if v, ok := obj["mpls"]; ok {
		_ = v
		if v != "" {
			m.MPLS = types.StringValue(v)
		} else {
			m.MPLS = types.StringNull()
		}
	} else {
		m.MPLS = types.StringNull()
	}
	if v, ok := obj["neighbour-discovery"]; ok {
		_ = v
		if v != "" {
			m.NeighbourDiscovery = types.StringValue(v)
		} else {
			m.NeighbourDiscovery = types.StringNull()
		}
	} else {
		m.NeighbourDiscovery = types.StringNull()
	}
	if v, ok := obj["networking"]; ok {
		_ = v
		if v != "" {
			m.Networking = types.StringValue(v)
		} else {
			m.Networking = types.StringNull()
		}
	} else {
		m.Networking = types.StringNull()
	}
	if v, ok := obj["ntp"]; ok {
		_ = v
		if v != "" {
			m.NTP = types.StringValue(v)
		} else {
			m.NTP = types.StringNull()
		}
	} else {
		m.NTP = types.StringNull()
	}
	if v, ok := obj["ospf"]; ok {
		_ = v
		if v != "" {
			m.OSPF = types.StringValue(v)
		} else {
			m.OSPF = types.StringNull()
		}
	} else {
		m.OSPF = types.StringNull()
	}
	if v, ok := obj["ovpn"]; ok {
		_ = v
		if v != "" {
			m.OVPN = types.StringValue(v)
		} else {
			m.OVPN = types.StringNull()
		}
	} else {
		m.OVPN = types.StringNull()
	}
	if v, ok := obj["pim"]; ok {
		_ = v
		if v != "" {
			m.Pim = types.StringValue(v)
		} else {
			m.Pim = types.StringNull()
		}
	} else {
		m.Pim = types.StringNull()
	}
	if v, ok := obj["profiling"]; ok {
		_ = v
		if v != "" {
			m.Profiling = types.StringValue(v)
		} else {
			m.Profiling = types.StringNull()
		}
	} else {
		m.Profiling = types.StringNull()
	}
	if v, ok := obj["queue-mgmt"]; ok {
		_ = v
		if v != "" {
			m.QueueMgmt = types.StringValue(v)
		} else {
			m.QueueMgmt = types.StringNull()
		}
	} else {
		m.QueueMgmt = types.StringNull()
	}
	if v, ok := obj["queuing"]; ok {
		_ = v
		if v != "" {
			m.Queuing = types.StringValue(v)
		} else {
			m.Queuing = types.StringNull()
		}
	} else {
		m.Queuing = types.StringNull()
	}
	if v, ok := obj["radius"]; ok {
		_ = v
		if v != "" {
			m.RADIUS = types.StringValue(v)
		} else {
			m.RADIUS = types.StringNull()
		}
	} else {
		m.RADIUS = types.StringNull()
	}
	if v, ok := obj["radv"]; ok {
		_ = v
		if v != "" {
			m.Radv = types.StringValue(v)
		} else {
			m.Radv = types.StringNull()
		}
	} else {
		m.Radv = types.StringNull()
	}
	if v, ok := obj["remote-access"]; ok {
		_ = v
		if v != "" {
			m.RemoteAccess = types.StringValue(v)
		} else {
			m.RemoteAccess = types.StringNull()
		}
	} else {
		m.RemoteAccess = types.StringNull()
	}
	if v, ok := obj["rip"]; ok {
		_ = v
		if v != "" {
			m.Rip = types.StringValue(v)
		} else {
			m.Rip = types.StringNull()
		}
	} else {
		m.Rip = types.StringNull()
	}
	if v, ok := obj["routing"]; ok {
		_ = v
		if v != "" {
			m.Routing = types.StringValue(v)
		} else {
			m.Routing = types.StringNull()
		}
	} else {
		m.Routing = types.StringNull()
	}
	if v, ok := obj["serial"]; ok {
		_ = v
		if v != "" {
			m.Serial = types.StringValue(v)
		} else {
			m.Serial = types.StringNull()
		}
	} else {
		m.Serial = types.StringNull()
	}
	if v, ok := obj["sniffing"]; ok {
		_ = v
		if v != "" {
			m.Sniffing = types.StringValue(v)
		} else {
			m.Sniffing = types.StringNull()
		}
	} else {
		m.Sniffing = types.StringNull()
	}
	if v, ok := obj["snmp"]; ok {
		_ = v
		if v != "" {
			m.SNMP = types.StringValue(v)
		} else {
			m.SNMP = types.StringNull()
		}
	} else {
		m.SNMP = types.StringNull()
	}
	if v, ok := obj["socks"]; ok {
		_ = v
		if v != "" {
			m.Socks = types.StringValue(v)
		} else {
			m.Socks = types.StringNull()
		}
	} else {
		m.Socks = types.StringNull()
	}
	if v, ok := obj["spi"]; ok {
		_ = v
		if v != "" {
			m.Spi = types.StringValue(v)
		} else {
			m.Spi = types.StringNull()
		}
	} else {
		m.Spi = types.StringNull()
	}
	if v, ok := obj["ssh"]; ok {
		_ = v
		if v != "" {
			m.SSH = types.StringValue(v)
		} else {
			m.SSH = types.StringNull()
		}
	} else {
		m.SSH = types.StringNull()
	}
	if v, ok := obj["ssl"]; ok {
		_ = v
		if v != "" {
			m.SSL = types.StringValue(v)
		} else {
			m.SSL = types.StringNull()
		}
	} else {
		m.SSL = types.StringNull()
	}
	if v, ok := obj["telnet"]; ok {
		_ = v
		if v != "" {
			m.Telnet = types.StringValue(v)
		} else {
			m.Telnet = types.StringNull()
		}
	} else {
		m.Telnet = types.StringNull()
	}
	if v, ok := obj["tftp"]; ok {
		_ = v
		if v != "" {
			m.Tftp = types.StringValue(v)
		} else {
			m.Tftp = types.StringNull()
		}
	} else {
		m.Tftp = types.StringNull()
	}
	if v, ok := obj["traffic-accounting"]; ok {
		_ = v
		if v != "" {
			m.TrafficAccounting = types.StringValue(v)
		} else {
			m.TrafficAccounting = types.StringNull()
		}
	} else {
		m.TrafficAccounting = types.StringNull()
	}
	if v, ok := obj["traffic-flow"]; ok {
		_ = v
		if v != "" {
			m.TrafficFlow = types.StringValue(v)
		} else {
			m.TrafficFlow = types.StringNull()
		}
	} else {
		m.TrafficFlow = types.StringNull()
	}
	if v, ok := obj["unclassified"]; ok {
		_ = v
		if v != "" {
			m.Unclassified = types.StringValue(v)
		} else {
			m.Unclassified = types.StringNull()
		}
	} else {
		m.Unclassified = types.StringNull()
	}
	if v, ok := obj["upnp"]; ok {
		_ = v
		if v != "" {
			m.Upnp = types.StringValue(v)
		} else {
			m.Upnp = types.StringNull()
		}
	} else {
		m.Upnp = types.StringNull()
	}
	if v, ok := obj["usb"]; ok {
		_ = v
		if v != "" {
			m.Usb = types.StringValue(v)
		} else {
			m.Usb = types.StringNull()
		}
	} else {
		m.Usb = types.StringNull()
	}
	if v, ok := obj["user-manager"]; ok {
		_ = v
		if v != "" {
			m.UserManager = types.StringValue(v)
		} else {
			m.UserManager = types.StringNull()
		}
	} else {
		m.UserManager = types.StringNull()
	}
	if v, ok := obj["vrrp"]; ok {
		_ = v
		if v != "" {
			m.VRRP = types.StringValue(v)
		} else {
			m.VRRP = types.StringNull()
		}
	} else {
		m.VRRP = types.StringNull()
	}
	if v, ok := obj["web-proxy"]; ok {
		_ = v
		if v != "" {
			m.WebProxy = types.StringValue(v)
		} else {
			m.WebProxy = types.StringNull()
		}
	} else {
		m.WebProxy = types.StringNull()
	}
	if v, ok := obj["winbox"]; ok {
		_ = v
		if v != "" {
			m.Winbox = types.StringValue(v)
		} else {
			m.Winbox = types.StringNull()
		}
	} else {
		m.Winbox = types.StringNull()
	}
	if v, ok := obj["wireguard"]; ok {
		_ = v
		if v != "" {
			m.Wireguard = types.StringValue(v)
		} else {
			m.Wireguard = types.StringNull()
		}
	} else {
		m.Wireguard = types.StringNull()
	}
	if v, ok := obj["wireless"]; ok {
		_ = v
		if v != "" {
			m.Wireless = types.StringValue(v)
		} else {
			m.Wireless = types.StringNull()
		}
	} else {
		m.Wireless = types.StringNull()
	}
	if v, ok := obj["www"]; ok {
		_ = v
		if v != "" {
			m.Www = types.StringValue(v)
		} else {
			m.Www = types.StringNull()
		}
	} else {
		m.Www = types.StringNull()
	}
	if v, ok := obj["zerotier"]; ok {
		_ = v
		if v != "" {
			m.Zerotier = types.StringValue(v)
		} else {
			m.Zerotier = types.StringNull()
		}
	} else {
		m.Zerotier = types.StringNull()
	}
}
