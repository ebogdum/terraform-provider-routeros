---
subcategory: "Guides"
page_title: "RouterOS: wireguard vpn example"
description: |-
  End-to-end working example using the routeros provider.
---

# wireguard vpn example

Example: WireGuard server interface plus a peer.

The router generates its own private key on first apply. The matching
public key is read back through the resource and emitted as an output so
you can hand it to the peer.

The peer's public key is supplied as an input variable. Generate it on
the peer side with `wg genkey | wg pubkey` and pass it in.


## Configuration

```terraform
# Example: WireGuard server interface plus a peer.
#
# The router generates its own private key on first apply. The matching
# public key is read back through the resource and emitted as an output so
# you can hand it to the peer.
#
# The peer's public key is supplied as an input variable. Generate it on
# the peer side with `wg genkey | wg pubkey` and pass it in.

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

variable "peer_public_key" {
  type        = string
  description = "Base64 WireGuard public key for the peer."
}

variable "peer_preshared_key" {
  type      = string
  sensitive = true
  default   = ""
}

provider "routeros" {
  host     = var.routeros_host
  username = var.routeros_user
  password = var.routeros_password
  insecure = true
}

# ---------------- WireGuard server interface ----------------

resource "routeros_interface_wireguard" "wg0" {
  name        = "tf-example-wg0"
  listen_port = "51820"
  comment     = "managed by terraform -- example/wireguard-vpn"
}

# ---------------- IP address on the WireGuard interface ----------------

resource "routeros_ip_address" "wg0" {
  address   = "10.99.99.1/24"
  interface = routeros_interface_wireguard.wg0.name
  comment   = "managed by terraform -- example/wireguard-vpn"
}

# ---------------- Peer ----------------

resource "routeros_interface_wireguard_peers" "laptop" {
  name                 = "tf-example-laptop"
  interface            = routeros_interface_wireguard.wg0.name
  public_key           = var.peer_public_key
  preshared_key        = var.peer_preshared_key
  allowed_address      = "10.99.99.2/32"
  persistent_keepalive = "25s"
  comment              = "managed by terraform -- laptop peer"
}

output "wireguard_listen_port" {
  value = routeros_interface_wireguard.wg0.listen_port
}

output "wireguard_server_address" {
  value = routeros_ip_address.wg0.address
}

output "wireguard_server_public_key" {
  value       = routeros_interface_wireguard.wg0.public_key
  description = "Distribute this to peers as their PublicKey."
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
