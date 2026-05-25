---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_policy"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_ip_ipsec_policy

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_ip_ipsec_policy" "policy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "encrypt"
  # active = false
  # dst_address = "10.99.0.0/24"
  # dst_port = "443"
  # group = "replace-me"
  # ipsec_protocols = "esp"
  # level = "unique"
  # nopeer = "replace-me"
  # notemplate = "replace-me"
  # peer = "replace-me"
  # proposal = "replace-me"
  # protocol = "icmp"
  # src_address = "10.99.0.0/24"
  # src_port = "443"
  # template = false
  # tunnel = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Optional) Type: `enum(discard|none|encrypt)`. Default: `2`.
* `active` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `dst_address` - (Optional) Type: `cidr`.
* `dst_port` - (Optional) Type: `int`.
* `group` - (Optional) Type: `string`.
* `ipsec_protocols` - (Optional) Type: `enum(|ah|esp)`. Default: `2`.
* `level` - (Optional) Type: `enum(use|require|unique)`. Default: `2`.
* `nopeer` - (Optional) Type: `string`.
* `notemplate` - (Optional) Type: `string`.
* `peer` - (Optional) Type: `string`.
* `proposal` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `enum(icmp|igmp|ggp|ip-encap|tcp|egp, ...)`. Default: `255`.
* `src_address` - (Optional) Type: `cidr`.
* `src_port` - (Optional) Type: `int`.
* `template` - (Optional) Type: `bool`.
* `tunnel` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `ph2_count` - Type: `int`.
* `ph2_state` - Type: `enum(spawning|starting|message-1-received|message-1-sent|message-2-received|message-2-sent, ...)`.
* `sa_dst_address` - Type: `string`.
* `sa_src_address` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_policy.example '*3'

# Named router
terraform import routeros_ip_ipsec_policy.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_policy.example 'home/my-resource-name'
```
