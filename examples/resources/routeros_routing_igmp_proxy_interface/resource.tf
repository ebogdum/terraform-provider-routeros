resource "routeros_routing_igmp_proxy_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alternative_subnets = "replace-me"
  # interface = "ether1"
  # threshold = 1
  # upstream = false
}
