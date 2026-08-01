---
subcategory: "System"
page_title: "RouterOS: routeros_system_watchdog"
description: |-
  Singleton; misconfig can soft-brick — skip in acc
---

# Resource: routeros_system_watchdog

Singleton; misconfig can soft-brick — skip in acc

## Example Usage

```terraform
resource "routeros_system_watchdog" "watchdog_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # auto_send_supout = false
  # automatic_supout = false
  # no_ping_delay = "replace-me"
  # ping_start_after_boot = "1h"
  # ping_timeout = "1h"
  # send_email_from = "replace-me"
  # send_email_to = "replace-me"
  # send_smtp_server = "replace-me"
  # watch_address = "10.99.0.0/24"
  # watchdog_timer = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `auto_send_supout` - (Optional) Type: `bool`. After the support output file is automatically generated, it can be sent by email.
* `automatic_supout` - (Optional) Type: `bool`. When software failure happens, a file named "autosupout.rif" is generated automatically. The previous "autosupout.rif" file is renamed to "autosupout.old.rif".
* `lockout_ack` - (Optional) Type: `bool`. Acknowledge that a `watch_address` the device cannot reach will reboot-loop it off the network.
* `no_ping_delay` - (Optional) Type: `string`. Specifies how long will it wait before trying to reach the watch-address.
* `ping_start_after_boot` - (Optional) Type: `string`.
* `ping_timeout` - (Optional) Type: `string`. Specifies the time interval in which the device will be pinged 6 times (after "no-ping-delay").
* `send_email_from` - (Optional) Type: `string`. The e-mail address to send the support output file from. If not set, the value set in /tool e-mail is used.
* `send_email_to` - (Optional) Type: `string`. The e-mail address to send the support output file to.
* `send_smtp_server` - (Optional) Type: `string`. SMTP server address to send the support output file through. If not set, the value set in /tool e-mail is used.
* `watch_address` - (Optional) Type: `string`. The system will reboot, in case 6 sequential pings to the given IP address will fail. If set to none this feature is disabled. By default, the router will reboot every 6 minutes if the watch-address is set and not reachable.
* `watchdog_timer` - (Optional) Type: `bool`. Whether to reboot if a system is unresponsive for a minute.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_watchdog.this 'home'
```
