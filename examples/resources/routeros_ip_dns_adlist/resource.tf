# DNS adlists require RouterOS 7.15 or newer, and DNS caching must be enabled
# (see routeros_ip_dns).

resource "routeros_ip_dns_adlist" "stevenblack" {
  # router = "my-router"  # which router to target; omit for the default
  url        = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
  ssl_verify = true
  comment    = "unified hosts: adware + malware"
}

# A list already downloaded onto the router's filesystem.
resource "routeros_ip_dns_adlist" "local" {
  file    = "blocklist.txt"
  comment = "locally curated"
}

# Staged but not active yet.
resource "routeros_ip_dns_adlist" "staged" {
  url      = "https://adaway.org/hosts.txt"
  disabled = true
  comment  = "evaluate before enabling"
}
