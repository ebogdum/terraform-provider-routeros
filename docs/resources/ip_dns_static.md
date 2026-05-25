---
subcategory: "DNS"
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
  # address_list = "replace-me"
  # cname = "replace-me"
  # forward_to = "replace-me"
  # match_subdomain = "replace-me"
  # mx_exchange = "replace-me"
  # mx_preference = "replace-me"
  # ns = "replace-me"
  # regexp = "replace-me"
  # srv_port = "443"
  # srv_priority = "replace-me"
  # srv_target = "replace-me"
  # srv_weight = "replace-me"
  # text = "replace-me"
  # ttl = "replace-me"
  # type = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `string`. Address to return. Default: `127.0.0.1`.
* `address_list` - (Optional) Type: `string`.
* `cname` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `forward_to` - (Optional) Type: `string`.
* `match_subdomain` - (Optional) Type: `string`.
* `mx_exchange` - (Optional) Type: `string`.
* `mx_preference` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. FQDN matched against incoming queries. Default: `tf-acc-test.invalid`.
* `ns` - (Optional) Type: `string`.
* `regexp` - (Optional) Type: `string`.
* `srv_port` - (Optional) Type: `string`.
* `srv_priority` - (Optional) Type: `string`.
* `srv_target` - (Optional) Type: `string`.
* `srv_weight` - (Optional) Type: `string`.
* `text` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`.

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
