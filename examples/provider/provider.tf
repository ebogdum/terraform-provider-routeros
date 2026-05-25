terraform {
  required_providers {
    routeros = {
      source  = "ebogdum/routeros"
      version = "~> 1.0"
    }
  }
}

# Manage one router (single-router back-compat).
provider "routeros" {
  alias    = "single"
  host     = "https://192.0.2.1"
  username = "admin"
  password = var.routeros_password
  insecure = true
}

# Manage a whole fleet of routers from one config.
provider "routeros" {
  routers = {
    core = {
      host     = "https://10.0.0.1"
      username = "admin"
      password = var.core_password
      insecure = true
    }
    edge_se = {
      host     = "https://10.0.1.1"
      username = "admin"
      password = var.edge_se_password
      insecure = true
    }
    edge_nw = {
      host     = "https://10.0.2.1"
      username = "admin"
      password = var.edge_nw_password
      insecure = true
    }
  }
}

# Per-resource router selection.
resource "routeros_ip_address" "core_lan" {
  router    = "core"
  address   = "192.168.1.1/24"
  interface = "bridge1"
}
