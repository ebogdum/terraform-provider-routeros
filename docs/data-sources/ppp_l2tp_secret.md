---
subcategory: "PPP"
page_title: "RouterOS: routeros_ppp_l2tp_secret"
description: |-
  L2TP CHAP/PAP shared secret entry. Schema varies across ROS releases. Skipped from automated acc tests.
---

# Data Source: routeros_ppp_l2tp_secret

L2TP CHAP/PAP shared secret entry. Schema varies across ROS releases. Skipped from automated acc tests.

## Example Usage

```terraform
data "routeros_ppp_l2tp_secret" "l2tp_secret_example" {
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
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `secret` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

