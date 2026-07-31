---
subcategory: "SNMP"
page_title: "RouterOS: routeros_snmp_community"
description: |-
  RouterOS resource.
---

# Resource: routeros_snmp_community

Manages the RouterOS `/snmp/community` menu.

## Example Usage

```terraform
resource "routeros_snmp_community" "community_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # addresses = "10.99.0.0/24"
  # authentication_password = "REDACTED"
  # authentication_protocol = "md5"
  # encryption_password = "REDACTED"
  # encryption_protocol = "des"
  # read_access = true
  # security = "none"
  # write_access = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `addresses` - (Optional) Type: `string`.
* `authentication_password` - (Optional) Type: `string`. **Sensitive.**
* `authentication_protocol` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `encryption_password` - (Optional) Type: `string`. **Sensitive.**
* `encryption_protocol` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `read_access` - (Optional) Type: `bool`.
* `security` - (Optional) Type: `string`.
* `write_access` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_snmp_community.example '*3'

# Named router
terraform import routeros_snmp_community.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_snmp_community.example 'home/my-resource-name'
```
