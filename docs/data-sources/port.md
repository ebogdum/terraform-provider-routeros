---
subcategory: "Port"
page_title: "RouterOS: routeros_port"
description: |-
  RouterOS resource.
---

# Data Source: routeros_port

Manages the RouterOS `/port` menu.

## Example Usage

```terraform
data "routeros_port" "port_example" {
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
* `baud_rate` - (Optional) Type: `string`. Baud rate (speed) used by the port. If set to auto , then RouterOS tries to detect baud rate automatically.
* `data_bits` - (Optional) Type: `string`. The number of data bits in each character. 7 - true ASCII 8 - any data (matches the size of a byte).
* `dtr` - (Optional) Type: `string`. Whether to enable RS-232 DTR signal circuit used by flow control.
* `flow_control` - (Optional) Type: `string`. method of flow control to pause and resume the transmission of data.
* `name` - (Optional) Type: `string`. Name of the port.
* `parity` - (Optional) Type: `string`. Error detection method. If enabled, extra bit is sent to detect the communication errors. In most cases parity is set to none and errors are handled by the communication protocol.
* `rts` - (Optional) Type: `string`. Whether to enable RS-232 RTS signal circuit used by flow control.
* `stop_bits` - (Optional) Type: `string`. Stop bits sent after each character. Electronic devices usually uses 1 stop bit.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

