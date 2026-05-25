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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `exchange_mode` - (Optional) Type: `enum(base|main|aggressive|ike2)`. Default: `2`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `passive` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `int`.
* `profile` - (Optional) Type: `string`.
* `responder` - (Optional) Type: `bool`.
* `send_initial_contact` - (Optional) Type: `bool`. Default: `1`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.

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
