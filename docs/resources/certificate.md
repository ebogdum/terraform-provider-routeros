---
subcategory: "Certificates"
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
  # acme = false
  # add_acme = "replace-me"
  # authority = false
  # card_reinstall = "replace-me"
  # card_verify = "replace-me"
  # common_name = "replace-me"
  # country = "replace-me"
  # create_cert_request = "replace-me"
  # crl = false
  # days_valid = 365
  # digest_algorithm = "md5"
  # dynamic = false
  # expired = false
  # export = "replace-me"
  # has_acme_status = "replace-me"
  # import = "replace-me"
  # issued = false
  # key_size = "2048"
  # key_usage = "109"
  # locality = "replace-me"
  # name = "tf-example"
  # notsealed = "replace-me"
  # organization = "replace-me"
  # private_key = "REDACTED"
  # revoke = "replace-me"
  # revoked = false
  # sealed = "replace-me"
  # sealed_and_hide = "replace-me"
  # sign = "replace-me"
  # sign_via_scep = "replace-me"
  # smart_card_key = "REDACTED"
  # state = "replace-me"
  # subject_alt_name = "replace-me"
  # trust_store = "4.294967295e+09"
  # trusted = true
  # type = 0
  # unit = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
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
