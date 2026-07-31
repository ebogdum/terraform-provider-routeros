---
subcategory: "PPP"
page_title: "RouterOS: routeros_ppp_l2tp_secret"
description: |-
  L2TP CHAP/PAP shared secret entry. Schema varies across ROS releases. Skipped from automated acc tests.
---

# Resource: routeros_ppp_l2tp_secret

L2TP CHAP/PAP shared secret entry. Schema varies across ROS releases. Skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_ppp_l2tp_secret" "l2tp_secret_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # secret = "REDACTED"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `secret` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ppp_l2tp_secret.example '*3'

# Named router
terraform import routeros_ppp_l2tp_secret.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ppp_l2tp_secret.example 'home/my-resource-name'
```
