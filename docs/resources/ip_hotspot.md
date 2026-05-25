---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_hotspot"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot

Manages the RouterOS `/ip/hotspot` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot" "hotspot_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "tf-example"

  disabled = false

  # Optional attributes (uncomment as needed):
  # dns_name = "replace-me"
  # hotspot_address = "10.99.0.0/24"
  # html_directory = "replace-me"
  # html_directory_override = "replace-me"
  # http_cookie_lifetime = "replace-me"
  # http_proxy = "replace-me"
  # https_redirect = "replace-me"
  # keepalive_timeout = "replace-me"
  # login_by = "replace-me"
  # mac_auth_password = "REDACTED"
  # nas_port_type = "replace-me"
  # profile = "replace-me"
  # radius_accounting = "replace-me"
  # radius_default_domain = "replace-me"
  # radius_interim_update = "replace-me"
  # radius_location_name = "replace-me"
  # radius_mac_format = "replace-me"
  # rate_limit = "replace-me"
  # smtp_server = "replace-me"
  # split_user_domain = "replace-me"
  # ssl_certificate = "replace-me"
  # trial_uptime = "replace-me"
  # trial_user_profile = "replace-me"
  # use_radius = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dns_name` - (Optional) Type: `string`. DNS name of the HotSpot server. This is the DNS name used as the name of the HotSpot server (i.e., it appears as the location of the login page). This name will automatically be added as a static DNS entry in the DNS cache.
* `hotspot_address` - (Optional) Type: `string`. IP address of HotSpot service.
* `html_directory` - (Optional) Type: `string`. Directory name in which HotSpot HTML pages are stored (by default hotspot directory). It is possible to specify different directory with modified HTML pages. To change HotSpot login page, get HotSpot files from your router, change and upload them back to same location. F ull path must be typed in html-directory field, including "/flash/(hotspot_dir)".
* `html_directory_override` - (Optional) Type: `string`. Alternative path for hotspot html files. It should be used only when customized hotspot html files are stored on external storage.
* `http_cookie_lifetime` - (Optional) Type: `string`. HTTP cookie validity time, the option is related to cookie HotSpot login method.
* `http_proxy` - (Optional) Type: `string`. Address and port of the proxy server for HotSpot service, when default value is used all request are resolved by the local /ip proxy.
* `https_redirect` - (Optional) Type: `string`. Whether to redirect unauthenticated user to hotspot login page, if user is visiting a https:// url. Since certificate domain name will mismatch, often this leads to errors, so you can set this parameter to "no" and all https requests will simply be rejected and user will have to visit a http page.
* `interface` - (Required) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `login_by` - (Optional) Type: `string`. Used HotSpot authentication method mac-cookie - enables login by mac cookie method cookie - may only be used with other HTTP authentication method. HTTP cookie is generated, when user authenticates in HotSpot for the first time. User is not asked for the login/password and authenticated automatically, until cookie-lifetime is active http-chap - login/password is required for the user to authenticate in HotSpot. CHAP challenge-response method with MD5 hashing algorithm is used for protecting passwords. http-pap - login/password is required for user to authenticate in HotSpot. Username and password are sent over network in plain text. https - login/password is required for user to authenticate in HotSpot. Client login/password exchange between client and server is encrypted with SSL tunnel.  mac - client is authenticated without asking login form. Client MAC-address is added to /ip hotspot user database, client is authenticated as soon as connected to the HotSpot trial - client is allowed to use internet without HotSpot login for the specified amount of time.
* `mac_auth_password` - (Optional) Type: `string`. Used together with MAC authentication, field used to specify password for the users to be authenticated by their MAC addresses. The following option is required, when specific RADIUS server rejects authentication for the clients with blank password.
* `name` - (Required) Type: `string`. Descriptive name of the profile. Default: `tf_acc_hotspot`.
* `nas_port_type` - (Optional) Type: `string`. NAS-Port-Type value to be sent to RADIUS server, NAS-Port-Type values are described in the RADIUS RFC 2865. This optional value attribute indicates the type of the physical port of the HotSpot server.
* `profile` - (Optional) Type: `string`.
* `radius_accounting` - (Optional) Type: `string`. Send RADIUS server accounting information for each user, when yes is used.
* `radius_default_domain` - (Optional) Type: `string`. Default domain to use for RADIUS requests. Allows to use separate RADIUS server per /ip hotspot profile . If used, same domain name should be specified under /radius domain value.
* `radius_interim_update` - (Optional) Type: `string`. How often to send accounting updates . When received is set, interim-time is used from RADIUS server. 0s is the same as received .
* `radius_location_name` - (Optional) Type: `string`. RADIUS-Location-Id to be sent to RADIUS server. Used to identify location of the HotSpot server during the communication with RADIUS server. Value is optional and used together with RADIUS server.
* `radius_mac_format` - (Optional) Type: `string`. Option to set format of user mac-address, that is sent to RADIUS server during AAA session.
* `rate_limit` - (Optional) Type: `string`. Rate limitation in form of rx-rate[/tx-rate] [rx-burst-rate[/tx-burst-rate] [rx-burst-threshold[/tx-burst-threshold] [rx-burst-time[/tx-burst-time]]]] [priority] [rx-rate-min[/tx-rate-min]] from the point of view of the router (so "rx" is client upload, and "tx" is client download). All rates should be numbers with optional 'k' (1,000s) or 'M' (1,000,000s). If tx-rate is not specified, rx-rate is as tx-rate too. Same goes for tx-burst-rate and tx-burst-threshold and tx-burst-time. If both rx-burst-threshold and tx-burst-threshold are not specified (but burst-rate is specified), rx-rate and tx-rate is used as burst thresholds. If both rx-burst-time and tx-burst-time are not specified, 1s is used as default. rx-rate-min and tx-rate min are the values of limit-at properties.
* `smtp_server` - (Optional) Type: `string`. SMTP server address to be used to redirect HotSpot users SMTP requests.
* `split_user_domain` - (Optional) Type: `string`. Split username from domain name when the username is given in "user@domain" or in "domain\user" format from RADIUS server.
* `ssl_certificate` - (Optional) Type: `string`. Name of the SSL certificate on the router to to use only for HTTPS authentication.
* `trial_uptime` - (Optional) Type: `string`. Used only with trial authentication method. First time value specifies, how long trial user identified by MAC address can use access to public networks without HotSpot authentication. Second time value specifies amount of time, that has to pass until user is allowed to use trial again.
* `trial_user_profile` - (Optional) Type: `string`. Specifies hotspot user profile for trial users.
* `use_radius` - (Optional) Type: `string`. Use RADIUS to authenticate HotSpot users.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot.example '*3'

# Named router
terraform import routeros_ip_hotspot.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot.example 'home/my-resource-name'
```
