# Terraform Provider for MikroTik RouterOS 7

[![Terraform Registry](https://img.shields.io/badge/terraform-registry-7B42BC?logo=terraform&logoColor=white)](https://registry.terraform.io/providers/ebogdum/routeros/latest)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/ebogdum/terraform-provider-routeros.svg)](https://pkg.go.dev/github.com/ebogdum/terraform-provider-routeros)

Manage MikroTik **RouterOS 7.x** devices as code with Terraform. Complete
coverage of every device menu -- **420 menus, 182 resources, 276 data sources,
75 actions, 3289 properties** -- generated from a schema validated
property-by-property against a live router.

Manage one MikroTik router or an entire fleet from a single Terraform
configuration. Apply firewall rules in deterministic order, import existing
device state, detect and correct out-of-band drift, and ship secrets safely
through the plugin framework's sensitive-value handling.

> **Keywords:** terraform-provider-mikrotik, terraform-provider-routeros,
> mikrotik terraform, routeros terraform, mikrotik infrastructure as code,
> routeros REST API, mikrotik automation, network as code, mikrotik CHR,
> mikrotik fleet management, terraform firewall mikrotik, mikrotik IPsec
> terraform.

## Table of contents

- [Why this provider](#why-this-provider)
- [Quick start](#quick-start)
- [Multi-router fleet management](#multi-router-fleet-management)
- [Coverage](#coverage)
- [Safety guards](#safety-guards)
- [Drift detection and reconciliation](#drift-detection-and-reconciliation)
- [Resources, data sources, and actions](#resources-data-sources-and-actions)
- [Authentication and TLS](#authentication-and-tls)
- [Development](#development)
- [Release process](#release-process)
- [Upgrading](#upgrading)
- [Release history](#release-history)
- [License](#license)

## Why this provider

| Feature | This provider |
|---|---|
| RouterOS API | REST (HTTPS) |
| Menu coverage | 420 menus (every menu surfaced over REST) |
| Resources / data sources / actions | 182 / 276 / 75 |
| Multi-router from one provider block | Yes (named map, no provider aliases) |
| Deterministic firewall ordering | Yes (`position` integer; stable across destroy/recreate) |
| Lockout safety guards | Firewall, user, user-group, mac-server |
| Sensitive field redaction | 28 properties marked sensitive |
| Out-of-band drift detection | Yes (verified end-to-end) |
| Terraform import | Yes (`<router>/<.id>` format) |
| Schema source | Live device + WebFig skin files + Confluence docs + per-property device validation |
| Plugin framework | terraform-plugin-framework (v1.19.0) |
| Minimum Terraform | 1.4 |
| Minimum RouterOS | 7.1 (REST API requirement) |

## Quick start

```hcl
terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 2.0"
    }
  }
}

provider "routeros" {
  host     = "https://192.0.2.1"
  username = "admin"
  password = var.routeros_password
  insecure = true # set to false in production with a real cert
}

resource "routeros_ip_address" "lan" {
  address   = "192.168.88.1/24"
  interface = "bridge1"
  comment   = "Managed by Terraform"
}

resource "routeros_ip_firewall_filter" "allow_established" {
  chain            = "input"
  action           = "accept"
  connection_state = "established,related"
  position         = 100
  comment          = "Managed by Terraform"
  lockout_ack      = true
}
```

`terraform init && terraform apply` and you are done.

### Install for local development

```sh
# 1. Build the provider binary
make build

# 2. Tell Terraform to use the local binary instead of the registry
cat > ~/.terraformrc <<EOF
provider_installation {
  dev_overrides {
    "ebogdum/routeros" = "$(pwd)/bin"
  }
  direct {}
}
EOF
```

## Multi-router fleet management

Manage every router in your network from a single Terraform configuration.
No provider aliases, no duplicate blocks -- just a named map.

```hcl
provider "routeros" {
  routers = {
    core = {
      host     = "https://10.0.0.1"
      username = "admin"
      password = var.core_password
      insecure = true
    }
    edge_se = {
      host     = "https://10.0.1.1"
      username = "admin"
      password = var.edge_se_password
      insecure = true
    }
    edge_nw = {
      host     = "https://10.0.2.1"
      username = "admin"
      password = var.edge_nw_password
      insecure = true
    }
  }
}

# Push the same identity policy to every router with one resource.
resource "routeros_system_identity" "label" {
  for_each = toset(["core", "edge_se", "edge_nw"])
  router   = each.key
  name     = each.key
}

# Cross-router data flow: feed core's address-list into edge_se's firewall.
data "routeros_ip_firewall_address_list" "core_lan" {
  router = "core"
}

resource "routeros_ip_firewall_filter" "edge_to_core" {
  router           = "edge_se"
  chain            = "forward"
  action           = "accept"
  src_address_list = data.routeros_ip_firewall_address_list.core_lan.records[0].list
  position         = 200
  comment          = "Allow core LAN inbound"
}
```

Omit `router =` to target the default -- the entry named `default`, or the
first router in sorted order if there is no `default`.

## Coverage

Every menu the REST API surfaces is mapped. Categories include:

- **IP**: address, route, pool, ARP, DHCP server/client/relay, DNS, firewall
  (filter, NAT, mangle, raw, address-list, layer7-protocol, service-port),
  IPsec (incl. policy-group, QKD keys), hotspot (incl. service-port), proxy,
  service, neighbor (incl. discovery-settings), kid-control, traffic-flow
  (incl. IPFIX), media, cloud (incl. advanced), DNS (static, forwarders,
  adlist), settings.
- **IPv6**: address, route, pool, DHCP client/server/relay (incl. relay
  options), firewall family, neighbor discovery (incl. default prefix),
  settings.
- **Interface**: bridge (port, vlan, msti, mst-config, settings, filter, nat,
  calea, host), bonding, vlan, vrrp, wireguard (peers), wifi (wave2 +
  legacy), 6to4, eoip, eoipv6, gre, gre6, ipip, ipipv6, l2tp client/server,
  list (members), macsec, macvlan, ovpn client/server, ppp, pppoe, pptp,
  sstp, veth, vrrp, vti.
- **Routing**: BGP (connection, template), OSPF (instance, area,
  interface-template, static-neighbor), RIP (instance, interface, neighbor),
  ISIS (instance, interface-template), filter (rule, select-rule, chain),
  BFD, RPKI (session), IGMP-Proxy, PIM-SM, table, rule.
- **System**: identity, clock, NTP client/server, scheduler, script,
  package (incl. local-update mirror), resource (incl. IRQ affinity, USB
  settings), health settings, routerboard (settings and the mode/WPS/reset
  button bindings), backup, logging (action), users (user, group, ssh-keys,
  aaa), notes, history, watchdog, leds.
- **Tools**: bandwidth-test, netwatch, fetch, sniffer, traffic-generator
  (with packet templates), email, sms, romon, mac-server, snmp-get,
  snmp-walk.
- **Certificate**: full lifecycle (sign, export, import, SCEP, ACME, CRL).
- **Queue**: simple, tree, type, interface.
- **SNMP**: community, send-trap.
- **Container**, **LCD**, **IoT/LoRa**, **disk**, **environment**, **RADIUS**.

Every resource has a matching data source for querying existing state.

## Safety guards

The provider refuses to apply changes that would obviously sever management
access. Each guard is conservative and per-resource -- set
`lockout_ack = true` on the specific resource to override.

| Resource | What it refuses |
|---|---|
| `routeros_ip_firewall_filter`, `routeros_ipv6_firewall_filter` | `chain=input|forward` + `action=drop|reject|tarpit` with no narrowing match (would drop your own session) |
| `routeros_user` | Deleting or disabling the last admin account |
| `routeros_user_group` | Removing `api`, `web`, `winbox`, `ssh`, `password`, `policy`, or `write` from the `full` group |
| `routeros_tool_mac_server` | Emptying `allowed-interface-list` (kills MAC-Winbox recovery) |

Sensitive properties (passwords, secrets, keys, OTP, SIM PIN -- 28 total
across the schema) are marked `Sensitive: true`. Terraform redacts them in
plan output, state files, and CLI display.

## Drift detection and reconciliation

The provider's Read implementation queries the device directly on every plan.
If a row is changed outside Terraform (via Winbox, CLI, or another tool),
the next plan shows the drift and the following apply reconciles it.

This is verified by an explicit acceptance test (`TestAccDriftCorrectionIPAddress`)
that:
1. Creates a row through Terraform.
2. Mutates that row's `comment` field via raw REST, outside Terraform.
3. Re-applies the original config.
4. Asserts the field is restored.

## Resources, data sources, and actions

### Resources

Every collection-style menu (`/ip/address`, `/ip/firewall/filter`, ...) is a
resource. Every singleton menu (`/system/identity`, `/ip/dns`, ...) is a
resource that maps to a single device-wide value.

### Data sources

Every collection-style menu is also a data source -- query existing rows by
property filter, project specific fields via `proplist`.

```hcl
data "routeros_ip_address" "wan_ips" {
  filter = {
    interface = "ether1"
  }
}

output "wan_ip" {
  value = data.routeros_ip_address.wan_ips.records[0].address
}
```

### Actions

Action resources represent RouterOS commands (`/system/reboot`, `/log/info`,
`/tool/fetch`, `/certificate/sign`, `/file/read`, etc.). Re-running the
action is triggered by changing the `trigger` attribute (string).

```hcl
resource "routeros_log_info" "marker" {
  trigger = "v1"
  message = "Deployed by Terraform run ${terraform.workspace}"
}
```

## Authentication and TLS

The provider speaks RouterOS REST over HTTPS by default. Three authentication
modes are supported:

| Mode | How |
|---|---|
| **Insecure HTTP** | `host = "http://..."` (lab gear only) |
| **HTTPS with self-signed cert** | `host = "https://..."` + `insecure = true` |
| **HTTPS with verified cert** | `host = "https://..."` + `ca_cert = file("ca.pem")` |

Environment variables provide a shorthand for single-router setups:

| Env var | Equivalent attribute |
|---|---|
| `ROUTEROS_HOST` | `host` |
| `ROUTEROS_USER` | `username` |
| `ROUTEROS_PASSWORD` | `password` |
| `ROUTEROS_CA_CERT` | `ca_cert` |
| `ROUTEROS_INSECURE` | `insecure` |
| `ROUTEROS_VERSION` | `ros_version` |

## Development

Requires Go 1.26+ (see [go.mod](go.mod)) and Terraform 1.4+.

```sh
make build              # local provider binary
make test               # unit tests
make testacc            # acceptance -- RUNS DESTRUCTIVE WRITES; requires ROUTEROS_HOST/USER/PASSWORD
make release-snapshot   # goreleaser dry-run (no publish, no sign)
make release            # signed release (requires GPG_FINGERPRINT)
make vet                # go vet (with and without the acceptance build tag)
```

Layout:

```
internal/client/         REST client, multi-router registry, type coercion, retry
                         classification, ordered-move support.
internal/schemautil/     Validators, plan modifiers, lockout guards.
internal/provider/       Provider entrypoint + generated resources +
                         handwritten lifecycle/drift/action/import tests.
docs/                    Terraform Registry markdown.
examples/                Per-resource .tf snippets + examples/curated/home-router/.
```

The Go code in `internal/provider/` is the committed output of a code
generator that lives in a separate, non-published toolchain repo. This
repo carries only the compiled provider, its handwritten runtime, and
its release plumbing.

## Release process

Releases are cut by pushing a `v*` semver tag:

```sh
git tag -s v1.2.3 -m 'release v1.2.3'
git push origin v1.2.3
```

GitHub Actions then:

1. Imports the GPG signing key (`GPG_PRIVATE_KEY` + `PASSPHRASE` repo secrets).
2. Runs `goreleaser release --clean`, which builds 13 archives across
   linux / darwin / windows / freebsd × amd64 / arm64 / 386 / arm.
3. Signs the SHA256SUMS file with the imported GPG key.
4. Creates a draft GitHub release with the auto-generated changelog body
   for manual review and publish.
5. The Terraform Registry picks up the published release and indexes it.

CI on every PR runs: `vet`, unit tests, build, golangci-lint, gen-diff
(ensures generated files match what the schema produces), and
`tfplugindocs validate`. Acceptance tests run on `main` against a self-hosted
runner with a real CHR matrix (ROS 7.20, 7.22, latest).

## Upgrading

v2.0.0 renames and retypes attributes that could not have worked on v1. See
the [v2 upgrade guide](docs/guides/upgrading-to-v2.md) for the edits needed.

## Release history

See [CHANGELOG.md](CHANGELOG.md).

## License

MIT -- see [LICENSE](LICENSE).
