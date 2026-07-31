package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &IPCloudResource{}
	_ resource.ResourceWithImportState = &IPCloudResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPCloudResource struct {
	reg *client.Registry
}

type IPCloudModel struct {
	ID                             types.String `tfsdk:"id"`
	BackToHomeVPN                  types.String `tfsdk:"back_to_home_vpn"`
	DdnsEnabled                    types.String `tfsdk:"ddns_enabled"`
	DdnsUpdateInterval             types.String `tfsdk:"ddns_update_interval"`
	DNSName                        types.String `tfsdk:"dns_name"`
	PublicAddress                  types.String `tfsdk:"public_address"`
	PublicAddressIvp6              types.String `tfsdk:"public_address_ivp6"`
	Status                         types.String `tfsdk:"status"`
	UpdateTime                     types.Bool   `tfsdk:"update_time"`
	VPNDNSName                     types.String `tfsdk:"vpn_dns_name"`
	VPNInterface                   types.String `tfsdk:"vpn_interface"`
	VPNPeerPrivateKey              types.String `tfsdk:"vpn_peer_private_key"`
	VPNPeerPublicKey               types.String `tfsdk:"vpn_peer_public_key"`
	VPNPort                        types.String `tfsdk:"vpn_port"`
	VPNPreferRelayCode             types.String `tfsdk:"vpn_prefer_relay_code"`
	VPNPrivateKey                  types.String `tfsdk:"vpn_private_key"`
	VPNPublicKey                   types.String `tfsdk:"vpn_public_key"`
	VPNRelayAddressess             types.String `tfsdk:"vpn_relay_addressess"`
	VPNRelayAddressessIPV6         types.String `tfsdk:"vpn_relay_addressess_ipv6"`
	VPNRelayCodes                  types.String `tfsdk:"vpn_relay_codes"`
	VPNRelayIpv4Status             types.String `tfsdk:"vpn_relay_ipv4_status"`
	VPNRelayIPV6Status             types.String `tfsdk:"vpn_relay_ipv6_status"`
	VPNRelayRtts                   types.String `tfsdk:"vpn_relay_rtts"`
	VPNStatus                      types.String `tfsdk:"vpn_status"`
	VPNWireguardClientConfig       types.String `tfsdk:"vpn_wireguard_client_config"`
	VPNWireguardClientConfigQrcode types.String `tfsdk:"vpn_wireguard_client_config_qrcode"`
	Warning                        types.String `tfsdk:"warning"`
	Router                         types.String `tfsdk:"router"`
}

func NewIPCloudResource() resource.Resource { return &IPCloudResource{} }

func (r *IPCloudResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_cloud"
}

func (r *IPCloudResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPCloudResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "MikroTik Cloud (DDNS) singleton — async DDNS state propagation makes acc tests flaky.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"back_to_home_vpn": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Enables or revokes and disables the Back to Home service. ddns-enabled has to be set to yes, for BTH to function.",
			},
			"ddns_enabled": schema.StringAttribute{Optional: true, Computed: true,
				Description: "If set to yes , then the device will send an encrypted message to MikroTik's Cloud server. The server will then decrypt the message and verify that the sender is an authentic MikroTik device. If all is OK, then MikroTik's Cloud server will create a DDNS record for this device and send a response to the device. Every minute the IP/Cloud service on the router will check if the WAN IP address matches the one sent to MikroTik's Cloud server and will send an encrypted update to the cloud server if the IP address changes. If set to auto, ddns will only be enabled if Back To Home is enabled. prior to the 7.17 versions, the default value was \"no\".",
			},
			"ddns_update_interval": schema.StringAttribute{Optional: true, Computed: true,
				Description: "If set DDNS will attempt to connect IP Cloud servers at the set interval. If set to none it will continue to internally check IP address update and connect to IP Cloud servers as needed. Useful if the IP address used is not on the router itself and thus, cannot be checked as a value internal to the router.",
			},
			"dns_name": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Shows the DNS name assigned to the device. Name consists of 12 characters serial number appended by . sn.mynetname.net . This field is visible only after at least one ddns-request is successfully completed.",
			},
			"public_address": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Shows the device's IPv4 address that was sent to the cloud server. This field is visible only after at least one IP Cloud request was successfully completed.",
			},
			"public_address_ivp6": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Shows the device's IPv6 address that was sent to the cloud server. This field is visible only after at least one IP Cloud request was successfully completed.",
			},
			"status": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Contains text string that describes the current dns-service state. The messages are self explanatory updating... updated Error: no Internet connection Error: request timed out Error: REJECTED. Contact MikroTik support Error: internal error - should not happen. One possible cause is if the router runs out of memory",
			},
			"update_time": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "If set to yes then router clock will be set to time, provided by the cloud server IF there is no NTP or SNTP client enabled. If set to no , then IP/Cloud service will never update the device's clock. If update-time is set to yes , Clock will be updated even when ddns-enabled is set to auto.",
			},
			"vpn_dns_name": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Shows the DNS name assigned to the device. Name consists of product serial number appended by \u00a0 .vpn.mynetname.net . This field is visible only after at least one ddns-request is successfully completed.",
			},
			"vpn_interface": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Name of the created interface for Back to Home WireGuard ® tunnel.",
			},
			"vpn_peer_private_key": schema.StringAttribute{Optional: true,
				Sensitive: true, Computed: true,
				Description: "Peer private key",
			},
			"vpn_peer_public_key": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Peer public key",
			},
			"vpn_port": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Port used by BTH VPN.",
			},
			"vpn_prefer_relay_code": schema.StringAttribute{Optional: true, Computed: true,
				Description: "You can enter relay code that will be preferred for BTH connection, if not set, relay with smallest RTT will be chosen.",
			},
			"vpn_private_key": schema.StringAttribute{Optional: true,
				Sensitive: true, Computed: true,
				Description: "Private key for BTH",
			},
			"vpn_public_key": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Public key for BTH",
			},
			"vpn_relay_addressess": schema.StringAttribute{Optional: true, Computed: true,
				Description: "IPv4 address of the relay",
			},
			"vpn_relay_addressess_ipv6": schema.StringAttribute{Optional: true, Computed: true,
				Description: "IPv6 address of the relay",
			},
			"vpn_relay_codes": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Available VPN relay codes, which can be referenced in vpn-prefer-relay-code. All available relays will be shown here.",
			},
			"vpn_relay_ipv4_status": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Status on connection to relay and detailed information about relay",
			},
			"vpn_relay_ipv6_status": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Status on connection to relay and detailed information about relay",
			},
			"vpn_relay_rtts": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Round trip time in milliseconds for each available relay, values are shown both for IPv4 and IPv6.",
			},
			"vpn_status": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Contains text string that describes the current BTH state.",
			},
			"vpn_wireguard_client_config": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Configuration that can be entered in your preferred WireGuard® client. Only one client at a time will be available to use this config.",
			},
			"vpn_wireguard_client_config_qrcode": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Scannable QR Code for your preferred WireGuard® client. Only one client at a time will be available to use this config.",
			},
			"warning": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Shows a warning message if the IP address sent by the device differs from the IP address in the UDP packet header as visible by MikroTik's Cloud server. Typically this happens if the device is behind NAT. Example: \"DDNS server received a request from IP 123.123.123.123 but your local IP was 192.168.88.23; DDNS service might not work\"",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPCloudResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPCloudModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPCloudUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPCloudResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPCloudModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPCloudUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPCloudResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPCloudModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/cloud")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/cloud failed", err.Error())
		return
	}
	iPCloudApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/cloud", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPCloudResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPCloudResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/cloud" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/cloud", types.StringValue(routerName))))...)
}

func iPCloudUpsert(ctx context.Context, reg *client.Registry, plan *IPCloudModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.BackToHomeVPN.IsNull() || plan.BackToHomeVPN.IsUnknown()) {
		body["back-to-home-vpn"] = plan.BackToHomeVPN.ValueString()
	}
	if !(plan.DdnsEnabled.IsNull() || plan.DdnsEnabled.IsUnknown()) {
		body["ddns-enabled"] = plan.DdnsEnabled.ValueString()
	}
	if !(plan.DdnsUpdateInterval.IsNull() || plan.DdnsUpdateInterval.IsUnknown()) {
		body["ddns-update-interval"] = plan.DdnsUpdateInterval.ValueString()
	}
	if !(plan.UpdateTime.IsNull() || plan.UpdateTime.IsUnknown()) {
		body["update-time"] = client.FormatBool(plan.UpdateTime.ValueBool())
	}
	if !(plan.VPNPreferRelayCode.IsNull() || plan.VPNPreferRelayCode.IsUnknown()) {
		body["vpn-prefer-relay-code"] = plan.VPNPreferRelayCode.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/cloud", body)
	if err != nil {
		diags.AddError("Upsert /ip/cloud failed", err.Error())
		return
	}
	iPCloudApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/cloud", plan.Router))
}

func iPCloudApply(ctx context.Context, obj client.Object, m *IPCloudModel) {
	_ = ctx
	if v, ok := obj["back-to-home-vpn"]; ok {
		_ = v
		if v != "" {
			m.BackToHomeVPN = types.StringValue(v)
		} else {
			m.BackToHomeVPN = types.StringNull()
		}
	}
	if v, ok := obj["ddns-enabled"]; ok {
		_ = v
		if v != "" {
			m.DdnsEnabled = types.StringValue(v)
		} else {
			m.DdnsEnabled = types.StringNull()
		}
	}
	if v, ok := obj["ddns-update-interval"]; ok {
		_ = v
		if v != "" {
			m.DdnsUpdateInterval = types.StringValue(v)
		} else {
			m.DdnsUpdateInterval = types.StringNull()
		}
	}
	if v, ok := obj["dns-name"]; ok {
		_ = v
		if v != "" {
			m.DNSName = types.StringValue(v)
		} else {
			m.DNSName = types.StringNull()
		}
	}
	if v, ok := obj["public-address"]; ok {
		_ = v
		if v != "" {
			m.PublicAddress = types.StringValue(v)
		} else {
			m.PublicAddress = types.StringNull()
		}
	}
	if v, ok := obj["public-address-ivp6"]; ok {
		_ = v
		if v != "" {
			m.PublicAddressIvp6 = types.StringValue(v)
		} else {
			m.PublicAddressIvp6 = types.StringNull()
		}
	}
	if v, ok := obj["status"]; ok {
		_ = v
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	}
	if v, ok := obj["update-time"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UpdateTime = types.BoolValue(b)
		} else {
			m.UpdateTime = types.BoolNull()
		}
	}
	if v, ok := obj["vpn-dns-name"]; ok {
		_ = v
		if v != "" {
			m.VPNDNSName = types.StringValue(v)
		} else {
			m.VPNDNSName = types.StringNull()
		}
	}
	if v, ok := obj["vpn-interface"]; ok {
		_ = v
		if v != "" {
			m.VPNInterface = types.StringValue(v)
		} else {
			m.VPNInterface = types.StringNull()
		}
	}
	if v, ok := obj["vpn-peer-private-key"]; ok {
		_ = v
		if v != "" {
			m.VPNPeerPrivateKey = types.StringValue(v)
		} else {
			m.VPNPeerPrivateKey = types.StringNull()
		}
	}
	if v, ok := obj["vpn-peer-public-key"]; ok {
		_ = v
		if v != "" {
			m.VPNPeerPublicKey = types.StringValue(v)
		} else {
			m.VPNPeerPublicKey = types.StringNull()
		}
	}
	if v, ok := obj["vpn-port"]; ok {
		_ = v
		if v != "" {
			m.VPNPort = types.StringValue(v)
		} else {
			m.VPNPort = types.StringNull()
		}
	}
	if v, ok := obj["vpn-prefer-relay-code"]; ok {
		_ = v
		if v != "" {
			m.VPNPreferRelayCode = types.StringValue(v)
		} else {
			m.VPNPreferRelayCode = types.StringNull()
		}
	}
	if v, ok := obj["vpn-private-key"]; ok {
		_ = v
		if v != "" {
			m.VPNPrivateKey = types.StringValue(v)
		} else {
			m.VPNPrivateKey = types.StringNull()
		}
	}
	if v, ok := obj["vpn-public-key"]; ok {
		_ = v
		if v != "" {
			m.VPNPublicKey = types.StringValue(v)
		} else {
			m.VPNPublicKey = types.StringNull()
		}
	}
	if v, ok := obj["vpn-relay-addressess"]; ok {
		_ = v
		if v != "" {
			m.VPNRelayAddressess = types.StringValue(v)
		} else {
			m.VPNRelayAddressess = types.StringNull()
		}
	}
	if v, ok := obj["vpn-relay-addressess-ipv6"]; ok {
		_ = v
		if v != "" {
			m.VPNRelayAddressessIPV6 = types.StringValue(v)
		} else {
			m.VPNRelayAddressessIPV6 = types.StringNull()
		}
	}
	if v, ok := obj["vpn-relay-codes"]; ok {
		_ = v
		if v != "" {
			m.VPNRelayCodes = types.StringValue(v)
		} else {
			m.VPNRelayCodes = types.StringNull()
		}
	}
	if v, ok := obj["vpn-relay-ipv4-status"]; ok {
		_ = v
		if v != "" {
			m.VPNRelayIpv4Status = types.StringValue(v)
		} else {
			m.VPNRelayIpv4Status = types.StringNull()
		}
	}
	if v, ok := obj["vpn-relay-ipv6-status"]; ok {
		_ = v
		if v != "" {
			m.VPNRelayIPV6Status = types.StringValue(v)
		} else {
			m.VPNRelayIPV6Status = types.StringNull()
		}
	}
	if v, ok := obj["vpn-relay-rtts"]; ok {
		_ = v
		if v != "" {
			m.VPNRelayRtts = types.StringValue(v)
		} else {
			m.VPNRelayRtts = types.StringNull()
		}
	}
	if v, ok := obj["vpn-status"]; ok {
		_ = v
		if v != "" {
			m.VPNStatus = types.StringValue(v)
		} else {
			m.VPNStatus = types.StringNull()
		}
	}
	if v, ok := obj["vpn-wireguard-client-config"]; ok {
		_ = v
		if v != "" {
			m.VPNWireguardClientConfig = types.StringValue(v)
		} else {
			m.VPNWireguardClientConfig = types.StringNull()
		}
	}
	if v, ok := obj["vpn-wireguard-client-config-qrcode"]; ok {
		_ = v
		if v != "" {
			m.VPNWireguardClientConfigQrcode = types.StringValue(v)
		} else {
			m.VPNWireguardClientConfigQrcode = types.StringNull()
		}
	}
	if v, ok := obj["warning"]; ok {
		_ = v
		if v != "" {
			m.Warning = types.StringValue(v)
		} else {
			m.Warning = types.StringNull()
		}
	}
}
