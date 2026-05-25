---
subcategory: "System"
page_title: "RouterOS: routeros_system_leds"
description: |-
  LED bindings — type/leds values depend on the specific device's available LEDs; not portable in an auto-test.
---

# Data Source: routeros_system_leds

LED bindings — type/leds values depend on the specific device's available LEDs; not portable in an auto-test.

## Example Usage

```terraform
data "routeros_system_leds" "leds_example" {
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
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`. Name of the interface which will be used for led status. Applicable only if   type   is interface specific.
* `leds` - (Optional) Type: `string`. List of led names used for a status report. For example, wireless signal strength will require more than one led.
* `type` - (Optional) Type: `string`. Type of the status: align-down   - light the led if the w60g device needs to be aligned downwards for the best signal quality align-left   - light the led if the w60g device needs to be aligned to the left align-right   - light the led if the w60g device needs to be aligned to the right align-up   - light the led if the w60g device needs to be aligned upwards ap-cap   - blink on CAP initializing with CAPsMAN, steady on once connected fan-fault   - light the led when any of the devices controlled fans stop working flash-access   - blink the led on flash access interface-activity   - blink the led on interface (traffic) activity interface-receive   - blink the led on interface received a traffic interface-speed   - light the led when interface works in 10Gbit rate interface-speed-1G   - light the led when interface works in 1Gbit rate interface-speed-25G   - light the led when interface works in 25Gbit rate interface-speed-100G - light the led when interface works in 100Gbit rate interface-status   - light the led on interface status change interface-transmit   - blink the led on interface transmitted traffic modem-signal   - blink the led on 3G modem signal (either USB or miniPCIe) modem-technology   - turns on LEDs in order of modem technology generation: GSM; 3G; LTE; single led turns on only when LTE is active. off   - turn off the led on   - turn on the led poe-fault   - light the led when PoE out budget is close to the maximum supported limit poe-out   - light the led when interface PoE out turns on wireless-signal-strength   - light the leds displaying wireless signal (requires more than one led) wireless-status   - light the led on wireless status change.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

