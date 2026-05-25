resource "routeros_routing_bgp_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # as = "replace-me"
  # cluster_id = "replace-me"
  # ignore_as_path = "replace-me"
  # name = "tf-example"
  # router_id = "replace-me"
  # routing_table = "main"
  # vrf = "main"
}
