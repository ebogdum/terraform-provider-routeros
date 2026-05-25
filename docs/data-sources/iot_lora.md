---
subcategory: "System & misc"
page_title: "RouterOS: routeros_iot_lora"
description: |-
  Requires LoRa hardware/package
---

# Data Source: routeros_iot_lora

Requires LoRa hardware/package

## Example Usage

```terraform
data "routeros_iot_lora" "lora_example" {
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
* `antenna_gain` - (Optional) Type: `string`. Antenna gain in dBi. This value should be equal to  setup-antenna-gain minus cable-loss . Using 6.5 dBi antenna, 6.5 is the value to be configured (not taking into account cable loss).  Output power of the gateway is dictated by the server. The gateway will calculate its actual output power by subtracting antenna-gain setting from server_value (value received in the downlink message).
* `channel_plan` - (Optional) Type: `string`. Frequency plans for various regions.
* `disabled` - (Optional) Type: `string`. Whether LoRaWAN gateway is disabled.
* `forward` - (Optional) Type: `string`. Defines what kind of packets should be forwarded to Network server: crc-validtaion - Forward valid packets with correct CRC. dev-addr-validtaion - Checks if DevAddr of the packet corresponds to the NetID and if not, drops the packet. The following sequence happens: 1) Dev. Addr value gets "obtained" from the received LoRa packet; 2) Dev. Addr is "compared" against "valid" Net IDs list; 3) If there is no Net ID for the Dev. Addr, the packet is not forwarded; 4) If Net ID is valid, Dev. Addr range is valid, the packet is forwarded. proprietary-traffic - Checks the content of the LoRa packet and if the "type" of the frame is "proprietary", the packet is not forwarded.
* `gateway_id` - (Optional) Type: `string`. Gateway ID or Gateway EUI, is used when registering the gateway with the server.
* `lbt_enabled` - (Optional) Type: `string`. Whether gateway should use LBT (Listen Before Talk) protocol.
* `listen_time` - (Optional) Type: `string`. Time in microseconds to track RSSI before TX (used when lbt-enabled=yes ).
* `name` - (Optional) Type: `string`. Name of LoRaWAN gateway.
* `network` - (Optional) Type: `string`. Whether sync word should (network=private) or should not (network=public) be used.
* `rssi_threshold` - (Optional) Type: `string`. RSSI value to determine whether forwarder may use specific channel to talk. If RSSI value is below rssi-threshold , channel could be used (used when lbt-enabled=yes ).
* `servers` - (Optional) Type: `string`. Name of the server from the /iot lora servers section.
* `spoof_gps` - (Optional) Type: `string`. Set custom GPS location: Latitude [-90..90] Longitude [-180..180] Altitude( m ) [-2147483648..2147483647].
* `src_address` - (Optional) Type: `string`. Specifies uplink packet source address if necessary (address should match an address configured on the RB).

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

