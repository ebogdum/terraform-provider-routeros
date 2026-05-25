resource "routeros_routing_id" "id_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # select_dynamic_id = "only static"
  # select_from_vrf = "replace-me"
}
