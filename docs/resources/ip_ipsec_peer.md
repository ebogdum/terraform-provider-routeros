---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_peer"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_ipsec_peer

Manages the RouterOS `/ip/ipsec/peer` menu.

## Example Usage

```terraform
resource "routeros_ip_ipsec_peer" "peer_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # exchange_mode = "aggressive"
  # local_address = "10.99.0.1"
  # name = "tf-example"
  # passive = false
  # port = "443"
  # profile = "replace-me"
  # responder = false
  # send_initial_contact = true
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `exchange_mode` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `passive` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `int`.
* `ppk_secret` - (Optional) Type: `string`. RouterOS `ppk-secret`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `responder` - (Read-only) Type: `bool`.
* `send_initial_contact` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_peer.example '*3'

# Named router
terraform import routeros_ip_ipsec_peer.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_peer.example 'home/my-resource-name'
```
