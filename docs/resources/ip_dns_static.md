---
subcategory: "DNS"
page_title: "RouterOS: routeros_ip_dns_static"
description: |-
  A DNS A/AAAA/CNAME/MX/... static entry. Requires either `name` or `regexp`.
---

# Resource: routeros_ip_dns_static

A DNS A/AAAA/CNAME/MX/... static entry. Requires either `name` or `regexp`.

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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`. IPv4/IPv6 address to return. Required when `type` is `"A"` or `"AAAA"` (the default); must be left unset for other types (`CNAME`, `FWD`, `MX`, `NS`, `SRV`, `TXT`, or a `regexp`-matched entry).
* `address_list` - (Optional) Type: `string`.
* `cname` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `forward_to` - (Optional) Type: `string`.
* `match_subdomain` - (Optional) Type: `string`.
* `mx_exchange` - (Optional) Type: `string`.
* `mx_preference` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`. FQDN matched against incoming queries. Provide this or `regexp`, not both.
* `ns` - (Optional) Type: `string`.
* `regexp` - (Optional) Type: `string`.
* `srv_port` - (Optional) Type: `string`.
* `srv_priority` - (Optional) Type: `string`.
* `srv_target` - (Optional) Type: `string`.
* `srv_weight` - (Optional) Type: `string`.
* `text` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`. Record type: `A`, `AAAA`, `CNAME`, `FWD`, `MX`, `NS`, `NXDOMAIN`, `SRV` or `TXT`. Defaults to `A`. Case-sensitive: RouterOS rejects a lower-case value.

## Attribute Reference

* `id` - RouterOS internal .id.


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
