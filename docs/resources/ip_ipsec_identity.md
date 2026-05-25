---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_identity"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_ip_ipsec_identity

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_ip_ipsec_identity" "identity_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auth_method = "rsa key"
  # generate_policy = "no"
  # match_by = "certificate"
  # notrack_chain = "replace-me"
  # peer = "replace-me"
  # policy_template_group = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auth_method` - (Optional) Type: `enum(pre shared key|rsa key|digital signature|pre shared key xauth|rsa signature hybrid|eap radius, ...)`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `generate_policy` - (Optional) Type: `enum(no|port override|port strict)`.
* `match_by` - (Optional) Type: `enum(certificate|remote id)`. Default: `100`.
* `notrack_chain` - (Optional) Type: `string`.
* `peer` - (Optional) Type: `string`.
* `policy_template_group` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_identity.example '*3'

# Named router
terraform import routeros_ip_ipsec_identity.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_identity.example 'home/my-resource-name'
```
