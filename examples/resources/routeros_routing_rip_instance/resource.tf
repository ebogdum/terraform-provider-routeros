resource "routeros_routing_rip_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # afi = "replace-me"
  # name = "tf-example"
  # originate_default = "replace-me"
  # redistribute = "replace-me"
  # route_gc_timeout = "replace-me"
  # route_timeout = "replace-me"
  # routing_table = "main"
  # update_interval = "replace-me"
  # vrf = "main"
}
