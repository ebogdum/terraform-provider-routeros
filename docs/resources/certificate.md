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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `acme` - (Read-only) Type: `bool`.
* `acme_status` - (Read-only) Type: `string`.
* `add_acme` - (Read-only) Type: `string`.
* `akid` - (Read-only) Type: `string`.
* `authority` - (Read-only) Type: `bool`.
* `ca` - (Read-only) Type: `string`.
* `ca_crl_host` - (Read-only) Type: `string`.
* `ca_fingerprint` - (Read-only) Type: `string`.
* `card_reinstall` - (Read-only) Type: `string`.
* `card_verify` - (Read-only) Type: `string`.
* `common_name` - (Optional) Type: `string`.
* `country` - (Optional) Type: `string`.
* `create_cert_request` - (Read-only) Type: `string`.
* `crl` - (Read-only) Type: `bool`.
* `days_valid` - (Optional) Type: `int`.
* `digest_algorithm` - (Optional) Type: `string`.
* `dynamic` - (Read-only) Type: `bool`.
* `expired` - (Read-only) Type: `bool`.
* `expires_after` - (Read-only) Type: `string`.
* `export` - (Read-only) Type: `string`.
* `fingerprint` - (Read-only) Type: `string`.
* `has_acme_status` - (Read-only) Type: `string`.
* `import` - (Read-only) Type: `string`.
* `invalid_after` - (Read-only) Type: `string`.
* `invalid_before` - (Read-only) Type: `string`.
* `issued` - (Read-only) Type: `bool`.
* `issuer` - (Read-only) Type: `string`.
* `key_size` - (Optional) Type: `string`.
* `key_type` - (Read-only) Type: `string`.
* `key_usage` - (Optional) Type: `string`.
* `locality` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `notsealed` - (Read-only) Type: `string`.
* `organization` - (Optional) Type: `string`.
* `private_key` - (Read-only) Type: `bool`. **Sensitive.**
* `req_fingerprint` - (Read-only) Type: `string`.
* `revoke` - (Read-only) Type: `string`.
* `revoked` - (Read-only) Type: `bool`.
* `revoked_time` - (Read-only) Type: `string`.
* `scep_url` - (Read-only) Type: `string`.
* `sealed` - (Read-only) Type: `string`.
* `sealed_and_hide` - (Read-only) Type: `string`.
* `serial_number` - (Read-only) Type: `string`.
* `sign` - (Read-only) Type: `string`.
* `sign_via_scep` - (Read-only) Type: `string`.
* `skid` - (Read-only) Type: `string`.
* `smart_card_key` - (Read-only) Type: `bool`.
* `state` - (Optional) Type: `string`.
* `subject_alt_name` - (Optional) Type: `string`.
* `trust_store` - (Optional) Type: `string`.
* `trusted` - (Optional) Type: `bool`.
* `type` - (Read-only) Type: `int`.
* `unit` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
