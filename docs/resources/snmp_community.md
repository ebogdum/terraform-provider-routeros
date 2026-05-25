---
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
  # authentication_protocol = "MD5"
  # encryption_password = "REDACTED"
  # encryption_protocol = "DES"
  # read_access = true
  # security = "none"
  # write_access = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `addresses` - (Optional) Type: `cidr`.
* `authentication_password` - (Optional) Type: `string`. **Sensitive.**
* `authentication_protocol` - (Optional) Type: `enum(MD5|SHA1)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `encryption_password` - (Optional) Type: `string`. **Sensitive.**
* `encryption_protocol` - (Optional) Type: `enum(DES|AES)`.
* `name` - (Required) Type: `string`. Default: `tf-acc-snmp-community`.
* `read_access` - (Optional) Type: `bool`. Default: `1`.
* `security` - (Optional) Type: `enum(none|authorized|private)`.
* `write_access` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

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
