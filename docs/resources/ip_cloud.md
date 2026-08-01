---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_cloud"
description: |-
  MikroTik Cloud (DDNS) singleton — async DDNS state propagation makes acc tests flaky.
---

# Resource: routeros_ip_cloud

MikroTik Cloud (DDNS) singleton — async DDNS state propagation makes acc tests flaky.

## Example Usage

```terraform
resource "routeros_ip_cloud" "cloud_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # back_to_home_vpn = "replace-me"
  # ddns_enabled = "replace-me"
  # ddns_update_interval = "replace-me"
  # dns_name = "replace-me"
  # public_address = "10.99.0.0/24"
  # public_address_ivp6 = "replace-me"
  # status = "replace-me"
  # update_time = false
  # vpn_dns_name = "replace-me"
  # vpn_interface = "replace-me"
  # vpn_peer_private_key = "REDACTED"
  # vpn_peer_public_key = "REDACTED"
  # vpn_port = "443"
  # vpn_prefer_relay_code = "replace-me"
  # vpn_private_key = "REDACTED"
  # vpn_public_key = "REDACTED"
  # vpn_relay_addressess = "replace-me"
  # vpn_relay_addressess_ipv6 = "replace-me"
  # vpn_relay_codes = "replace-me"
  # vpn_relay_ipv4_status = "replace-me"
  # vpn_relay_ipv6_status = "replace-me"
  # vpn_relay_rtts = "replace-me"
  # vpn_status = "replace-me"
  # vpn_wireguard_client_config = "replace-me"
  # vpn_wireguard_client_config_qrcode = "replace-me"
  # warning = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `back_to_home_vpn` - (Optional) Type: `string`. Enables or revokes and disables the Back to Home service. ddns-enabled has to be set to yes, for BTH to function.
* `ddns_enabled` - (Optional) Type: `string`. If set to yes , then the device will send an encrypted message to MikroTik's Cloud server. The server will then decrypt the message and verify that the sender is an authentic MikroTik device. If all is OK, then MikroTik's Cloud server will create a DDNS record for this device and send a response to the device. Every minute the IP/Cloud service on the router will check if the WAN IP address matches the one sent to MikroTik's Cloud server and will send an encrypted update to the cloud server if the IP address changes. If set to auto, ddns will only be enabled if Back To Home is enabled. prior to the 7.17 versions, the default value was "no".
* `ddns_update_interval` - (Optional) Type: `string`. If set DDNS will attempt to connect IP Cloud servers at the set interval. If set to none it will continue to internally check IP address update and connect to IP Cloud servers as needed. Useful if the IP address used is not on the router itself and thus, cannot be checked as a value internal to the router.
* `dns_name` - (Optional) Type: `string`. Shows the DNS name assigned to the device. Name consists of 12 characters serial number appended by . sn.mynetname.net . This field is visible only after at least one ddns-request is successfully completed.
* `public_address` - (Read-only) Type: `string`. Shows the device's IPv4 address that was sent to the cloud server. This field is visible only after at least one IP Cloud request was successfully completed.
* `public_address_ipv6` - (Read-only) Type: `string`. Shows the device's IPv6 address that was sent to the cloud server. This field is visible only after at least one IP Cloud request was successfully completed.
* `status` - (Optional) Type: `string`. Contains text string that describes the current dns-service state. The messages are self explanatory updating... updated Error: no Internet connection Error: request timed out Error: REJECTED. Contact MikroTik support Error: internal error - should not happen. One possible cause is if the router runs out of memory
* `update_time` - (Optional) Type: `bool`. If set to yes then router clock will be set to time, provided by the cloud server IF there is no NTP or SNTP client enabled. If set to no , then IP/Cloud service will never update the device's clock. If update-time is set to yes , Clock will be updated even when ddns-enabled is set to auto.
* `vpn_dns_name` - (Optional) Type: `string`. Shows the DNS name assigned to the device. Name consists of product serial number appended by .vpn.mynetname.net . This field is visible only after at least one ddns-request is successfully completed.
* `vpn_interface` - (Optional) Type: `string`. Name of the created interface for Back to Home WireGuard ® tunnel.
* `vpn_peer_private_key` - (Optional) Type: `string`. Peer private key **Sensitive.**
* `vpn_peer_public_key` - (Optional) Type: `string`. Peer public key
* `vpn_port` - (Optional) Type: `string`. Port used by BTH VPN.
* `vpn_prefer_relay_code` - (Optional) Type: `string`. You can enter relay code that will be preferred for BTH connection, if not set, relay with smallest RTT will be chosen.
* `vpn_private_key` - (Optional) Type: `string`. Private key for BTH **Sensitive.**
* `vpn_public_key` - (Optional) Type: `string`. Public key for BTH
* `vpn_relay_addressess` - (Optional) Type: `string`. IPv4 address of the relay
* `vpn_relay_addressess_ipv6` - (Optional) Type: `string`. IPv6 address of the relay
* `vpn_relay_codes` - (Optional) Type: `string`. Available VPN relay codes, which can be referenced in vpn-prefer-relay-code. All available relays will be shown here.
* `vpn_relay_ipv4_status` - (Optional) Type: `string`. Status on connection to relay and detailed information about relay
* `vpn_relay_ipv6_status` - (Optional) Type: `string`. Status on connection to relay and detailed information about relay
* `vpn_relay_rtts` - (Optional) Type: `string`. Round trip time in milliseconds for each available relay, values are shown both for IPv4 and IPv6.
* `vpn_status` - (Optional) Type: `string`. Contains text string that describes the current BTH state.
* `vpn_wireguard_client_config` - (Optional) Type: `string`. Configuration that can be entered in your preferred WireGuard® client. Only one client at a time will be available to use this config.
* `vpn_wireguard_client_config_qrcode` - (Optional) Type: `string`. Scannable QR Code for your preferred WireGuard® client. Only one client at a time will be available to use this config.
* `warning` - (Optional) Type: `string`. Shows a warning message if the IP address sent by the device differs from the IP address in the UDP packet header as visible by MikroTik's Cloud server. Typically this happens if the device is behind NAT. Example: "DDNS server received a request from IP 123.123.123.123 but your local IP was 192.168.88.23; DDNS service might not work"

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_cloud.this 'home'
```
