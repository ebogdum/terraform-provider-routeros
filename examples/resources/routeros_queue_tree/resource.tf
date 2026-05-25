resource "routeros_queue_tree" "tree_example" {
  # router = "my-router"  # which router to target; omit for the default
  name   = "tf-example"
  parent = "global"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # burst_limit = "replace-me"
  # burst_threshold = "replace-me"
  # burst_time = "replace-me"
  # limit_at = "replace-me"
  # max_limit = "replace-me"
  # packet_mark = "replace-me"
  # place_before = "replace-me"
  # priority = "replace-me"
  # queue = "replace-me"
}
