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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `code` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `raw_value` - (Read-only) Type: `string`.
* `value` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
