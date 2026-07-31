---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_service"
description: |-
  One row of RouterOS /ip/service -- the management services (api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl).
---

# Resource: routeros_ip_service

Manages one row of RouterOS `/ip/service` — the management services `api`,
`api-ssl`, `ftp`, `ssh`, `telnet`, `winbox`, `www` and `www-ssl`.

RouterOS ships a fixed set of service rows: they can be enabled, disabled and
reconfigured, but never added or removed. This resource therefore **adopts** the
existing service named by `name` rather than creating one. Consequences:

* `name` is immutable — changing it replaces the resource, which just means
  adopting a different row.
* `terraform destroy` only forgets the row. A service you disabled stays
  disabled; the provider will not silently re-enable telnet or FTP for you. Set
  `disabled = false` and apply if you want a service back on.
* Applying to a service name that does not exist on the device is an error, and
  the message lists the names the router actually exposes.

TFTP is **not** part of `/ip/service`. The RouterOS TFTP server only answers for
paths listed in `/ip/tftp` — see `routeros_ip_tftp` (and
`routeros_ip_tftp_settings`). Removing or disabling those entries turns TFTP off.

## Safety

The resource refuses a change that would leave *every* management service
(`winbox`, `ssh`, `api`, `api-ssl`, `www`, `www-ssl`, `telnet`) disabled,
because that locks you out of the router. Set `lockout_ack = true` on the
resource to override the guard when you know what you are doing (for example
when management moves to MAC-Winbox only).

## Example Usage

```terraform
# Turn off the legacy plaintext services.
resource "routeros_ip_service" "telnet" {
  # router = "my-router"  # which router to target; omit for the default
  name     = "telnet"
  disabled = true
}

resource "routeros_ip_service" "ftp" {
  name     = "ftp"
  disabled = true
}

resource "routeros_ip_service" "api" {
  name     = "api"
  disabled = true
}

# Plaintext API off, TLS API on and restricted to the management subnet.
resource "routeros_ip_service" "api_ssl" {
  name        = "api-ssl"
  disabled    = false
  port        = 8729
  address     = "10.0.0.0/24"
  certificate = "api-cert"
  tls_version = "only-v1.2"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`. Comma-separated list of IP/IPv6 prefixes allowed to reach the service. Empty means any source.
* `certificate` - (Optional) Type: `string`. Certificate used by the TLS-enabled services (api-ssl, www-ssl, reverse-proxy). `none` to unset. RouterOS rejects this on the plaintext services.
* `disabled` - (Optional) Type: `bool`. Whether the service is disabled. Set to true to turn the service off.
* `lockout_ack` - (Optional) Type: `bool`. Acknowledge that this change may leave no enabled management service, locking you out of the router.
* `max_sessions` - (Optional) Type: `int`. Maximum number of concurrent sessions (RouterOS 7.13+).
* `name` - (Required) Type: `string`. Service name as listed by `/ip/service print`: api, api-ssl, ftp, ssh, telnet, winbox, www, www-ssl, and reverse-proxy on newer releases. Deliberately not validated against a fixed list -- the set varies by RouterOS version; an unknown name is reported at apply time with the names the device actually exposes.
* `port` - (Optional) Type: `int`. TCP port the service listens on.
* `tls_version` - (Optional) Type: `string`. Minimum accepted TLS version: `any` or `only-v1.2`. TLS-enabled services only (api-ssl, www-ssl, reverse-proxy).
* `vrf` - (Optional) Type: `string`. VRF the service is bound to.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Services are imported by name (recommended) or by RouterOS `.id`, optionally
prefixed by the router name:

```sh
# Default router, by service name
terraform import routeros_ip_service.telnet 'telnet'

# Named router, by service name
terraform import routeros_ip_service.telnet 'home/telnet'

# By RouterOS .id
terraform import routeros_ip_service.telnet 'home/*1'
```
