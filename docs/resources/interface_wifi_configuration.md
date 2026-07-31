---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_configuration"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_configuration

Manages the RouterOS `/interface/wifi/configuration` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # x2g_probe_delay = "replace-me"
  # x3gpp_info = "replace-me"
  # x3gpp_info_raw = "replace-me"
  # aaa = "replace-me"
  # antenna_gain = "replace-me"
  # authentication_types = "replace-me"
  # band = "replace-me"
  # beacon_interval = "replace-me"
  # beacon_protection = "replace-me"
  # bridge = "bridge1"
  # bridge_cost = "replace-me"
  # bridge_horizon = "replace-me"
  # called_format = "replace-me"
  # calling_format = "replace-me"
  # chains = "replace-me"
  # channel = "replace-me"
  # channel_width = "replace-me"
  # ciphers = "tkip"
  # client_isolation = "replace-me"
  # connect_group = "replace-me"
  # connect_priority = "replace-me"
  # connection_capabilities = "replace-me"
  # country = "replace-me"
  # datapath = "replace-me"
  # deprioritize_unii_3_4 = "replace-me"
  # dgaf = "replace-me"
  # dh_groups = "replace-me"
  # disable_pmkid = "replace-me"
  # distance = "replace-me"
  # domain_names = "replace-me"
  # dtim_period = "replace-me"
  # eap_accounting = "replace-me"
  # eap_anonymous_identity = "replace-me"
  # eap_certificate_mode = "replace-me"
  # eap_methods = "replace-me"
  # eap_password = "REDACTED"
  # eap_tls_certificate = "replace-me"
  # eap_username = "replace-me"
  # esr = "replace-me"
  # frequency = "replace-me"
  # ft_enabled = "replace-me"
  # ft_mobility_domain = "replace-me"
  # ft_nas_identifier = "replace-me"
  # ft_over_ds = "replace-me"
  # ft_preserve_vlan_id = "replace-me"
  # ft_r0_key_lifetime = "replace-me"
  # ft_reassoc_deadline = "replace-me"
  # group_encryption = "replace-me"
  # group_key_update = "replace-me"
  # hessid = "replace-me"
  # hide_ssid = "replace-me"
  # hotspot_2_0 = "replace-me"
  # hw_protection_mode = "replace-me"
  # installation = "replace-me"
  # interface_list = "replace-me"
  # interim_update = "replace-me"
  # internet = "replace-me"
  # interworking = "replace-me"
  # ipv4_availability = "replace-me"
  # ipv6_availability = "replace-me"
  # mac_caching = "replace-me"
  # management_encryption = "replace-me"
  # management_protection = "replace-me"
  # manager = "replace-me"
  # max_clients = "replace-me"
  # max_tx_power = "replace-me"
  # mode = "replace-me"
  # multi_passphrase_group = "replace-me"
  # multicast_enhance = "replace-me"
  # nas_identifier = "replace-me"
  # neighbor_group = "replace-me"
  # network_type = "replace-me"
  # open_flow_switch = "replace-me"
  # openflow = "replace-me"
  # operational_classes = "replace-me"
  # operator_names = "replace-me"
  # owe_transition_interface = "replace-me"
  # passphrase = "replace-me"
  # password_format = "replace-me"
  # qo_s_classifier = "replace-me"
  # realms = "replace-me"
  # realms_raw = "replace-me"
  # reselect_interval = "replace-me"
  # reselect_time = "replace-me"
  # roaming_ois = "replace-me"
  # rrm = "replace-me"
  # sae_anti_clogging_threshold = "replace-me"
  # sae_max_failure_rate = "replace-me"
  # sae_pwe = "replace-me"
  # secondary_frequency = "replace-me"
  # security = "replace-me"
  # skip_dfs_channels = "replace-me"
  # ssid = "replace-me"
  # station_roaming = "replace-me"
  # steering = "replace-me"
  # traffic_processing = "replace-me"
  # transition_request_count = "replace-me"
  # transition_threshold = "replace-me"
  # transition_threshold_period = "replace-me"
  # transition_threshold_time = "replace-me"
  # transition_time = "replace-me"
  # tx_chains = "replace-me"
  # tx_power = "replace-me"
  # types = "wpa-psk"
  # uesa = "replace-me"
  # username_format = "replace-me"
  # venue = "replace-me"
  # venue_names = "replace-me"
  # vlan_id = "replace-me"
  # wan_at_capacity = "replace-me"
  # wan_downlink = "replace-me"
  # wan_downlink_load = "replace-me"
  # wan_measurement_duration = "replace-me"
  # wan_status = "replace-me"
  # wan_symmetric = "replace-me"
  # wan_uplink = "replace-me"
  # wan_uplink_load = "replace-me"
  # wnm = "replace-me"
  # wps = "replace-me"
}
```

## Sub-object properties

RouterOS exposes this menu's sub-objects under **dotted** REST names --
`configuration.ssid`, `security.passphrase`, `channel.band` and so on. The
provider keeps them as flat Terraform attributes (`ssid`, `passphrase`, `band`)
and maps each one to its dotted wire name for you, so configuration written
against these attributes now reaches the device.

In this menu `ssid`, `mode`, `country` and `manager` are genuine top-level
properties, not sub-object members -- only `security.*`, `datapath.*`,
`channel.*`, `aaa.*` and `steering.*` are dotted. The section names
(`security`, `datapath`, `channel`, ...) remain top-level references to a named
profile.

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
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
* `eap_password` - (Optional) Type: `string`. **Sensitive.**
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
* `name` - (Required) Type: `string`.
* `nas_identifier` - (Optional) Type: `string`.
* `neighbor_group` - (Optional) Type: `string`.
* `network_type` - (Optional) Type: `string`.
* `open_flow_switch` - (Optional) Type: `string`.
* `openflow` - (Optional) Type: `string`.
* `operational_classes` - (Optional) Type: `string`.
* `operator_names` - (Optional) Type: `string`.
* `owe_transition_interface` - (Optional) Type: `string`.
* `passphrase` - (Optional) Type: `string`. **Sensitive.**
* `password_format` - (Optional) Type: `string`.
* `qo_s_classifier` - (Optional) Type: `string`.
* `qos_classifier` - (Optional) Type: `string`. RouterOS `qos-classifier`.
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
* `x2g_probe_delay` - (Optional) Type: `string`.
* `x3gpp_info` - (Optional) Type: `string`.
* `x3gpp_info_raw` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_configuration.example '*3'

# Named router
terraform import routeros_interface_wifi_configuration.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_configuration.example 'home/my-resource-name'
```
