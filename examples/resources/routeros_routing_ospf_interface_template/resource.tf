resource "routeros_routing_ospf_interface_template" "interface_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # area = "replace-me"
  # auth_id = "replace-me"
  # auth_key = "REDACTED"
  # cost = 0
  # dead_interval = "1h"
  # hello_interval = "1h"
  # instance_id = 0
  # interfaces = "replace-me"
  # networks = "replace-me"
  # passive = "replace-me"
  # prefix_list = "replace-me"
  # priority = 0
  # retransmit_interval = "1h"
  # transmit_delay = 0
  # use_bfd = "replace-me"
  # vlink_neighbor_id = "replace-me"
  # vlink_transit_area = "replace-me"
}
