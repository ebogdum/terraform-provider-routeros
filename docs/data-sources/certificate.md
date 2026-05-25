---
page_title: "RouterOS: routeros_certificate"
description: |-
  Certificate template/store. RouterOS issues self-signed CAs and CA-signed
---

# Data Source: routeros_certificate

Certificate template/store. RouterOS issues self-signed CAs and CA-signed
leaf certs in two steps: declare the template via this resource, then
trigger signing via the /certificate/sign action.


## Example Usage

```terraform
data "routeros_certificate" "certificate_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

