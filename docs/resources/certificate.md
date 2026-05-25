---
page_title: "RouterOS: routeros_certificate"
description: |-
  Certificate template/store. RouterOS issues self-signed CAs and CA-signed
---

# Resource: routeros_certificate

Certificate template/store. RouterOS issues self-signed CAs and CA-signed
leaf certs in two steps: declare the template via this resource, then
trigger signing via the /certificate/sign action.


## Example Usage

```terraform
resource "routeros_certificate" "certificate_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # common_name = "replace-me"
  # country = "replace-me"
  # days_valid = 365
  # digest_algorithm = "md5"
  # key_size = "2048"
  # key_usage = "109"
  # locality = "replace-me"
  # name = "tf-example"
  # organization = "replace-me"
  # state = "replace-me"
  # subject_alt_name = "replace-me"
  # trust_store = "4.294967295e+09"
  # trusted = true
  # unit = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `common_name` - (Optional) Type: `string`.
* `country` - (Optional) Type: `string`.
* `days_valid` - (Optional) Type: `int`. Default: `365`.
* `digest_algorithm` - (Optional) Type: `enum(md5|sha1|sha256|sha384|sha512)`.
* `key_size` - (Optional) Type: `enum(prime256v1|secp384r1|secp521r1|1024|1536|2048, ...)`. Default: `2048`.
* `key_usage` - (Optional) Type: `string`. Default: `109`.
* `locality` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `organization` - (Optional) Type: `string`.
* `state` - (Optional) Type: `string`.
* `subject_alt_name` - (Optional) Type: `string`.
* `trust_store` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `trusted` - (Optional) Type: `bool`. Default: `1`.
* `unit` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_certificate.example '*3'

# Named router
terraform import routeros_certificate.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_certificate.example 'home/my-resource-name'
```
