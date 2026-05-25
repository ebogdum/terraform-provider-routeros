resource "routeros_interface_mesh" "mesh_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "enabled"
  # arp_timeout = "1h"
  # mesh_portal = false
  # mtu = 1500
  # reoptimize_paths = false
}
