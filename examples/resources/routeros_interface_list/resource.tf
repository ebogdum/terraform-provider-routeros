resource "routeros_interface_list" "list_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # exclude = "replace-me"
  # include = "replace-me"
}
