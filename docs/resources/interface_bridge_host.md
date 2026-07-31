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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `aged` - (Read-only) Type: `bool`.
* `aged_on_peer` - (Read-only) Type: `bool`.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `external_fdb` - (Read-only) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `local` - (Read-only) Type: `bool`.
* `mac_address` - (Optional) Type: `string`.
* `on_interface` - (Read-only) Type: `string`.
* `remote_ip` - (Read-only) Type: `string`.
* `vid` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
