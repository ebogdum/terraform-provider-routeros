---
subcategory: "SNMP"
page_title: "RouterOS: routeros_snmp"
description: |-
  Singleton; live-edit risks orphaning agent — skip in acc
---

# Resource: routeros_snmp

Singleton; live-edit risks orphaning agent — skip in acc

## Example Usage

```terraform
resource "routeros_snmp" "snmp_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # contact = "replace-me"
  # enabled = false
  # engine_id_suffix = "replace-me"
  # location = "replace-me"
  # src_address = "10.99.0.0/24"
  # trap_community = "replace-me"
  # trap_generators = "replace-me"
  # trap_interfaces = "replace-me"
  # trap_target = "replace-me"
  # trap_version = 0
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `contact` - (Optional) Type: `string`. Contact information.
* `enabled` - (Optional) Type: `bool`. Used to disable/enable SNMP service.
* `engine_id_suffix` - (Optional) Type: `string`.
* `location` - (Optional) Type: `string`. Location information.
* `src_address` - (Optional) Type: `ip`. Force the router to always use the same IP source address for all of the SNMP messages.
* `trap_community` - (Optional) Type: `string`. Which communities configured in the community menu to use when sending out the trap.
* `trap_generators` - (Optional) Type: `string`. What action will generate traps: interfaces - interface changes; start-trap - SNMP server starting on the router temp-exception - send trap when temperature reached 100c (or value configured for cpu-overtemp-temperature at /system health ).
* `trap_interfaces` - (Optional) Type: `string`. List of interfaces that traps are going to be sent out.
* `trap_target` - (Optional) Type: `string`. IP (IPv4 or IPv6) addresses of SNMP data collectors that have to receive the trap.
* `trap_version` - (Optional) Type: `int`. A version of SNMP protocol to use for trap.
* `vrf` - (Optional) Type: `string`. Set VRF on which service is listening for incoming connections.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `engine_id` - Type: `string`. For SNMP v3, used as part of the identifier. You can configure the suffix part of the engine id using this argument. If the SNMP client is not capable to detect set engine-id value then this prefix hex has to be used 0x80003a8c04.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_snmp.this 'home'
```
