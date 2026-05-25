resource "routeros_routing_filter_num_list" "num_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
