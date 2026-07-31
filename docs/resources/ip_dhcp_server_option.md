---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_option"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_server_option

Manages the RouterOS `/ip/dhcp-server/option` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_option" "option_example" {
  # router = "my-router"  # which router to target; omit for the default
  code = 60
  name = "tf-example"
  value = "'tf-acc'"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # force = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `code` - (Required) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `force` - (Optional) Type: `bool`.
* `name` - (Required) Type: `string`.
* `raw_value` - (Read-only) Type: `string`.
* `value` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server_option.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_option.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_option.example 'home/my-resource-name'
```
