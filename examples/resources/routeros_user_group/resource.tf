resource "routeros_user_group" "group_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # policy = ["read"]
  # skin = "replace-me"
}
