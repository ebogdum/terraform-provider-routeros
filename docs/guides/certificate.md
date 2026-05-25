---
subcategory: "Guides"
page_title: "RouterOS: certificate example"
description: |-
  End-to-end working example using the routeros provider.
---

# certificate example

Example: stand up a self-signed Certificate Authority on the router and
sign a leaf certificate (suitable for HTTPS/WebFig, OVPN server, IPsec, etc.).

The /certificate/sign action is per-row -- the cert to sign is selected
via the action's `target_id` attribute, which references the cert
resource's `.id`. Change the action's `trigger` to "v2" etc. to re-sign.


## Configuration

```terraform
# Example: stand up a self-signed Certificate Authority on the router and
# sign a leaf certificate (suitable for HTTPS/WebFig, OVPN server, IPsec, etc.).
#
# The /certificate/sign action is per-row -- the cert to sign is selected
# via the action's `target_id` attribute, which references the cert
# resource's `.id`. Change the action's `trigger` to "v2" etc. to re-sign.

terraform {
  required_version = ">= 1.4"
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = ">= 1.0"
    }
  }
}

variable "routeros_host" { type = string }
variable "routeros_user" { type = string }
variable "routeros_password" {
  type      = string
  sensitive = true
}

provider "routeros" {
  host     = var.routeros_host
  username = var.routeros_user
  password = var.routeros_password
  insecure = true
}

# ---------------- CA template ----------------

resource "routeros_certificate" "ca" {
  name        = "tf-example-ca"
  common_name = "tf-example-ca"
  days_valid  = "3650"
  key_size    = "2048"
  key_usage   = "key-cert-sign,crl-sign"
}

# ---------------- Leaf template ----------------

resource "routeros_certificate" "leaf" {
  name        = "tf-example-leaf"
  common_name = "router.example.local"
  days_valid  = "365"
  key_size    = "2048"
  key_usage   = "digital-signature,key-encipherment,tls-server"
}

# ---------------- Sign the CA (self-sign) ----------------

resource "routeros_certificate_sign" "ca" {
  trigger   = "v1"
  target_id = routeros_certificate.ca.id
}

# ---------------- Sign the leaf with the CA ----------------

resource "routeros_certificate_sign" "leaf" {
  trigger    = "v1"
  target_id  = routeros_certificate.leaf.id
  params     = { ca = routeros_certificate.ca.name }
  depends_on = [routeros_certificate_sign.ca]
}

output "ca_name"   { value = routeros_certificate.ca.name }
output "leaf_name" { value = routeros_certificate.leaf.name }
```

## Apply

```sh
terraform apply -auto-approve \
  -var 'routeros_host=https://192.0.2.1' \
  -var 'routeros_user=admin' \
  -var 'routeros_password=...'
```

Each apply is **idempotent**: re-running returns *No changes*. `terraform destroy` removes everything the guide creates.
