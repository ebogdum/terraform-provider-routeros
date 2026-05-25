---
page_title: "RouterOS: routeros_interface_macsec"
description: |-
  MACsec needs a base interface; skipped on hardware without a suitable parent.
---

# Resource: routeros_interface_macsec

MACsec needs a base interface; skipped on hardware without a suitable parent.

## Example Usage

```terraform
resource "routeros_interface_macsec" "macsec_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # mtu = "replace-me"
  # name = "tf-example"
  # profile = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `profile` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_macsec.example '*3'

# Named router
terraform import routeros_interface_macsec.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_macsec.example 'home/my-resource-name'
```
