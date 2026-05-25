---
page_title: "RouterOS: routeros_ip_vrf"
description: |-
  VRF creation on CHR sometimes returns "interrupted" if attempted alongside other routing changes. Skipped from automated acc tests.
---

# Resource: routeros_ip_vrf

VRF creation on CHR sometimes returns "interrupted" if attempted alongside other routing changes. Skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_ip_vrf" "vrf_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # interfaces = "replace-me"
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `interfaces` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `builtin` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_vrf.example '*3'

# Named router
terraform import routeros_ip_vrf.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_vrf.example 'home/my-resource-name'
```
