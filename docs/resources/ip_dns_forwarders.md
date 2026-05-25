---
subcategory: "DNS"
page_title: "RouterOS: routeros_ip_dns_forwarders"
description: |-
  Discovered; required address must be a valid resolvable IP
---

# Resource: routeros_ip_dns_forwarders

Discovered; required address must be a valid resolvable IP

## Example Usage

```terraform
resource "routeros_ip_dns_forwarders" "forwarders_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dns_servers = "replace-me"
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dns_servers` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dns_forwarders.example '*3'

# Named router
terraform import routeros_ip_dns_forwarders.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dns_forwarders.example 'home/my-resource-name'
```
