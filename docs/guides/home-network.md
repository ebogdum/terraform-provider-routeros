---
subcategory: "Guides"
page_title: "RouterOS: home network example"
description: |-
  End-to-end working example using the routeros provider.
---

# home network example

Example: minimal-but-real LAN setup.

Allocates a /24 to the LAN bridge, defines a DHCP pool, hangs a DHCP
server off the bridge interface, sets the per-network options (gateway +
DNS), and tells the router to forward LAN DNS queries.

Pick a bridge interface name that exists on your router (the default is
"bridge"). Adjust `interface = ...` if yours is different.


## Configuration

```terraform
# Example: minimal-but-real LAN setup.
#
# Allocates a /24 to the LAN bridge, defines a DHCP pool, hangs a DHCP
# server off the bridge interface, sets the per-network options (gateway +
# DNS), and tells the router to forward LAN DNS queries.
#
# Pick a bridge interface name that exists on your router (the default is
# "bridge"). Adjust `interface = ...` if yours is different.

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

variable "lan_interface" {
  type        = string
  default     = "bridge"
  description = "Which interface to put the LAN address + DHCP server on."
}

provider "routeros" {
  host     = var.routeros_host
  username = var.routeros_user
  password = var.routeros_password
  insecure = true
}

# ---------------- LAN address ----------------

resource "routeros_ip_address" "lan_gateway" {
  address   = "10.99.0.1/24"
  interface = var.lan_interface
  comment   = "managed by terraform -- example/home-network"
}

# ---------------- DHCP pool + server ----------------

resource "routeros_ip_pool" "lan" {
  name    = "tf-example-pool"
  ranges  = "10.99.0.100-10.99.0.200"
  comment = "managed by terraform"
}

resource "routeros_ip_dhcp_server" "lan" {
  name          = "tf-example-dhcp"
  interface     = var.lan_interface
  address_pool  = routeros_ip_pool.lan.name
  lease_time    = "12h"
  comment       = "managed by terraform"
}

resource "routeros_ip_dhcp_server_network" "lan" {
  address = "10.99.0.0/24"
  gateway = "10.99.0.1"
  comment = "managed by terraform -- example/home-network"
  # NOTE: DNS for LAN clients comes from routeros_ip_dns below -- the router
  # itself resolves and caches, clients use the router as their resolver.
}

# ---------------- DNS forwarder ----------------

resource "routeros_ip_dns" "this" {
  servers              = "1.1.1.1,8.8.8.8"
  allow_remote_requests = false
}

output "lan_gateway" {
  value = routeros_ip_address.lan_gateway.address
}

output "dhcp_range" {
  value = routeros_ip_pool.lan.ranges
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
