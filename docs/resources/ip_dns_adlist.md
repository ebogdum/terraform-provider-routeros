---
subcategory: "DNS"
page_title: "RouterOS: routeros_ip_dns_adlist"
description: |-
  A DNS adlist (blocklist) entry in RouterOS /ip/dns/adlist.
---

# Resource: routeros_ip_dns_adlist

A DNS adlist (blocklist) entry in RouterOS `/ip/dns/adlist`. The router loads a
hosts-format or adblock-format list and answers matching queries with
`0.0.0.0`, which is the built-in ad/malware blocking introduced in **RouterOS
7.15**. Earlier versions do not have this menu and will return a 404 on apply.

Supply exactly one of:

* `url` — the router downloads and periodically refreshes the list, or
* `file` — a list already stored on the router's filesystem.

The provider rejects a configuration that sets both or neither at plan time.

DNS caching must be enabled for adlists to take effect (see `routeros_ip_dns`).
Large lists consume router RAM; check `name_count` after applying.

## Example Usage

```terraform
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
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the adlist entry is disabled.
* `file` - (Optional) Type: `string`. Name of a blocklist file already present on the router. Mutually exclusive with `url`.
* `match_count` - (Read-only) Type: `int`. Read-only: number of queries blocked by this list.
* `name_count` - (Read-only) Type: `int`. Read-only: number of names loaded from this list.
* `ssl_verify` - (Optional) Type: `bool`. Verify the TLS certificate when downloading `url`.
* `url` - (Optional) Type: `string`. HTTP(S) URL of a hosts-format or adblock-format blocklist. Mutually exclusive with `file`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, or by `url`, `file` or `comment`. Because a
URL contains `/`, use the `::` separator to name a router explicitly:

```sh
# Default router, .id = *3
terraform import routeros_ip_dns_adlist.example '*3'

# Named router, by .id
terraform import routeros_ip_dns_adlist.example 'home::*3'

# Named router, by url
terraform import routeros_ip_dns_adlist.example 'home::https://adaway.org/hosts.txt'
```
