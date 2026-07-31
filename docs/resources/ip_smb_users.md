---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_smb_users"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_ip_smb_users

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_ip_smb_users" "users_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # password = "REDACTED"
  # read_only = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `read_only` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_smb_users.example '*3'

# Named router
terraform import routeros_ip_smb_users.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_smb_users.example 'home/my-resource-name'
```
