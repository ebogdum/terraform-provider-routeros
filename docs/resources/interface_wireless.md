---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_wireless"
description: |-
  802.11abgn interface needs a master-interface from the physical wireless hardware. Skipped on CHR.
---

# Resource: routeros_interface_wireless

802.11abgn interface needs a master-interface from the physical wireless hardware. Skipped on CHR.

## Example Usage

```terraform
resource "routeros_interface_wireless" "wireless_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # adaptive_noise_immunity = "replace-me"
  # allow_sharedkey = "replace-me"
  # ampdu_priorities = "replace-me"
  # amsdu_limit = "replace-me"
  # amsdu_threshold = "replace-me"
  # antenna_gain = "replace-me"
  # antenna_mode = "replace-me"
  # area = "replace-me"
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # band = "replace-me"
  # basic_rates_b = "replace-me"
  # bridge_mode = "replace-me"
  # burst_time = "replace-me"
  # channel_width = "replace-me"
  # compression = "replace-me"
  # country = "replace-me"
  # default_ap_tx_limit = "replace-me"
  # default_authentication = "replace-me"
  # default_client_tx_limit = "replace-me"
  # default_forwarding = "replace-me"
  # disable_running_check = "replace-me"
  # disconnect_timeout = "replace-me"
  # distance = "replace-me"
  # frame_lifetime = "replace-me"
  # frequency = "replace-me"
  # frequency_mode = "replace-me"
  # frequency_offset = "replace-me"
  # guard_interval = "replace-me"
  # hide_ssid = "replace-me"
  # ht_basic_mcs = "replace-me"
  # ht_supported_mcs = "replace-me"
  # hw_fragmentation_threshold = "replace-me"
  # hw_protection_mode = "replace-me"
  # hw_protection_threshold = "replace-me"
  # hw_retries = "replace-me"
  # installation = "replace-me"
  # interworking_profile = "replace-me"
  # keepalive_frames = "replace-me"
  # l2mtu = "replace-me"
  # mac_address = "10.99.0.0/24"
  # master_interface = "ether1"
  # max_station_count = "replace-me"
  # mode = "replace-me"
  # mtu = "replace-me"
  # multicast_buffering = "replace-me"
  # multicast_helper = "replace-me"
  # name = "example"
  # noise_floor_threshold = "replace-me"
  # nv2_cell_radius = "replace-me"
  # nv2_noise_floor_offset = "replace-me"
  # nv2_preshared_key = "REDACTED"
  # nv2_qos = "replace-me"
  # nv2_queue_count = "replace-me"
  # nv2_security = "replace-me"
  # on_fail_retry_time = "replace-me"
  # periodic_calibration = "replace-me"
  # periodic_calibration_interval = "replace-me"
  # preamble_mode = "replace-me"
  # prism_cardtype = "replace-me"
  # proprietary_extensions = "replace-me"
  # radio_name = "replace-me"
  # rate_selection = "replace-me"
  # rate_set = "replace-me"
  # rx_chains = "replace-me"
  # scan_list = "replace-me"
  # secondary_channel = "replace-me"
  # security_profile = "replace-me"
  # skip_dfs_channels = "replace-me"
  # ssid = "replace-me"
  # station_bridge_clone_mac = "replace-me"
  # station_roaming = "replace-me"
  # supported_rates_b = "replace-me"
  # tdma_period_size = "replace-me"
  # tx_chains = "replace-me"
  # tx_power = "replace-me"
  # tx_power_mode = "replace-me"
  # update_stats_interval = "replace-me"
  # vht_basic_mcs = "replace-me"
  # vht_supported_mcs = "replace-me"
  # wds_cost_range = "replace-me"
  # wds_default_bridge = "replace-me"
  # wds_default_cost = "replace-me"
  # wds_ignore_ssid = "replace-me"
  # wds_mode = "replace-me"
  # wireless_protocol = "replace-me"
  # wmm_support = "replace-me"
  # wps_mode = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `adaptive_noise_immunity` - (Optional) Type: `string`. This property is only effective for cards based on Atheros chipset.
* `allow_sharedkey` - (Optional) Type: `string`. Allow WEP Shared Key clients to connect. Note that no authentication is done for these clients (WEP Shared keys are not compared to anything) - they are just accepted at once (if access list allows that).
* `ampdu_priorities` - (Optional) Type: `string`. Frame priorities for which AMPDU sending (aggregating frames and sending using block acknowledgment) should get negotiated and used. Using AMPDUs will increase throughput, but may increase latency, therefore, may not be desirable for real-time traffic (voice, video). Due to this, by default AMPDUs are enabled only for best-effort traffic.
* `amsdu_limit` - (Optional) Type: `string`. Max AMSDU that device is allowed to prepare when negotiated. AMSDU aggregation may significantly increase throughput especially for small frames, but may increase latency in case of packet loss due to retransmission of aggregated frame. Sending and receiving AMSDUs will also increase CPU usage.
* `amsdu_threshold` - (Optional) Type: `string`. Max frame size to allow including in AMSDU.
* `antenna_gain` - (Optional) Type: `string`. Antenna gain in dBi, used to calculate maximum transmit power according to country regulations.
* `antenna_mode` - (Optional) Type: `string`. Select antenna to use for transmitting and for receiving ant-a - use only 'a' antenna ant-b - use only 'b' antenna txa-rxb - use antenna 'a' for transmitting, antenna 'b' for receiving rxa-txb - use antenna 'b' for transmitting, antenna 'a' for receiving.
* `area` - (Optional) Type: `string`. Identifies group of wireless networks. This value is announced by AP, and can be matched in connect-list by area-prefix . This is a proprietary extension.
* `arp` - (Optional) Type: `string`. Read more >>.
* `arp_timeout` - (Optional) Type: `string`. ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in /ip settings , default is 30s.
* `band` - (Optional) Type: `string`. Defines set of used data rates, channel frequencies and widths.
* `basic_rates_b` - (Optional) Type: `string`. List of basic rates, used for 2.4ghz-b, 2.4ghz-b/g and 2.4ghz-onlyg bands. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP. This property has effect only in AP modes, and when value of rate-set is configured.
* `bridge_mode` - (Optional) Type: `string`. Allows to use station-bridge mode. Read more >>.
* `burst_time` - (Optional) Type: `string`. Time in microseconds which will be used to send data without stopping. Note that no other wireless cards in that network will be able to transmit data during burst-time microseconds. This setting is available only for AR5000, AR5001X, and AR5001X+ chipset based cards.
* `channel_width` - (Optional) Type: `string`. Use of extension channels (e.g. Ce, eC etc) allows additional 20MHz extension channels and if it should be located below or above the control (main) channel. Extension channel allows 802.11n devices to use up to 40MHz (802.11ac up to 160MHz) of spectrum in total thus increasing max throughput. Channel widths with XX and XXXX extensions automatically scan for a less crowded control channel frequency based on the number of concurrent devices running in every frequency and chooses the “C” - Control channel frequency automatically.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `compression` - (Optional) Type: `string`. Setting this property to yes will allow the use of the hardware compression. Wireless interface must have support for hardware compression. Connections with devices that do not use compression will still work.
* `country` - (Optional) Type: `string`. Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default value of scan-list . Value no_country_set is an FCC compliant set of channels.
* `default_ap_tx_limit` - (Optional) Type: `string`. This is the value of ap-tx-limit for clients that do not match any entry in the access-list . 0 means no limit.
* `default_authentication` - (Optional) Type: `string`. For AP mode, this is the value of authentication for clients that do not match any entry in the access-list . For station mode, this is the value of connect for APs that do not match any entry in the connect-list.
* `default_client_tx_limit` - (Optional) Type: `string`. This is the value of client-tx-limit for clients that do not match any entry in the access-list . 0 means no limit.
* `default_forwarding` - (Optional) Type: `string`. This is the value of forwarding for clients that do not match any entry in the access-list.
* `disable_running_check` - (Optional) Type: `string`. When set to yes interface will always have running flag. If value is set to no' , the router determines whether the card is up and running - for AP one or more clients have to be registered to it, for station, it should be connected to an AP.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `disconnect_timeout` - (Optional) Type: `string`. This interval is measured from third sending failure on the lowest data rate. At this point 3 * ( hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully during disconnect-timeout , the connection is closed, and this event is logged as "extensive data loss". Successful frame transmission resets this timer.
* `distance` - (Optional) Type: `string`. How long to wait for confirmation of unicast frames ( ACKs ) before considering transmission unsuccessful, or in short ACK-Timeout . Distance value has these behaviors: Dynamic - causes AP to detect and use the smallest timeout that works with all connected clients. Indoor - uses the default ACK timeout value that the hardware chip manufacturer has set. Number - uses the input value in formula: ACK-timeout = (( distance * 1000) + 299) / 300 us; Acknowledgments are not used in Nstreme/NV2 protocols.
* `frame_lifetime` - (Optional) Type: `string`. Discard frames that have been queued for sending longer than frame-lifetime . By default, when value of this property is 0 , frames are discarded only after connection is closed.
* `frequency` - (Optional) Type: `string`. Channel frequency value in MHz on which AP will operate. Allowed values depend on the selected band, and are restricted by country setting and wireless card capabilities. This setting has no effect if interface is in any of station modes, or in wds-slave mode, or if DFS is active. Note : If using mode "superchannel", any frequency supported by the card will be accepted, but on the RouterOS client, any non-standard frequency must be configured in the scan-list , otherwise it will not be scanning in non-standard range. In Winbox, scanlist frequencies are in bold , any other frequency means the clients will need scan-list configured.
* `frequency_mode` - (Optional) Type: `string`. Three frequency modes are available: regulatory-domain - Limit available channels and maximum transmit power for each channel according to the value of country manual-txpower - Same as above, but do not limit maximum transmit power. superchannel - Conformance Testing Mode. Allow all channels supported by the card. List of available channels for each band can be seen in /interface wireless info allowed-channels . This mode allows you to test wireless channels outside the default scan-list and/or regulatory domain. This mode should only be used in controlled environments, or if you have special permission to use it in your region. Before v4.3 this was called Custom Frequency Upgrade, or Superchannel. Since RouterOS v4.3 this mode is available without special key upgrades to all installations.
* `frequency_offset` - (Optional) Type: `string`. Allows to specify offset if the used wireless card operates at a different frequency than is shown in RouterOS, in case a frequency converter is used in the card. So if your card works at 4000MHz but RouterOS shows 5000MHz, set offset to 1000MHz and it will be displayed correctly. The value is in MHz and can be positive or negative.
* `guard_interval` - (Optional) Type: `string`. Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how this may affect throughput). "any" will use either short or long, depending on data rate, "long" will use long.
* `hide_ssid` - (Optional) Type: `string`. yes - AP does not include SSID in the beacon frames, and does not reply to probe requests that have broadcast SSID. no - AP includes SSID in the beacon frames, and replies to probe requests that have broadcast SSID. This property has an effect only in AP mode. Setting it to yes can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP.
* `ht_basic_mcs` - (Optional) Type: `string`. Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.
* `ht_supported_mcs` - (Optional) Type: `string`. Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.
* `hw_fragmentation_threshold` - (Optional) Type: `string`. Specifies maximum fragment size in bytes when transmitted over the wireless medium. 802.11 standard packet (MSDU in 802.11 terminologies) fragmentation allows packets to be fragmented before transmitting over a wireless medium to increase the probability of successful transmission (only fragments that did not transmit correctly are retransmitted). Note that transmission of a fragmented packet is less efficient than transmitting unfragmented packet because of protocol overhead and increased resource usage at both - transmitting and receiving party.
* `hw_protection_mode` - (Optional) Type: `string`. Frame protection support property read more >>.
* `hw_protection_threshold` - (Optional) Type: `string`. Frame protection support property read more >>.
* `hw_retries` - (Optional) Type: `string`. Number of times sending frame is retried without considering it a transmission failure. Data-rate is decreased upon failure and the frame is sent again. Three sequential failures on the lowest supported rate suspend transmission to this destination for the duration of on-fail-retry-time . After that, the frame is sent again. The frame is being retransmitted until transmission success, or until the client is disconnected after disconnect-timeout . The frame can be discarded during this time if frame-lifetime is exceeded.
* `installation` - (Optional) Type: `string`. Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.
* `interworking_profile` - (Optional) Type: `string`.
* `keepalive_frames` - (Optional) Type: `string`. Applies only if wireless interface is in mode= ap-bridge . If a client has not communicated for around 20 seconds, AP sends a "keepalive-frame". Note , disabling the feature can lead to "ghost" clients in registration-table.
* `l2mtu` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `master_interface` - (Optional) Type: `string`. Name of wireless interface that has virtual-ap capability. Virtual AP interface will only work if master interface is in ap-bridge , bridge , station or wds-slave mode. This property is only for virtual AP interfaces.
* `max_station_count` - (Optional) Type: `string`. Maximum number of associated clients. WDS links also count toward this limit.
* `mode` - (Optional) Type: `string`. Selection between different station and access point (AP) modes. : station - Basic station mode. Find and connect to acceptable AP. station-wds - Same as station , but create WDS link with AP, using proprietary extension. AP configuration has to allow WDS links with this device. Note that this mode does not use entries in wds. station-pseudobridge - Same as station , but additionally perform MAC address translation of all traffic. Allows interface to be bridged. station-pseudobridge-clone - Same as station-pseudobridge , but use station-bridge-clone-mac address to connect to AP. station-bridge - Provides support for transparent protocol-independent L2 bridging on the station device. RouterOS AP accepts clients in station-bridge mode when enabled using bridge-mode parameter. In this mode, the AP maintains a forwarding table with information on which MAC addresses are reachable over which station device. Only works with RouterOS APs. With station-bridge mode, it is not possible to connect to CAPsMAN controlled CAP. The 'wireless' station-bridge mode,  is incompatible with APs running the newer 'wifi' package and vice versa. AP modes: ap-bridge - Basic access point mode. bridge - Same as ap-bridge , but limited to one associated client. wds-slave - Same as ap-bridge , but scan for AP with the same ssid and establishes WDS link. If this link is lost or cannot be established, then continue scanning. If dfs-mode is radar-detect , then APs with enabled hide-ssid will not be found during scanning. Special modes: alignment-only - Put the interface in a continuous transmit mode that is used for aiming the remote antenna. nstreme-dual-slave - allow this interface to be used in nstreme-dual setup. MAC address translation in pseudobridge modes works by inspecting packets and building a table of corresponding IP and MAC addresses. All packets are sent to AP with the MAC address used by pseudobridge, and MAC addresses of received packets are restored from the address translation table. There is a single entry in the address translation table for all non-IP packets, hence more than one host in the bridged network cannot reliably use non-IP protocols. Note: Currently IPv6 doesn't work over Pseudobridge.
* `mtu` - (Optional) Type: `string`.
* `multicast_buffering` - (Optional) Type: `string`. For a client that has power saving, buffer multicast packets until next beacon time. A client should wake up to receive a beacon, by receiving beacon it sees that there are multicast packets pending, and it should wait for multicast packets to be sent.
* `multicast_helper` - (Optional) Type: `string`. When set to full , multicast packets will be sent with a unicast destination MAC address, resolving multicast problem on the wireless link. This option should be enabled only on the access point, clients should be configured in station-bridge mode. Available starting from v5.15. disabled - disables the helper and sends multicast packets with multicast destination MAC addresses dhcp - dhcp packet mac addresses are changed to unicast mac addresses prior to sending them out full - all multicast packet mac address are changed to unicast mac addresses prior to sending them out default - default choice that currently is set to dhcp . Value can be changed in future releases.
* `name` - (Optional) Type: `string`. name of the interface.
* `noise_floor_threshold` - (Optional) Type: `string`. For advanced use only, as it can badly affect the performance of the interface. It is possible to manually set noise floor threshold value. By default, it is dynamically calculated. This property also affects received signal strength. This property is only effective on non-AC chips.
* `nv2_cell_radius` - (Optional) Type: `string`. Setting affects the size of contention time slot that AP allocates for clients to initiate connection and also size of time slots used for estimating distance to client. When setting is too small, clients that are farther away may have trouble connecting and/or disconnect with "ranging timeout" error. Although during normal operation the effect of this setting should be negligible, in order to maintain maximum performance, it is advised to not increase this setting if not necessary, so AP is not reserving time that is actually never used, but instead allocates it for actual data transfer. on AP: distance to farthest client in km on station: no effect.
* `nv2_noise_floor_offset` - (Optional) Type: `string`.
* `nv2_preshared_key` - (Optional) Type: `string`. **Sensitive.**
* `nv2_qos` - (Optional) Type: `string`. Sets the packet priority mechanism, firstly data from high priority queue is sent, then lower queue priority data until 0 queue priority is reached. When link is full with high priority queue data, lower priority data is not sent. Use it very carefully, setting works on AP frame-priority - manual setting that can be tuned with Mangle rules. default - default setting where small packets receive priority for best latency.
* `nv2_queue_count` - (Optional) Type: `string`.
* `nv2_security` - (Optional) Type: `string`.
* `on_fail_retry_time` - (Optional) Type: `string`. After third sending failure on the lowest data rate, wait for specified time interval before retrying.
* `periodic_calibration` - (Optional) Type: `string`. Setting default enables periodic calibration if info default-periodic-calibration property is enabled . Value of that property depends on the type of wireless card. This property is only effective for cards based on Atheros chipset.
* `periodic_calibration_interval` - (Optional) Type: `string`. This property is only effective for cards based on Atheros chipset.
* `preamble_mode` - (Optional) Type: `string`. Short preamble mode is an option of 802.11b standard that reduces per-frame overhead. On AP: long - Do not use short preamble. short - Announce short preamble capability. Do not accept connections from clients that do not have this capability. both - Announce short preamble capability. On station: long - do not use short preamble. short - do not connect to AP if it does not support short preamble. both - Use short preamble if AP supports it.
* `prism_cardtype` - (Optional) Type: `string`. Specify type of the installed Prism wireless card.
* `proprietary_extensions` - (Optional) Type: `string`. RouterOS includes proprietary information in an information element of management frames. This parameter controls how this information is included. pre-2.9.25 - This is older method. It can interoperate with newer versions of RouterOS. This method is incompatible with some clients, for example, Centrino based ones. post-2.9.25 - This uses standardized way of including vendor specific information, that is compatible with newer wireless clients.
* `radio_name` - (Optional) Type: `string`. Descriptive name of the device, that is shown in registration table entries on the remote devices. This is a proprietary extension.
* `rate_selection` - (Optional) Type: `string`. Starting from v5.9 default value is advanced since legacy mode was inefficient.
* `rate_set` - (Optional) Type: `string`. Two options are available: default - default basic and supported rate sets are used. Values from basic-rates and supported-rates parameters have no effect. configured - use values from basic-rates , supported-rates , basic-mcs , mcs . Read more >> .
* `rx_chains` - (Optional) Type: `string`. Which antennas to use for receive. In current MikroTik routers, both RX and TX chain must be enabled, for the chain to be enabled.
* `scan_list` - (Optional) Type: `string`. The default value is all channels from selected band that are supported by card and allowed by the country and frequency-mode settings (this list can be seen in info ). For default scan list in 5ghz band channels are taken with 20MHz step, in 5ghz-turbo band - with 40MHz step, for all other bands - with 5MHz step. If scan-list is specified manually, then all matching channels are taken. (Example: scan-list = default,5200-5245,2412-2427 - This will use the default value of scan list for current band, and add to it supported frequencies from 5200-5245 or 2412-2427 range.) Since RouterOS v6.0 with Winbox or Webfig, for inputting of multiple frequencies, add each frequency or range of frequencies into separate multiple scan-lists. Using a comma to separate frequencies is no longer supported in Winbox/Webfig since v6.0. Since RouterOS v6.35 (wireless-rep) scan-list support step feature where it is possible to manually specify the scan step. Example: scan-list = 5500-5600:20 will generate such scan-list values 5500,5520,5540,5560,5580,5600.
* `secondary_channel` - (Optional) Type: `string`. Specifies secondary channel, required to enable 80+80MHz transmission. To disable 80+80MHz functionality, set secondary-channel to "" or unset the value via CLI/GUI.
* `security_profile` - (Optional) Type: `string`. Name of profile from security-profiles.
* `skip_dfs_channels` - (Optional) Type: `string`. These values are used to skip all DFS channels or specifically skip DFS CAC channels in range 5600-5650MHz which detection could go up to 10min.
* `ssid` - (Optional) Type: `string`. SSID (service set identifier) is a name that identifies wireless network.
* `station_bridge_clone_mac` - (Optional) Type: `string`. This property has effect only in the station-pseudobridge-clone mode. Use this MAC address when connection to AP. If this value is 00:00:00:00:00:00 , station will initially use MAC address of the wireless interface. As soon as packet with MAC address of another device needs to be transmitted, station will reconnect to AP using that address.
* `station_roaming` - (Optional) Type: `string`. Station Roaming feature is available only for 802.11 wireless protocol and only for station modes. Read more >>.
* `supported_rates_b` - (Optional) Type: `string`. List of supported rates, used for 2ghz-b , 2ghz-b/g and 2ghz-b/g/n bands. Two devices will communicate only using rates that are supported by both devices. This property has effect only when value of rate-set is configured .
* `tdma_period_size` - (Optional) Type: `string`. Specifies TDMA period in milliseconds. It could help on the longer distance links, it could slightly increase bandwidth, while latency is increased too.
* `tx_chains` - (Optional) Type: `string`. Which antennas to use for transmitting. In current MikroTik routers, both RX and TX chain must be enabled, for the chain to be enabled.
* `tx_power` - (Optional) Type: `string`. For 802.11ac wireless interface it's total power but for 802.11a/b/g/n it's power per chain.
* `tx_power_mode` - (Optional) Type: `string`. sets up tx-power mode for wireless card default - use values stored in the card all-rates-fixed - use same transmit power for all data rates. Can damage the card if transmit power is set above rated value of the card for used rate. manual-table - define transmit power for each rate separately. Can damage the card if transmit power is set above rated value of the card for used rate. card-rates - use transmit power calculated for each rate based on value of tx-power parameter. Legacy mode only compatible with currently discontinued products.
* `update_stats_interval` - (Optional) Type: `string`. How often to request update of signals strength and ccq values from clients. Access to registration-table also triggers update of these values. This is proprietary extension.
* `vht_basic_mcs` - (Optional) Type: `string`. Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9.
* `vht_supported_mcs` - (Optional) Type: `string`. Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9.
* `wds_cost_range` - (Optional) Type: `string`. Bridge port cost of WDS links are automatically adjusted, depending on measured link throughput. Port cost is recalculated and adjusted every 5 seconds if it has changed by more than 10%, or if more than 20 seconds have passed since the last adjustment. Setting this property to 0 disables automatic cost adjustment. Automatic adjustment does not work for WDS links that are manually configured as a bridge port.
* `wds_default_bridge` - (Optional) Type: `string`. When WDS link is established and status of the wds interface becomes running , it will be added as a bridge port to the bridge interface specified by this property. When WDS link is lost, wds interface is removed from the bridge. If wds interface is already included in a bridge setup when WDS link becomes active, it will not be added to bridge specified by , and will (needs editing).
* `wds_default_cost` - (Optional) Type: `string`. Initial bridge port cost of the WDS links.
* `wds_ignore_ssid` - (Optional) Type: `string`. By default, WDS link between two APs can be created only when they work on the same frequency and have the same SSID value. If this property is set to yes , then SSID of the remote AP will not be checked. This property has no effect on connections from clients in station-wds mode. It also does not work if wds-mode is static-mesh or dynamic-mesh .
* `wds_mode` - (Optional) Type: `string`. Controls how WDS links with other devices (APs and clients in station-wds mode) are established. disabled does not allow WDS links. static only allows WDS links that are manually configured in WDS dynamic also allows WDS links with devices that are not configured in WDS, by creating required entries dynamically. Such dynamic WDS entries are removed automatically after the connection with the other AP is lost. -mesh modes use different (better) method for establishing link between AP, that is not compatible with APs in non-mesh mode. This method avoids one-sided WDS links that are created only by one of the two APs. Such links cannot pass any data.When AP or station is establishing WDS connection with another AP, it uses connect-list to check whether this connection is allowed. If station in station-wds mode is establishing connection with AP, AP uses access-list to check whether this connection is allowed.If mode is station-wds, then this property has no effect.
* `wireless_protocol` - (Optional) Type: `string`. Specifies protocol used on wireless interface; unspecified - protocol mode used on previous RouterOS versions (v3.x, v4.x). Nstreme is enabled by old enable-nstreme setting, Nv2 configuration is not possible. any  : on AP - regular 802.11 Access Point or Nstreme Access Point; on station - selects Access Point without specific sequence, it could be changed by connect-list rules. nstreme - enables Nstreme protocol (the same as old enable-nstreme setting). nv2 - enables Nv2 protocol. nv2 nstreme  : on AP - uses first wireless-protocol setting, always Nv2; on station - searches for Nv2 Access Point, then for Nstreme Access Point. nv2 nstreme 802.11 - on AP - uses first wireless-protocol setting, always Nv2; on station - searches for Nv2 Access Point, then for Nstreme Access Point, then for regular 802.11 Access Point. Warning! Nv2 doesn't have support for Virtual AP.
* `wmm_support` - (Optional) Type: `string`. Specifies whether to enable WMM .  Only applies to bands B and G. Other bands will have it enabled regardless of this setting.
* `wps_mode` - (Optional) Type: `string`. Read more >>.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wireless.example '*3'

# Named router
terraform import routeros_interface_wireless.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wireless.example 'home/my-resource-name'
```
