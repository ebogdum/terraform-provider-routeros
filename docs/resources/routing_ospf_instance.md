---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_instance"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_ospf_instance

Manages the RouterOS `/routing/ospf/instance` menu.

## Example Usage

```terraform
resource "routeros_routing_ospf_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # domain_id = "replace-me"
  # domain_tag = "replace-me"
  # in_filter = "replace-me"
  # mpls_te_address = "10.99.0.0/24"
  # mpls_te_area = "replace-me"
  # name = "tf-example"
  # originate_default = "replace-me"
  # out_filter = "replace-me"
  # out_filter_select = "replace-me"
  # redistribute = "replace-me"
  # router_id = "replace-me"
  # routing_table = "main"
  # version = "2"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `domain_id` - (Optional) Type: `string`.
* `domain_tag` - (Optional) Type: `string`.
* `in_filter` - (Read-only) Type: `string`.
* `in_filter_chain` - (Optional) Type: `string`. RouterOS `in-filter-chain`.
* `invalid` - (Read-only) Type: `bool`.
* `mpls_te_address` - (Optional) Type: `string`.
* `mpls_te_area` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `originate_default` - (Optional) Type: `string`.
* `out_filter` - (Read-only) Type: `string`.
* `out_filter_chain` - (Optional) Type: `string`. RouterOS `out-filter-chain`.
* `out_filter_select` - (Optional) Type: `string`.
* `redistribute` - (Optional) Type: `string`.
* `router_id` - (Optional) Type: `string`.
* `routing_table` - (Optional) Type: `string`.
* `use_dn` - (Optional) Type: `string`. RouterOS `use-dn`.
* `version` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_ospf_instance.example '*3'

# Named router
terraform import routeros_routing_ospf_instance.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_ospf_instance.example 'home/my-resource-name'
```
