---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_interface"
description: |-
  CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.
---

# Resource: routeros_caps_man_interface

CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.

## Example Usage

```terraform
resource "routeros_caps_man_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp_timeout = "replace-me"
  # mac_address = "10.99.0.0/24"
  # master_interface = "ether1"
  # name = "tf-example"
  # radio_mac = "02:00:00:00:00:01"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `arp` - (Optional) Type: `string`. RouterOS `arp`.
* `arp_timeout` - (Optional) Type: `string`.
* `authentication_types` - (Optional) Type: `string`. RouterOS `authentication-types`.
* `band` - (Optional) Type: `string`. RouterOS `band`.
* `basic` - (Optional) Type: `string`. RouterOS `basic`.
* `bridge` - (Optional) Type: `string`. RouterOS `bridge`.
* `bridge_cost` - (Optional) Type: `string`. RouterOS `bridge-cost`.
* `bridge_horizon` - (Optional) Type: `string`. RouterOS `bridge-horizon`.
* `channel` - (Optional) Type: `string`. RouterOS `channel`.
* `client_to_client_forwarding` - (Optional) Type: `string`. RouterOS `client-to-client-forwarding`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `configuration` - (Optional) Type: `string`. RouterOS `configuration`.
* `control_channel_width` - (Optional) Type: `string`. RouterOS `control-channel-width`.
* `country` - (Optional) Type: `string`. RouterOS `country`.
* `datapath` - (Optional) Type: `string`. RouterOS `datapath`.
* `disable_pmkid` - (Optional) Type: `string`. RouterOS `disable-pmkid`.
* `disable_running_check` - (Optional) Type: `string`. RouterOS `disable-running-check`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `disconnect_timeout` - (Optional) Type: `string`. RouterOS `disconnect-timeout`.
* `distance` - (Optional) Type: `string`. RouterOS `distance`.
* `eap_methods` - (Optional) Type: `string`. RouterOS `eap-methods`.
* `eap_radius_accounting` - (Optional) Type: `string`. RouterOS `eap-radius-accounting`.
* `encryption` - (Optional) Type: `string`. RouterOS `encryption`.
* `extension_channel` - (Optional) Type: `string`. RouterOS `extension-channel`.
* `frame_lifetime` - (Optional) Type: `string`. RouterOS `frame-lifetime`.
* `frequency` - (Optional) Type: `string`. RouterOS `frequency`.
* `group_encryption` - (Optional) Type: `string`. RouterOS `group-encryption`.
* `group_key_update` - (Optional) Type: `string`. RouterOS `group-key-update`.
* `guard_interval` - (Optional) Type: `string`. RouterOS `guard-interval`.
* `hide_ssid` - (Optional) Type: `string`. RouterOS `hide-ssid`.
* `ht_basic_mcs` - (Optional) Type: `string`. RouterOS `ht-basic-mcs`.
* `ht_supported_mcs` - (Optional) Type: `string`. RouterOS `ht-supported-mcs`.
* `hw_protection_mode` - (Optional) Type: `string`. RouterOS `hw-protection-mode`.
* `hw_retries` - (Optional) Type: `string`. RouterOS `hw-retries`.
* `installation` - (Optional) Type: `string`. RouterOS `installation`.
* `interface_list` - (Optional) Type: `string`. RouterOS `interface-list`.
* `keepalive_frames` - (Optional) Type: `string`. RouterOS `keepalive-frames`.
* `l2mtu` - (Optional) Type: `string`. RouterOS `l2mtu`.
* `load_balancing_group` - (Optional) Type: `string`. RouterOS `load-balancing-group`.
* `local_forwarding` - (Optional) Type: `string`. RouterOS `local-forwarding`.
* `mac_address` - (Optional) Type: `string`.
* `master_interface` - (Optional) Type: `string`.
* `max_sta_count` - (Optional) Type: `string`. RouterOS `max-sta-count`.
* `mode` - (Optional) Type: `string`. RouterOS `mode`.
* `mtu` - (Optional) Type: `string`. RouterOS `mtu`.
* `multicast_helper` - (Optional) Type: `string`. RouterOS `multicast-helper`.
* `name` - (Optional) Type: `string`.
* `openflow_switch` - (Optional) Type: `string`. RouterOS `openflow-switch`.
* `passphrase` - (Optional) Type: `string`. RouterOS `passphrase`. **Sensitive.**
* `radio_mac` - (Optional) Type: `string`.
* `radio_name` - (Optional) Type: `string`. RouterOS `radio-name`.
* `rates` - (Optional) Type: `string`. RouterOS `rates`.
* `reselect_interval` - (Optional) Type: `string`. RouterOS `reselect-interval`.
* `rx_chains` - (Optional) Type: `string`. RouterOS `rx-chains`.
* `save_selected` - (Optional) Type: `string`. RouterOS `save-selected`.
* `secondary_frequency` - (Optional) Type: `string`. RouterOS `secondary-frequency`.
* `security` - (Optional) Type: `string`. RouterOS `security`.
* `skip_dfs_channels` - (Optional) Type: `string`. RouterOS `skip-dfs-channels`.
* `ssid` - (Optional) Type: `string`. RouterOS `ssid`.
* `supported` - (Optional) Type: `string`. RouterOS `supported`.
* `tls_certificate` - (Optional) Type: `string`. RouterOS `tls-certificate`.
* `tls_mode` - (Optional) Type: `string`. RouterOS `tls-mode`.
* `tx_chains` - (Optional) Type: `string`. RouterOS `tx-chains`.
* `tx_power` - (Optional) Type: `string`. RouterOS `tx-power`.
* `vht_basic_mcs` - (Optional) Type: `string`. RouterOS `vht-basic-mcs`.
* `vht_supported_mcs` - (Optional) Type: `string`. RouterOS `vht-supported-mcs`.
* `vlan_id` - (Optional) Type: `string`. RouterOS `vlan-id`.
* `vlan_mode` - (Optional) Type: `string`. RouterOS `vlan-mode`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_interface.example '*3'

# Named router
terraform import routeros_caps_man_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_interface.example 'home/my-resource-name'
```
