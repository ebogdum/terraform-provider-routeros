resource "routeros_mpls_ldp_accept_filter" "accept_filter_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # prefix = "replace-me"
  # vrf = "main"
}
