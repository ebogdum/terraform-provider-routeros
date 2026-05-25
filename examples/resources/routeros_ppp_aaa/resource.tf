resource "routeros_ppp_aaa" "aaa_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # enable_ipv6_accounting = false
  # interim_update = "1h"
  # use_circuit_id_in_nas_port_id = false
  # use_radius = false
}
