---
subcategory: "Guides"
page_title: "RouterOS: system example"
description: |-
  End-to-end working example using the routeros provider.
---

# system example

Example: a baseline "system" config every router benefits from -- identity,
NTP, a script, a scheduler that runs the script every hour, and a log/info
action that leaves an audit trail entry on every apply.


## Configuration

```terraform
# Example: a baseline "system" config every router benefits from -- identity,
# NTP, a script, a scheduler that runs the script every hour, and a log/info
# action that leaves an audit trail entry on every apply.

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

variable "router_identity" {
  type    = string
  default = "tf-example-router"
}

provider "routeros" {
  host     = var.routeros_host
  username = var.routeros_user
  password = var.routeros_password
  insecure = true
}

# ---------------- Identity ----------------

resource "routeros_system_identity" "this" {
  name = var.router_identity
}

# ---------------- NTP ----------------

resource "routeros_system_ntp_client" "this" {
  enabled = true
  servers = "pool.ntp.org"
  mode    = "unicast"
}

# ---------------- Script ----------------

resource "routeros_system_script" "log_uptime" {
  name    = "tf-example-log-uptime"
  owner   = var.routeros_user
  source  = ":log info (\"uptime: \" . [/system resource get uptime])"
  policy  = ["read", "write", "policy", "test"]
  comment = "managed by terraform"
}

# ---------------- Scheduler ----------------

resource "routeros_system_scheduler" "hourly_uptime" {
  name     = "tf-example-hourly-uptime"
  on_event = routeros_system_script.log_uptime.name
  interval = "1h"
  comment  = "managed by terraform"
}

# ---------------- Audit log on every apply ----------------

resource "routeros_log_info" "deploy_marker" {
  # Change `trigger` to re-fire the action on the next apply.
  # Use formatdate("YYYY-MM-DD-hh-mm", timestamp()) if you want it to
  # re-fire on every plan/apply, but be aware that introduces churn.
  trigger = "v1"
  message = "terraform apply against ${var.router_identity}"
}

output "identity"    { value = routeros_system_identity.this.name }
output "ntp_servers" { value = routeros_system_ntp_client.this.servers }
```

## Apply

```sh
terraform apply -auto-approve \
  -var 'routeros_host=https://192.0.2.1' \
  -var 'routeros_user=admin' \
  -var 'routeros_password=...'
```

Each apply is **idempotent**: re-running returns *No changes*. `terraform destroy` removes everything the guide creates.
