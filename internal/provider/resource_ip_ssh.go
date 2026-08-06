package provider

import (
	"context"
	"fmt"
	"strings"

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
	_ resource.Resource                = &IPSSHResource{}
	_ resource.ResourceWithImportState = &IPSSHResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPSSHResource struct {
	reg *client.Registry
}

type IPSSHModel struct {
	ID                             types.String `tfsdk:"id"`
	Ciphers                        types.String `tfsdk:"ciphers"`
	ForwardingEnabled              types.String `tfsdk:"forwarding_enabled"`
	HostKeySize                    types.Int64  `tfsdk:"host_key_size"`
	HostKeyType                    types.String `tfsdk:"host_key_type"`
	PasswordAuthentication         types.String `tfsdk:"password_authentication"`
	PublickeyAuthenticationOptions types.String `tfsdk:"publickey_authentication_options"`
	StrongCrypto                   types.Bool   `tfsdk:"strong_crypto"`
	Router                         types.String `tfsdk:"router"`
}

func NewIPSSHResource() resource.Resource { return &IPSSHResource{} }

func (r *IPSSHResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ssh"
}

func (r *IPSSHResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPSSHResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ssh`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ciphers": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Allow to configure SSH ciphers.",
			},
			"forwarding_enabled": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Allows to control which SSH forwarding method to allow: `no` - SSH forwarding is disabled; `local` - Allow SSH clients to originate connections from the server(router), this setting controls also dynamic forwarding; `remote` - Allow SSH clients to listen on the server(router) and forward incoming connections; `both` - Allow both local and remote forwarding methods.",
				Validators:  []validator.String{schemautil.OneOf([]string{"both", "local", "no", "remote"}...)},
			},
			"host_key_size": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "RSA key size when host key is being regenerated.",
			},
			"host_key_type": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Select host key type",
			},
			"password_authentication": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Whether to allow password login at the same time when public key authorization is configured for a user.",
			},
			"publickey_authentication_options": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Sets public key authentication options. The touch-required option causes public key authentication using a FIDO authenticator algorithm\u00a0to always require the signature to attest that a physically present user explicitly confirmed the authentication (usually by touching the authenticator). The verify-required option requires a FIDO key signature attest that the user was verified, e.g. via a PIN.",
			},
			"strong_crypto": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Use stronger encryption, HMAC algorithms, use bigger DH primes and disallow weaker ones: use 256 and 192 bit encryption instead of 128 bits; disable null encryption; use sha256 for hashing instead of sha1; disable md5; use 2048bit prime for Diffie-Hellman exchange instead of 1024bit.",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPSSHResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSSHModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPSSHUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSSHResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPSSHModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPSSHModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPSSHUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSSHResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSSHModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/ssh")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/ssh failed", err.Error())
		return
	}
	iPSSHApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/ssh", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPSSHResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPSSHResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/ssh" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/ssh", types.StringValue(routerName))))...)
}

func iPSSHUpsert(ctx context.Context, reg *client.Registry, plan, state *IPSSHModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Ciphers.IsNull() || plan.Ciphers.IsUnknown()) && (state == nil || !plan.Ciphers.Equal(state.Ciphers)) {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !(plan.ForwardingEnabled.IsNull() || plan.ForwardingEnabled.IsUnknown()) && (state == nil || !plan.ForwardingEnabled.Equal(state.ForwardingEnabled)) {
		body["forwarding-enabled"] = plan.ForwardingEnabled.ValueString()
	}
	if !(plan.HostKeySize.IsNull() || plan.HostKeySize.IsUnknown()) && (state == nil || !plan.HostKeySize.Equal(state.HostKeySize)) {
		body["host-key-size"] = client.FormatInt64(plan.HostKeySize.ValueInt64())
	}
	if !(plan.HostKeyType.IsNull() || plan.HostKeyType.IsUnknown()) && (state == nil || !plan.HostKeyType.Equal(state.HostKeyType)) {
		body["host-key-type"] = plan.HostKeyType.ValueString()
	}
	if !(plan.PasswordAuthentication.IsNull() || plan.PasswordAuthentication.IsUnknown()) && (state == nil || !plan.PasswordAuthentication.Equal(state.PasswordAuthentication)) {
		body["password-authentication"] = plan.PasswordAuthentication.ValueString()
	}
	if !(plan.PublickeyAuthenticationOptions.IsNull() || plan.PublickeyAuthenticationOptions.IsUnknown()) && (state == nil || !plan.PublickeyAuthenticationOptions.Equal(state.PublickeyAuthenticationOptions)) {
		body["publickey-authentication-options"] = plan.PublickeyAuthenticationOptions.ValueString()
	}
	if !(plan.StrongCrypto.IsNull() || plan.StrongCrypto.IsUnknown()) && (state == nil || !plan.StrongCrypto.Equal(state.StrongCrypto)) {
		body["strong-crypto"] = client.FormatBool(plan.StrongCrypto.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/ip/ssh", body)
	if err != nil {
		diags.AddError("Upsert /ip/ssh failed", err.Error())
		return
	}
	iPSSHApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/ssh", plan.Router))
}

func iPSSHApply(ctx context.Context, obj client.Object, m *IPSSHModel) {
	_ = ctx
	if v, ok := obj["ciphers"]; ok {
		_ = v
		if v != "" {
			m.Ciphers = types.StringValue(v)
		} else {
			m.Ciphers = types.StringNull()
		}
	}
	if v, ok := obj["forwarding-enabled"]; ok {
		_ = v
		if v != "" {
			m.ForwardingEnabled = types.StringValue(v)
		} else {
			m.ForwardingEnabled = types.StringNull()
		}
	}
	if v, ok := obj["host-key-size"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.HostKeySize = types.Int64Value(n)
		} else {
			m.HostKeySize = types.Int64Null()
		}
	}
	if v, ok := obj["host-key-type"]; ok {
		_ = v
		if v != "" {
			m.HostKeyType = types.StringValue(v)
		} else {
			m.HostKeyType = types.StringNull()
		}
	}
	if v, ok := obj["password-authentication"]; ok {
		_ = v
		if v != "" {
			m.PasswordAuthentication = types.StringValue(v)
		} else {
			m.PasswordAuthentication = types.StringNull()
		}
	}
	if v, ok := obj["publickey-authentication-options"]; ok {
		_ = v
		if v != "" {
			m.PublickeyAuthenticationOptions = types.StringValue(v)
		} else {
			m.PublickeyAuthenticationOptions = types.StringNull()
		}
	}
	if v, ok := obj["strong-crypto"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.StrongCrypto = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.StrongCrypto = types.BoolValue(true)
		} else {
			m.StrongCrypto = types.BoolNull()
		}
	}
}
