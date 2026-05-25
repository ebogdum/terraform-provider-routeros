---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_configuration"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_configuration

Manages the RouterOS `/interface/wifi/configuration` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_configuration" "configuration_example" {
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
* `x2g_probe_delay` - (Optional) Type: `string`.
* `x3gpp_info` - (Optional) Type: `string`.
* `x3gpp_info_raw` - (Optional) Type: `string`.
* `aaa` - (Optional) Type: `string`.
* `antenna_gain` - (Optional) Type: `string`.
* `authentication_types` - (Optional) Type: `string`.
* `band` - (Optional) Type: `string`.
* `beacon_interval` - (Optional) Type: `string`.
* `beacon_protection` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `bridge_cost` - (Optional) Type: `string`.
* `bridge_horizon` - (Optional) Type: `string`.
* `called_format` - (Optional) Type: `string`.
* `calling_format` - (Optional) Type: `string`.
* `chains` - (Optional) Type: `string`.
* `channel` - (Optional) Type: `string`.
* `channel_width` - (Optional) Type: `string`.
* `ciphers` - (Optional) Type: `string`.
* `client_isolation` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_group` - (Optional) Type: `string`.
* `connect_priority` - (Optional) Type: `string`.
* `connection_capabilities` - (Optional) Type: `string`.
* `country` - (Optional) Type: `string`.
* `datapath` - (Optional) Type: `string`.
* `deprioritize_unii_3_4` - (Optional) Type: `string`.
* `dgaf` - (Optional) Type: `string`.
* `dh_groups` - (Optional) Type: `string`.
* `disable_pmkid` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `distance` - (Optional) Type: `string`.
* `domain_names` - (Optional) Type: `string`.
* `dtim_period` - (Optional) Type: `string`.
* `eap_accounting` - (Optional) Type: `string`.
* `eap_anonymous_identity` - (Optional) Type: `string`.
* `eap_certificate_mode` - (Optional) Type: `string`.
* `eap_methods` - (Optional) Type: `string`.
* `eap_password` - (Optional) Type: `string`.
* `eap_tls_certificate` - (Optional) Type: `string`.
* `eap_username` - (Optional) Type: `string`.
* `esr` - (Optional) Type: `string`.
* `frequency` - (Optional) Type: `string`.
* `ft_enabled` - (Optional) Type: `string`.
* `ft_mobility_domain` - (Optional) Type: `string`.
* `ft_nas_identifier` - (Optional) Type: `string`.
* `ft_over_ds` - (Optional) Type: `string`.
* `ft_preserve_vlan_id` - (Optional) Type: `string`.
* `ft_r0_key_lifetime` - (Optional) Type: `string`.
* `ft_reassoc_deadline` - (Optional) Type: `string`.
* `group_encryption` - (Optional) Type: `string`.
* `group_key_update` - (Optional) Type: `string`.
* `hessid` - (Optional) Type: `string`.
* `hide_ssid` - (Optional) Type: `string`.
* `hotspot_2_0` - (Optional) Type: `string`.
* `hw_protection_mode` - (Optional) Type: `string`.
* `installation` - (Optional) Type: `string`.
* `interface_list` - (Optional) Type: `string`.
* `interim_update` - (Optional) Type: `string`.
* `internet` - (Optional) Type: `string`.
* `interworking` - (Optional) Type: `string`.
* `ipv4_availability` - (Optional) Type: `string`.
* `ipv6_availability` - (Optional) Type: `string`.
* `mac_caching` - (Optional) Type: `string`.
* `management_encryption` - (Optional) Type: `string`.
* `management_protection` - (Optional) Type: `string`.
* `manager` - (Optional) Type: `string`.
* `max_clients` - (Optional) Type: `string`.
* `max_tx_power` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `multi_passphrase_group` - (Optional) Type: `string`.
* `multicast_enhance` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_wcfg`.
* `nas_identifier` - (Optional) Type: `string`.
* `neighbor_group` - (Optional) Type: `string`.
* `network_type` - (Optional) Type: `string`.
* `open_flow_switch` - (Optional) Type: `string`.
* `openflow` - (Optional) Type: `string`.
* `operational_classes` - (Optional) Type: `string`.
* `operator_names` - (Optional) Type: `string`.
* `owe_transition_interface` - (Optional) Type: `string`.
* `passphrase` - (Optional) Type: `string`.
* `password_format` - (Optional) Type: `string`.
* `qo_s_classifier` - (Optional) Type: `string`.
* `realms` - (Optional) Type: `string`.
* `realms_raw` - (Optional) Type: `string`.
* `reselect_interval` - (Optional) Type: `string`.
* `reselect_time` - (Optional) Type: `string`.
* `roaming_ois` - (Optional) Type: `string`.
* `rrm` - (Optional) Type: `string`.
* `sae_anti_clogging_threshold` - (Optional) Type: `string`.
* `sae_max_failure_rate` - (Optional) Type: `string`.
* `sae_pwe` - (Optional) Type: `string`.
* `secondary_frequency` - (Optional) Type: `string`.
* `security` - (Optional) Type: `string`.
* `skip_dfs_channels` - (Optional) Type: `string`.
* `ssid` - (Optional) Type: `string`.
* `station_roaming` - (Optional) Type: `string`.
* `steering` - (Optional) Type: `string`.
* `traffic_processing` - (Optional) Type: `string`.
* `transition_request_count` - (Optional) Type: `string`.
* `transition_threshold` - (Optional) Type: `string`.
* `transition_threshold_period` - (Optional) Type: `string`.
* `transition_threshold_time` - (Optional) Type: `string`.
* `transition_time` - (Optional) Type: `string`.
* `tx_chains` - (Optional) Type: `string`.
* `tx_power` - (Optional) Type: `string`.
* `types` - (Optional) Type: `string`.
* `uesa` - (Optional) Type: `string`.
* `username_format` - (Optional) Type: `string`.
* `venue` - (Optional) Type: `string`.
* `venue_names` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.
* `wan_at_capacity` - (Optional) Type: `string`.
* `wan_downlink` - (Optional) Type: `string`.
* `wan_downlink_load` - (Optional) Type: `string`.
* `wan_measurement_duration` - (Optional) Type: `string`.
* `wan_status` - (Optional) Type: `string`.
* `wan_symmetric` - (Optional) Type: `string`.
* `wan_uplink` - (Optional) Type: `string`.
* `wan_uplink_load` - (Optional) Type: `string`.
* `wnm` - (Optional) Type: `string`.
* `wps` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

