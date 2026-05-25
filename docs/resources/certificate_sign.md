---
page_title: "RouterOS: routeros_certificate_sign"
description: |-
  Sign a certificate. Pass `certificate` (the cert name to sign), `ca` (signer
---

# Resource: routeros_certificate_sign

Sign a certificate. Pass `certificate` (the cert name to sign), `ca` (signer
cert name; omit for self-sign), and `name` (optional new name for the
signed cert). Triggered by changing `trigger`.


## Example Usage

```terraform
resource "routeros_certificate_sign" "sign_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

