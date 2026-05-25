resource "routeros_mpls_mangle" "mangle_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # chain = ""
  # exp = "0"
  # set_exp = "0"
  # set_mark = "replace-me"
}
