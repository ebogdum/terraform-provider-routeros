---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_netwatch"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_netwatch

Manages the RouterOS `/tool/netwatch` menu.

## Example Usage

```terraform
resource "routeros_tool_netwatch" "netwatch_example" {
  # router = "my-router"  # which router to target; omit for the default
  host = "1.1.1.1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # certificate = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # interval = "1m"
  # name = "tf-example"
  # port = "443"
  # src_address = "10.99.0.0/24"
  # timeout = "replace-me"
  # ttl = "replace-me"
  # type = "icmp"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accept_icmp_time_exceeded` - (Optional) Type: `string`. RouterOS `accept-icmp-time-exceeded`.
* `certificate` - (Optional) Type: `string`.
* `check_certificate` - (Optional) Type: `string`. RouterOS `check-certificate`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dns_server` - (Optional) Type: `string`.
* `down_script` - (Optional) Type: `string`. RouterOS `down-script`.
* `early_failure_detection` - (Optional) Type: `string`. RouterOS `early-failure-detection`.
* `early_success_detection` - (Optional) Type: `string`. RouterOS `early-success-detection`.
* `host` - (Required) Type: `string`.
* `http_codes` - (Optional) Type: `string`. RouterOS `http-codes`.
* `ignore_initial_down` - (Optional) Type: `string`. RouterOS `ignore-initial-down`.
* `ignore_initial_up` - (Optional) Type: `string`. RouterOS `ignore-initial-up`.
* `interval` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `packet_count` - (Optional) Type: `string`. RouterOS `packet-count`.
* `packet_interval` - (Optional) Type: `string`. RouterOS `packet-interval`.
* `packet_size` - (Optional) Type: `string`. RouterOS `packet-size`.
* `port` - (Optional) Type: `string`.
* `record_type` - (Optional) Type: `string`. RouterOS `record-type`.
* `src_address` - (Optional) Type: `string`.
* `start_delay` - (Optional) Type: `string`. RouterOS `start-delay`.
* `startup_delay` - (Optional) Type: `string`. RouterOS `startup-delay`.
* `test_script` - (Optional) Type: `string`. RouterOS `test-script`.
* `thr_avg` - (Optional) Type: `string`. RouterOS `thr-avg`.
* `thr_http_time` - (Optional) Type: `string`. RouterOS `thr-http-time`.
* `thr_jitter` - (Optional) Type: `string`. RouterOS `thr-jitter`.
* `thr_loss_count` - (Optional) Type: `string`. RouterOS `thr-loss-count`.
* `thr_loss_percent` - (Optional) Type: `string`. RouterOS `thr-loss-percent`.
* `thr_max` - (Optional) Type: `string`. RouterOS `thr-max`.
* `thr_stdev` - (Optional) Type: `string`. RouterOS `thr-stdev`.
* `thr_tcp_conn_time` - (Optional) Type: `string`. RouterOS `thr-tcp-conn-time`.
* `timeout` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`.
* `up_script` - (Optional) Type: `string`. RouterOS `up-script`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_netwatch.example '*3'

# Named router
terraform import routeros_tool_netwatch.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_netwatch.example 'home/my-resource-name'
```
