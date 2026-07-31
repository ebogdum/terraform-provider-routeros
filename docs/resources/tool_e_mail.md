---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_e_mail"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_e_mail

Manages the RouterOS `/tool/e-mail` menu.

## Example Usage

```terraform
resource "routeros_tool_e_mail" "e_mail_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # certificate_verification = false
  # from = "replace-me"
  # password = "REDACTED"
  # port = "443"
  # server = "replace-me"
  # tls = false
  # user = "myuser"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `certificate_verification` - (Optional) Type: `bool`.
* `from` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `port` - (Optional) Type: `int`.
* `server` - (Optional) Type: `string`.
* `tls` - (Optional) Type: `bool`.
* `user` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_e_mail.this 'home'
```
