---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_host"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_interface_bridge_host

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_interface_bridge_host" "host_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # aged = false
  # aged_on_peer = false
  # bridge = "bridge1"
  # external_fdb = false
  # interface = "ether1"
  # local = false
  # mac_address = "10.99.0.0/24"
  # vid = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `aged` - (Optional) Type: `bool`.
* `aged_on_peer` - (Optional) Type: `bool`.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `external_fdb` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `local` - (Optional) Type: `bool`.
* `mac_address` - (Optional) Type: `string`.
* `vid` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.
* `on_interface` - Type: `string`.
* `remote_ip` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge_host.example '*3'

# Named router
terraform import routeros_interface_bridge_host.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge_host.example 'home/my-resource-name'
```
