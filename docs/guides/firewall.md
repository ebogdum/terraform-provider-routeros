---
subcategory: "Guides"
page_title: "RouterOS: firewall example"
description: |-
  End-to-end working example using the routeros provider.
---

# firewall example

Example: a working stateful firewall ruleset.

Covers:
  - input chain: accept established/related, drop invalid, accept from a
    trusted address list, drop everything else from the WAN interface
  - forward chain: accept established/related, drop invalid
  - srcnat: masquerade outbound traffic
  - dstnat: forward TCP/443 from the WAN interface to a LAN host
  - a growable trusted-sources address list
  - position-based deterministic ordering (chain stays in this order across
    destroy/recreate)
  - lockout_ack = true on every drop rule, required by the provider's
    lockout safety guard

Apply:
  terraform apply \
    -var 'routeros_host=https://192.0.2.1' \
    -var 'routeros_user=admin' \
    -var 'routeros_password=...'


## Configuration

```terraform
# Example: a working stateful firewall ruleset.
#
# Covers:
#   - input chain: accept established/related, drop invalid, accept from a
#     trusted address list, drop everything else from the WAN interface
#   - forward chain: accept established/related, drop invalid
#   - srcnat: masquerade outbound traffic
#   - dstnat: forward TCP/443 from the WAN interface to a LAN host
#   - a growable trusted-sources address list
#   - position-based deterministic ordering (chain stays in this order across
#     destroy/recreate)
#   - lockout_ack = true on every drop rule, required by the provider's
#     lockout safety guard
#
# Apply:
#   terraform apply \
#     -var 'routeros_host=https://192.0.2.1' \
#     -var 'routeros_user=admin' \
#     -var 'routeros_password=...'

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

variable "wan_interface" {
  type        = string
  default     = "ether1"
  description = "Physical interface facing the internet."
}

provider "routeros" {
  host     = var.routeros_host
  username = var.routeros_user
  password = var.routeros_password
  insecure = true
}

# ---------------- Address list ----------------

resource "routeros_ip_firewall_address_list" "trusted_lan" {
  list    = "tf-example-trusted"
  address = "10.99.0.0/24"
  comment = "managed by terraform"
}

resource "routeros_ip_firewall_address_list" "trusted_mgmt_host" {
  list    = "tf-example-trusted"
  address = "10.99.1.10"
  comment = "managed by terraform"
}

# ---------------- INPUT chain ----------------

resource "routeros_ip_firewall_filter" "input_established" {
  chain            = "input"
  action           = "accept"
  connection_state = "established,related"
  position         = 9000
  comment          = "managed by terraform -- accept established/related"
  lockout_ack      = true
}

resource "routeros_ip_firewall_filter" "input_invalid" {
  chain            = "input"
  action           = "drop"
  connection_state = "invalid"
  position         = 9010
  comment          = "managed by terraform -- drop invalid"
  lockout_ack      = true
}

resource "routeros_ip_firewall_filter" "input_trusted" {
  chain            = "input"
  action           = "accept"
  src_address_list = routeros_ip_firewall_address_list.trusted_lan.list
  position         = 9020
  comment          = "managed by terraform -- allow trusted LAN"
  lockout_ack      = true
}

resource "routeros_ip_firewall_filter" "input_drop_wan" {
  chain        = "input"
  action       = "drop"
  in_interface = var.wan_interface
  position     = 9030
  comment      = "managed by terraform -- drop unsolicited from WAN"
  lockout_ack  = true
}

# ---------------- FORWARD chain ----------------

resource "routeros_ip_firewall_filter" "fwd_established" {
  chain            = "forward"
  action           = "accept"
  connection_state = "established,related"
  position         = 9100
  comment          = "managed by terraform -- established/related"
  lockout_ack      = true
}

resource "routeros_ip_firewall_filter" "fwd_drop_invalid" {
  chain            = "forward"
  action           = "drop"
  connection_state = "invalid"
  position         = 9110
  comment          = "managed by terraform -- drop invalid forward"
  lockout_ack      = true
}

# ---------------- NAT ----------------

resource "routeros_ip_firewall_nat" "masquerade" {
  chain         = "srcnat"
  action        = "masquerade"
  out_interface = var.wan_interface
  comment       = "managed by terraform -- masquerade LAN"
}

resource "routeros_ip_firewall_nat" "port_forward_https" {
  chain        = "dstnat"
  action       = "dst-nat"
  protocol     = "tcp"
  dst_port     = "443"
  in_interface = var.wan_interface
  to_addresses = "10.99.0.10"
  to_ports     = "443"
  comment      = "managed by terraform -- public 443 -> 10.99.0.10"
}
```

## Apply

```sh
terraform apply -auto-approve \
  -var 'routeros_host=https://192.0.2.1' \
  -var 'routeros_user=admin' \
  -var 'routeros_password=...'
```

Each apply is **idempotent**: re-running returns *No changes*. `terraform destroy` removes everything the guide creates.
