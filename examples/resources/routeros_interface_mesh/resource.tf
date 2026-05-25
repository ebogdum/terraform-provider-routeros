resource "routeros_interface_mesh" "mesh_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # admin_mac_address = "10.99.0.0/24"
  # arp = "enabled"
  # arp_timeout = "1h"
  # default_hoplimit = 32
  # mesh_portal = false
  # mesh_traceroute = "replace-me"
  # mtu = 1500
  # prep_lifetime = "300"
  # preq_destination_only = true
  # preq_reply_and_forward = true
  # preq_retries = 2
  # preq_waiting_time = 4
  # rann_interval = "10"
  # rann_lifetime = "22"
  # rann_propagation_delay = 500
  # reoptimize_paths = false
}
