---
subcategory: "Certificates"
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
* `acme` - (Optional) Type: `bool`.
* `add_acme` - (Optional) Type: `string`.
* `authority` - (Optional) Type: `bool`.
* `card_reinstall` - (Optional) Type: `string`.
* `card_verify` - (Optional) Type: `string`.
* `common_name` - (Optional) Type: `string`.
* `country` - (Optional) Type: `string`.
* `create_cert_request` - (Optional) Type: `string`.
* `crl` - (Optional) Type: `bool`.
* `days_valid` - (Optional) Type: `int`. Default: `365`.
* `digest_algorithm` - (Optional) Type: `enum(md5|sha1|sha256|sha384|sha512)`.
* `dynamic` - (Optional) Type: `bool`.
* `expired` - (Optional) Type: `bool`.
* `export` - (Optional) Type: `string`.
* `has_acme_status` - (Optional) Type: `string`.
* `import` - (Optional) Type: `string`.
* `issued` - (Optional) Type: `bool`.
* `key_size` - (Optional) Type: `enum(prime256v1|secp384r1|secp521r1|1024|1536|2048, ...)`. Default: `2048`.
* `key_usage` - (Optional) Type: `string`. Default: `109`.
* `locality` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `notsealed` - (Optional) Type: `string`.
* `organization` - (Optional) Type: `string`.
* `private_key` - (Optional) Type: `bool`.
* `revoke` - (Optional) Type: `string`.
* `revoked` - (Optional) Type: `bool`.
* `sealed` - (Optional) Type: `string`.
* `sealed_and_hide` - (Optional) Type: `string`.
* `sign` - (Optional) Type: `string`.
* `sign_via_scep` - (Optional) Type: `string`.
* `smart_card_key` - (Optional) Type: `bool`.
* `state` - (Optional) Type: `string`.
* `subject_alt_name` - (Optional) Type: `string`.
* `trust_store` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `trusted` - (Optional) Type: `bool`. Default: `1`.
* `type` - (Optional) Type: `int`.
* `unit` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `acme_status` - Type: `string`.
* `akid` - Type: `string`.
* `ca` - Type: `string`.
* `ca_crl_host` - Type: `string`.
* `ca_fingerprint` - Type: `string`.
* `expires_after` - Type: `string`.
* `fingerprint` - Type: `string`.
* `invalid_after` - Type: `string`.
* `invalid_before` - Type: `string`.
* `issuer` - Type: `string`.
* `key_type` - Type: `enum(rsa|dsa|ec)`.
* `req_fingerprint` - Type: `string`.
* `revoked_time` - Type: `string`.
* `scep_url` - Type: `string`.
* `serial_number` - Type: `string`.
* `skid` - Type: `string`.

