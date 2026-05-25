resource "routeros_routing_ospf_area" "area_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # area_id = "10.99.0.1"
  # default_cost = "replace-me"
  # instance = "replace-me"
  # name = "tf-example"
  # no_summaries = false
  # nssa_translator = "replace-me"
  # type = "default"
}
