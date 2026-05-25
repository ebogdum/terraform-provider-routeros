resource "routeros_partition" "partition_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # activate = "replace-me"
  # active = false
  # fallback_to = "replace-me"
  # name = "tf-example"
  # running = false
}
