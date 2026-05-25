---
page_title: "RouterOS: routeros_interface_ethernet_switch"
description: |-
  Switch chip menu varies by hardware; not on hAP/CHR
---

# Resource: routeros_interface_ethernet_switch

Switch chip menu varies by hardware; not on hAP/CHR

## Example Usage

```terraform
resource "routeros_interface_ethernet_switch" "switch_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # autorestart = "replace-me"
  # fasttrack_hw = "replace-me"
  # icmp_reply_on_error = "replace-me"
  # ipv6_hw = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `autorestart` - (Optional) Type: `string`. Automatically restarts the l3hw driver in case of an error. Otherwise, if an error occurs,   l3-hw-offloading   gets disabled, and the error code is displayed in the switch settings and   #monitor . Autorestart does not work for system failures, such as OOM (Out Of Memory).
* `fasttrack_hw` - (Optional) Type: `string`. Enables or disables FastTrack HW Offloading. Keep it enabled unless HW TCAM memory reservation is required, e.g., for dynamic switch ACL rules creation. Not all switch chips support FastTrack HW Offloading (see   hw-supports-fasttrack ).
* `icmp_reply_on_error` - (Optional) Type: `string`. Since the hardware cannot send ICMP messages, the packet must be redirected to the CPU to send an ICMP reply in case of an error (e.g., "Time Exceeded", "Fragmentation required", etc.). Enabling icmp-reply-on-error   helps with network diagnostics but may open potential vulnerabilities for DDoS attacks. Disabling icmp-reply-on-error silently drops the packets on the hardware level in case of an error.
* `ipv6_hw` - (Optional) Type: `string`. Enables or disables IPv6 Hardware Offloading. Since IPv6 routes occupy a lot of HW memory, enable it only if IPv6 traffic speed is significant enough to benefit from hardware routing.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ethernet_switch.example '*3'

# Named router
terraform import routeros_interface_ethernet_switch.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ethernet_switch.example 'home/my-resource-name'
```
