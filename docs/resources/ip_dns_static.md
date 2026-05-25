---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dns_static"
description: |-
  A DNS A/AAAA/CNAME/MX/... static entry. Requires either name OR regexp.
---

# Resource: routeros_ip_dns_static

A DNS A/AAAA/CNAME/MX/... static entry. Requires either name OR regexp.

## Example Usage

```terraform
resource "routeros_ip_dns_static" "static_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "127.0.0.1"
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ttl = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `string`. Address to return. Default: `127.0.0.1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`. FQDN matched against incoming queries. Default: `tf-acc-test.invalid`.
* `ttl` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dns_static.example '*3'

# Named router
terraform import routeros_ip_dns_static.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dns_static.example 'home/my-resource-name'
```
