---
subcategory: "PPP"
page_title: "RouterOS: routeros_ppp_secret"
description: |-
  RouterOS resource.
---

# Resource: routeros_ppp_secret

Manages the RouterOS `/ppp/secret` menu.

## Example Usage

```terraform
resource "routeros_ppp_secret" "secret_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # caller_id = "replace-me"
  # ipv6_routes = "replace-me"
  # limit_bytes_in = "replace-me"
  # limit_bytes_out = "replace-me"
  # local_address = "10.99.0.1"
  # password = "REDACTED"
  # profile = "replace-me"
  # remote_address = "10.99.0.1"
  # remote_ipv6_prefix = "replace-me"
  # routes = "replace-me"
  # service = "any"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `caller_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipv6_routes` - (Optional) Type: `string`.
* `limit_bytes_in` - (Optional) Type: `string`.
* `limit_bytes_out` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_pppsec`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_ipv6_prefix` - (Optional) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `service` - (Optional) Type: `enum(any|async|pptp|pppoe|l2tp|ovpn, ...)`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ppp_secret.example '*3'

# Named router
terraform import routeros_ppp_secret.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ppp_secret.example 'home/my-resource-name'
```
