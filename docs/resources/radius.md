---
page_title: "RouterOS: routeros_radius"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_radius

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_radius" "radius_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # accounting_backup = false
  # accounting_port = "443"
  # address = "replace-me"
  # authentication_port = "443"
  # called_id = "replace-me"
  # certificate = "replace-me"
  # domain = "example.local"
  # protocol = "udp"
  # radsec_timeout = "replace-me"
  # realm = "replace-me"
  # require_message_auth = ""
  # secret = "REDACTED"
  # service = "ppp"
  # src_address = "10.99.0.0/24"
  # timeout = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `accounting_backup` - (Optional) Type: `bool`. Whether the configuration is for the backup RADIUS server.
* `accounting_port` - (Optional) Type: `int`. RADIUS server port used for accounting. Default: `1813`.
* `address` - (Optional) Type: `string`. IPv4 or IPv6 address of RADIUS server. The following formats are accepted: -   ipv4 -   ipv4 @ vrf -   ipv6 -   ipv6 @ vrf.
* `authentication_port` - (Optional) Type: `int`. RADIUS server port used for authentication. Default: `1812`.
* `called_id` - (Optional) Type: `string`. Value depends on Point-to-Point protocol: PPPoE - service name, PPTP - server's IP address, L2TP - server's IP address.
* `certificate` - (Optional) Type: `string`. Certificate file to use for communicating with RADIUS Server with RadSec enabled.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `domain` - (Optional) Type: `string`. Microsoft Windows domain of client passed to RADIUS servers that require domain validation.
* `protocol` - (Optional) Type: `enum(udp|radsec)`. Specifies the protocol to use when communicating with the RADIUS Server.
* `radsec_timeout` - (Optional) Type: `string`. Timeout after which the request should be resent over RadSec protocol.
* `realm` - (Optional) Type: `string`. Explicitly stated realm (user domain), so the users do not have to provide proper ISP domain name in the user name.
* `require_message_auth` - (Optional) Type: `enum(|no|yes for request resp)`. Specifies if Message-Authenticator attributes are required.
* `secret` - (Optional) Type: `string`. The shared secret used to access the RADIUS server. **Sensitive.**
* `service` - (Optional) Type: `enum(ppp|login|hotspot|wireless|dhcp|ipsec, ...)`. Router services that will use this RADIUS server: hotspot - HotSpot authentication service login - router's local user authentication ppp - Point-to-Point clients authentication wireless - wireless client authentication dhcp - DHCP protocol client authentication (client's MAC address is sent as User-Name) ipsec - ipsec client authentification dot1x - dot1x authentification.
* `src_address` - (Optional) Type: `string`. Source IP/IPv6 address of the packets sent to the RADIUS server.
* `timeout` - (Optional) Type: `int`. Timeout after which the request should be resent.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_radius.example '*3'

# Named router
terraform import routeros_radius.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_radius.example 'home/my-resource-name'
```
