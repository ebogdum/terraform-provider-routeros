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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`.
* `active` - (Read-only) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dst_address` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `int`.
* `dynamic` - (Read-only) Type: `bool`.
* `group` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `ipsec_protocols` - (Optional) Type: `string`.
* `level` - (Optional) Type: `string`.
* `nopeer` - (Read-only) Type: `string`.
* `notemplate` - (Optional) Type: `string`.
* `peer` - (Optional) Type: `string`.
* `ph2_count` - (Read-only) Type: `int`.
* `ph2_state` - (Read-only) Type: `string`.
* `proposal` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `sa_dst_address` - (Read-only) Type: `string`.
* `sa_src_address` - (Read-only) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `int`.
* `template` - (Optional) Type: `bool`.
* `tunnel` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
