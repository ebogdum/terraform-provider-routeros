---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_attribute"
description: |-
  Requires user-manager package
---

# Resource: routeros_user_manager_attribute

Requires user-manager package

## Example Usage

```terraform
resource "routeros_user_manager_attribute" "attribute_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # default = "replace-me"
  # default_name = "replace-me"
  # name = "tf-example"
  # packet_types = "replace-me"
  # standard_name = "replace-me"
  # type_id = "replace-me"
  # value_type = "replace-me"
  # vendor_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `default` - (Optional) Type: `string`.
* `default_name` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`. Name of the attribute.
* `packet_types` - (Optional) Type: `string`. access-accept - use this attribute in RADIUS Access-Accept messages access-challenge - use this attribute in RADIUS Access-Challenge messages
* `standard_name` - (Optional) Type: `string`.
* `type_id` - (Optional) Type: `string`. Attribute identification number from the specific vendor's attribute database.
* `value_type` - (Optional) Type: `string`. hex ip-address - IPv4 or IPv6 IP address ip6-prefix - IPv6 prefix macro string uint32
* `vendor_id` - (Optional) Type: `string`. IANA allocated a specific enterprise identification number.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_manager_attribute.example '*3'

# Named router
terraform import routeros_user_manager_attribute.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_manager_attribute.example 'home/my-resource-name'
```
