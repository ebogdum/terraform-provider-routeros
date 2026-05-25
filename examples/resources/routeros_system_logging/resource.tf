resource "routeros_system_logging" "logging_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # name = "tf-example"
  # prefix = "replace-me"
  # regex = "replace-me"
  # topics = "replace-me"
}
