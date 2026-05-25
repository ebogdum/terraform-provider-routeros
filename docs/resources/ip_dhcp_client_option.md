---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_client_option"
description: |-
  DHCP client option requires an existing dhcp-client + named option. Skipped.
---

# Resource: routeros_ip_dhcp_client_option

DHCP client option requires an existing dhcp-client + named option. Skipped.

## Example Usage

```terraform
resource "routeros_ip_dhcp_client_option" "option_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # code = 0
  # name = "tf-example"
  # value = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `code` - (Optional) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `name` - (Optional) Type: `string`.
* `value` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `raw_value` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_client_option.example '*3'

# Named router
terraform import routeros_ip_dhcp_client_option.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_client_option.example 'home/my-resource-name'
```
