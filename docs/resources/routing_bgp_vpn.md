---
subcategory: "BGP"
page_title: "RouterOS: routeros_routing_bgp_vpn"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_routing_bgp_vpn

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_routing_bgp_vpn" "vpn_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # export_filter = "replace-me"
  # export_route_targets = "replace-me"
  # export_select = "replace-me"
  # import_filter = "replace-me"
  # import_route_targets = "replace-me"
  # instance = "replace-me"
  # label_allocation_policy = ""
  # redistribute = "replace-me"
  # route_distinguisher = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `export_filter` - (Read-only) Type: `string`.
* `export_filter_chain` - (Optional) Type: `string`. RouterOS `export.filter-chain`.
* `export_filter_select` - (Optional) Type: `string`. RouterOS `export.filter-select`.
* `export_redistribute` - (Optional) Type: `string`. RouterOS `export.redistribute`.
* `export_route_targets` - (Optional) Type: `string`.
* `export_select` - (Read-only) Type: `string`.
* `import_filter` - (Read-only) Type: `string`.
* `import_filter_chain` - (Optional) Type: `string`. RouterOS `import.filter-chain`.
* `import_route_targets` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `label_allocation_policy` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `redistribute` - (Read-only) Type: `string`.
* `route_distinguisher` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bgp_vpn.example '*3'

# Named router
terraform import routeros_routing_bgp_vpn.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bgp_vpn.example 'home/my-resource-name'
```
