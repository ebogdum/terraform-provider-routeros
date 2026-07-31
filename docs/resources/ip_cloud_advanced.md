---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_cloud_advanced"
description: |-
  Mirrors RouterOS /ip/cloud/advanced.
---

# Resource: routeros_ip_cloud_advanced

Mirrors RouterOS `/ip/cloud/advanced`.

## Example Usage

```terraform
resource "routeros_ip_cloud_advanced" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # use_local_address = true
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `use_local_address` - (Optional) Type: `bool`. RouterOS `use-local-address`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_cloud_advanced.this 'home'
```
