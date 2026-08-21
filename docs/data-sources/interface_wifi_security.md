---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_security"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_security

Manages the RouterOS `/interface/wifi/security` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_security" "security_example" {
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
* `eap_password` - (Optional) Type: `string`.
* `eap_tls_certificate` - (Optional) Type: `string`.
* `eap_username` - (Optional) Type: `string`.
* `encryption` - (Optional) Type: `string`.
* `ft_enabled` - (Optional) Type: `string`.
* `ft_mobility_domain` - (Optional) Type: `string`.
* `ft_nas_identifier` - (Optional) Type: `string`.
* `ft_over_ds` - (Optional) Type: `string`.
* `ft_preserve_vlan_id` - (Optional) Type: `string`.
* `ft_r0_key_lifetime` - (Optional) Type: `string`.
* `ft_reassoc_deadline` - (Optional) Type: `string`.
* `group_encryption` - (Optional) Type: `string`.
* `group_key_update` - (Optional) Type: `string`.
* `management_encryption` - (Optional) Type: `string`.
* `management_protection` - (Optional) Type: `string`.
* `multi_passphrase_group` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_wsec`.
* `owe_transition_interface` - (Optional) Type: `string`.
* `passphrase` - (Optional) Type: `string`.
* `sae_anti_clogging_threshold` - (Optional) Type: `string`.
* `sae_max_failure_rate` - (Optional) Type: `string`.
* `sae_pwe` - (Optional) Type: `string`.
* `types` - (Optional) Type: `string`.
* `wps` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

