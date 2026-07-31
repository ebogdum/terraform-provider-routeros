---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_identity"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_ip_ipsec_identity

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_ip_ipsec_identity" "identity_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auth_method = "rsa-key"
  # generate_policy = "no"
  # match_by = "certificate"
  # mode_configuration = "4.294967295e+09"
  # my_id_type = "fqdn"
  # notrack_chain = "replace-me"
  # peer = "replace-me"
  # policy_template_group = "replace-me"
  # remote_id_type = "fqdn"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `auth_method` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`. RouterOS `certificate`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `eap_methods` - (Optional) Type: `string`. RouterOS `eap-methods`.
* `generate_policy` - (Optional) Type: `string`.
* `key` - (Optional) Type: `string`. RouterOS `key`.
* `match_by` - (Optional) Type: `string`.
* `mode_config` - (Optional) Type: `string`. RouterOS `mode-config`.
* `mode_configuration` - (Read-only) Type: `string`.
* `my_id` - (Optional) Type: `string`. RouterOS `my-id`.
* `my_id_type` - (Read-only) Type: `string`.
* `notrack_chain` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. RouterOS `password`. **Sensitive.**
* `peer` - (Optional) Type: `string`.
* `policy_template_group` - (Optional) Type: `string`.
* `remote_certificate` - (Optional) Type: `string`. RouterOS `remote-certificate`.
* `remote_id` - (Optional) Type: `string`. RouterOS `remote-id`.
* `remote_id_type` - (Read-only) Type: `string`.
* `remote_key` - (Optional) Type: `string`. RouterOS `remote-key`.
* `secret` - (Optional) Type: `string`. RouterOS `secret`. **Sensitive.**
* `username` - (Optional) Type: `string`. RouterOS `username`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_identity.example '*3'

# Named router
terraform import routeros_ip_ipsec_identity.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_identity.example 'home/my-resource-name'
```
