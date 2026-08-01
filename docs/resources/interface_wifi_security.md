---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_security"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_security

Manages the RouterOS `/interface/wifi/security` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_security" "security_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # beacon_protection = "replace-me"
  # ciphers = "tkip"
  # connect_group = "replace-me"
  # connect_priority = "replace-me"
  # dh_groups = "replace-me"
  # disable_pmkid = "replace-me"
  # eap_accounting = "replace-me"
  # eap_anonymous_identity = "replace-me"
  # eap_certificate_mode = "replace-me"
  # eap_methods = "replace-me"
  # eap_password = "REDACTED"
  # eap_tls_certificate = "replace-me"
  # eap_username = "replace-me"
  # encryption = "replace-me"
  # ft_enabled = "replace-me"
  # ft_mobility_domain = "replace-me"
  # ft_nas_identifier = "replace-me"
  # ft_over_ds = "replace-me"
  # ft_r0_key_lifetime = "replace-me"
  # ft_reassoc_deadline = "replace-me"
  # group_encryption = "replace-me"
  # group_key_update = "replace-me"
  # management_encryption = "replace-me"
  # management_protection = "replace-me"
  # multi_passphrase_group = "replace-me"
  # owe_transition_interface = "replace-me"
  # passphrase = "replace-me"
  # sae_anti_clogging_threshold = "replace-me"
  # sae_max_failure_rate = "replace-me"
  # sae_pwe = "replace-me"
  # types = "wpa-psk"
  # wps = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `authentication_types` - (Optional) Type: `string`. RouterOS `authentication-types`.
* `beacon_protection` - (Optional) Type: `string`.
* `ciphers` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_group` - (Optional) Type: `string`.
* `connect_priority` - (Optional) Type: `string`.
* `dh_groups` - (Optional) Type: `string`.
* `disable_pmkid` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `eap_accounting` - (Optional) Type: `string`.
* `eap_anonymous_identity` - (Optional) Type: `string`.
* `eap_certificate_mode` - (Optional) Type: `string`.
* `eap_methods` - (Optional) Type: `string`.
* `eap_password` - (Optional) Type: `string`. **Sensitive.**
* `eap_tls_certificate` - (Optional) Type: `string`.
* `eap_username` - (Optional) Type: `string`.
* `encryption` - (Optional) Type: `string`.
* `ft` - (Optional) Type: `string`. RouterOS `ft`.
* `ft_enabled` - (Optional) Type: `string`.
* `ft_mobility_domain` - (Optional) Type: `string`.
* `ft_nas_identifier` - (Optional) Type: `string`.
* `ft_over_ds` - (Optional) Type: `string`.
* `ft_preserve_vlanid` - (Optional) Type: `string`. RouterOS `ft-preserve-vlanid`.
* `ft_r0_key_lifetime` - (Optional) Type: `string`.
* `ft_reassoc_deadline` - (Optional) Type: `string`.
* `ft_reassociation_deadline` - (Optional) Type: `string`. RouterOS `ft-reassociation-deadline`.
* `group_encryption` - (Optional) Type: `string`.
* `group_key_update` - (Optional) Type: `string`.
* `management_encryption` - (Optional) Type: `string`.
* `management_protection` - (Optional) Type: `string`.
* `multi_passphrase_group` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `owe_transition_interface` - (Optional) Type: `string`.
* `passphrase` - (Optional) Type: `string`. **Sensitive.**
* `sae_anti_clogging_threshold` - (Optional) Type: `string`.
* `sae_max_failure_rate` - (Optional) Type: `string`.
* `sae_pwe` - (Optional) Type: `string`.
* `types` - (Optional) Type: `string`.
* `wps` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_security.example '*3'

# Named router
terraform import routeros_interface_wifi_security.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_security.example 'home/my-resource-name'
```
