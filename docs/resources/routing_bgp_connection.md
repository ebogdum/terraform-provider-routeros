---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_bgp_connection"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_routing_bgp_connection

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_routing_bgp_connection" "connection_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # as = "replace-me"
  # connect = "replace-me"
  # hold_time = "replace-me"
  # instance = "replace-me"
  # keepalive_time = "replace-me"
  # listen = "replace-me"
  # multihop = "replace-me"
  # name = "tf-example"
  # nexthop_choice = "replace-me"
  # routing_table = "main"
  # tcp_md5_key = "REDACTED"
  # use_bfd = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `afi` - (Optional) Type: `string`.
* `as` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hold_time` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `keepalive_time` - (Optional) Type: `string`.
* `listen` - (Optional) Type: `string`.
* `multihop` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nexthop_choice` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `tcp_md5_key` - (Optional) Type: `string`. **Sensitive.**
* `use_bfd` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_bgp_connection.example '*3'

# Named router
terraform import routeros_routing_bgp_connection.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_bgp_connection.example 'home/my-resource-name'
```
