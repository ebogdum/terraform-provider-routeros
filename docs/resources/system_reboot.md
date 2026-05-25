---
page_title: "RouterOS: routeros_system_reboot"
description: |-
  Reboots the router. Verified working against a CHR VM (test PASSes with
---

# Resource: routeros_system_reboot

Reboots the router. Verified working against a CHR VM (test PASSes with
ROUTEROS_RUN_DESTRUCTIVE_ACTIONS=1, isolated invocation). Skipped from the
general acc suite because it takes the device out for ~30s, breaking every
subsequent test in the same run.

To verify locally on a disposable device:
  go test -tags acceptance -run '^TestAccSystemReboot$' -v ./internal/provider/...


## Example Usage

```terraform
resource "routeros_system_reboot" "reboot_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

