---
page_title: "RouterOS: routeros_ip_ipsec_identity"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Data Source: routeros_ip_ipsec_identity

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
data "routeros_ip_ipsec_identity" "identity_example" {
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
* `auth_method` - (Optional) Type: `enum(pre shared key|rsa key|digital signature|pre shared key xauth|rsa signature hybrid|eap radius, ...)`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `generate_policy` - (Optional) Type: `enum(no|port override|port strict)`.
* `match_by` - (Optional) Type: `enum(certificate|remote id)`. Default: `100`.
* `notrack_chain` - (Optional) Type: `string`.
* `peer` - (Optional) Type: `string`.
* `policy_template_group` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

