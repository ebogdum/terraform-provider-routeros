---
subcategory: "Interface"
page_title: "RouterOS: routeros_interface_sstp_server_server"
description: |-
  Mirrors RouterOS /interface/sstp-server/server.
---

# Resource: routeros_interface_sstp_server_server

Mirrors RouterOS `/interface/sstp-server/server`.

## Example Usage

```terraform
resource "routeros_interface_sstp_server_server" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # authentication = "replace-me"
  # certificate = "replace-me"
  # ciphers = "replace-me"
  # default_profile = "replace-me"
  # enabled = true
  # keepalive_timeout = 0
  # max_mru = 0
  # max_mtu = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `authentication` - (Optional) Type: `string`. RouterOS `authentication`.
* `certificate` - (Optional) Type: `string`. RouterOS `certificate`.
* `ciphers` - (Optional) Type: `string`. RouterOS `ciphers`.
* `default_profile` - (Optional) Type: `string`. RouterOS `default-profile`.
* `enabled` - (Optional) Type: `bool`. RouterOS `enabled`.
* `keepalive_timeout` - (Optional) Type: `int`. RouterOS `keepalive-timeout`.
* `max_mru` - (Optional) Type: `int`. RouterOS `max-mru`.
* `max_mtu` - (Optional) Type: `int`. RouterOS `max-mtu`.
* `mrru` - (Optional) Type: `string`. RouterOS `mrru`.
* `pfs` - (Optional) Type: `string`. Perfect forward secrecy: `no`, `yes` (offer PFS) or `required` (refuse clients without it).
* `port` - (Optional) Type: `int`. RouterOS `port`.
* `tls_version` - (Optional) Type: `string`. RouterOS `tls-version`.
* `verify_client_certificate` - (Optional) Type: `bool`. RouterOS `verify-client-certificate`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_sstp_server_server.this 'home'
```
