---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_key_qkd"
description: |-
  Mirrors RouterOS /ip/ipsec/key/qkd.
---

# Resource: routeros_ip_ipsec_key_qkd

Mirrors RouterOS `/ip/ipsec/key/qkd`.

## Example Usage

```terraform
resource "routeros_ip_ipsec_key_qkd" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # cache_size = 0
  # certificate = "replace-me"
  # enabled = true
  # key_size = 0
  # kme_id = "replace-me"
  # peer_sae_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`. RouterOS `address`.
* `cache_size` - (Optional) Type: `int`. RouterOS `cache-size`.
* `certificate` - (Optional) Type: `string`. RouterOS `certificate`.
* `enabled` - (Optional) Type: `bool`. RouterOS `enabled`.
* `key_size` - (Optional) Type: `int`. RouterOS `key-size`.
* `kme_id` - (Optional) Type: `string`. RouterOS `kme-id`.
* `peer_sae_id` - (Optional) Type: `string`. RouterOS `peer-sae-id`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_ip_ipsec_key_qkd.example 'home::*3'
```
