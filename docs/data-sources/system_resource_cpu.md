---
page_title: "RouterOS: routeros_system_resource_cpu"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_resource_cpu

Manages the RouterOS `/system/resource/cpu` menu.

## Example Usage

```terraform
data "routeros_system_resource_cpu" "cpu_example" {
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
* `architecture_name` - (Optional) Type: `string`. CPU architecture.
* `bad_blocks` - (Optional) Type: `string`. Shows percentage of bad blocks on the NAND.
* `board_name` - (Optional) Type: `string`. RouterBOARD model name.
* `build_time` - (Optional) Type: `string`. Installed RouterOS version build-time.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `cpu` - (Optional) Type: `string`. CPU model that is on the board.
* `cpu_count` - (Optional) Type: `string`. Number of CPUs present on the system. Each core is a separate CPU, Intel HT is also a separate CPU.
* `cpu_frequency` - (Optional) Type: `string`. Current CPU frequency.
* `cpu_load` - (Optional) Type: `string`. Percentage of used CPU resources. Combines all CPUs. Per-core CPU usage can be seen in  CPU submenu.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `disk` - (Optional) Type: `int`.
* `factory_software` - (Optional) Type: `string`. Minimal RouterOS version.
* `free_hdd_space` - (Optional) Type: `string`. Free space on hard drive or NAND.
* `free_memory` - (Optional) Type: `string`. The unused amount of RAM.
* `irq` - (Optional) Type: `int`.
* `load` - (Optional) Type: `int`.
* `platform` - (Optional) Type: `string`. Platform name.
* `profile` - (Optional) Type: `string`.
* `total_hdd_space` - (Optional) Type: `string`. Size of the hard drive or NAND.
* `total_memory` - (Optional) Type: `string`. Amount of installed RAM.
* `uptime` - (Optional) Type: `string`. Time interval passed since boot-up.
* `version` - (Optional) Type: `string`. Installed RouterOS version number.
* `write_sect_since_reboot` - (Optional) Type: `string`. A number of sector writes in HDD or NAND since the router was last time rebooted.
* `write_sect_total` - (Optional) Type: `string`. A number of sector writes in total.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

