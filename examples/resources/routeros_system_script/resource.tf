resource "routeros_system_script" "script_example" {
  # router = "my-router"  # which router to target; omit for the default
  name   = "tf-example"
  source = ":put \"hello\""

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # dont_require_permissions = false
  # policy = []
  # run_script = "replace-me"
}
