---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_mode_config"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_ipsec_mode_config

Manages the RouterOS `/ip/ipsec/mode-config` menu.

## Example Usage

```terraform
resource "routeros_ip_ipsec_mode_config" "mode_config_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # address_pool = "4.294967295e+09"
  # address_prefix_length = 24
  # connection_mark = "replace-me"
  # nonresp = "replace-me"
  # pool = "replace-me"
  # resp = "replace-me"
  # responder = false
  # sdns = "replace-me"
  # split_dns = "replace-me"
  # split_include = "replace-me"
  # src_address_list = "my-list"
  # static_dns = "replace-me"
  # system_dns = false
  # use_responder_dns = "exclusively"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `address_pool` - (Optional) Type: `string`.
* `address_prefix_length` - (Optional) Type: `int`.
* `connection_mark` - (Optional) Type: `string`.
* `default` - (Read-only) Type: `bool`.
* `name` - (Required) Type: `string`.
* `nonresp` - (Optional) Type: `string`.
* `pool` - (Read-only) Type: `string`.
* `resp` - (Optional) Type: `string`.
* `responder` - (Optional) Type: `bool`.
* `sdns` - (Read-only) Type: `string`.
* `split_dns` - (Optional) Type: `string`.
* `split_include` - (Optional) Type: `string`.
* `src_address_list` - (Optional) Type: `string`.
* `static_dns` - (Optional) Type: `string`.
* `system_dns` - (Optional) Type: `bool`.
* `use_responder_dns` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_mode_config.example '*3'

# Named router
terraform import routeros_ip_ipsec_mode_config.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_mode_config.example 'home/my-resource-name'
```
