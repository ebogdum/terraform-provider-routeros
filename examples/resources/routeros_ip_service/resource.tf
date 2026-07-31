# /ip/service rows are a fixed set. This resource adopts the row named by
# `name` and manages its settings; it never adds or removes rows.

# Turn off the legacy plaintext services.
resource "routeros_ip_service" "telnet" {
  # router = "my-router"  # which router to target; omit for the default
  name     = "telnet"
  disabled = true
}

resource "routeros_ip_service" "ftp" {
  name     = "ftp"
  disabled = true
}

resource "routeros_ip_service" "www" {
  name     = "www"
  disabled = true
}

# Plaintext API off, TLS API on and restricted to the management subnet.
resource "routeros_ip_service" "api" {
  name     = "api"
  disabled = true
}

resource "routeros_ip_service" "api_ssl" {
  name        = "api-ssl"
  disabled    = false
  port        = 8729
  address     = "10.0.0.0/24"
  certificate = "api-cert"
  tls_version = "only-v1.2"
}

# Keep SSH on a non-default port, reachable only from the management subnet.
resource "routeros_ip_service" "ssh" {
  name     = "ssh"
  disabled = false
  port     = 2200
  address  = "10.0.0.0/24"
}
