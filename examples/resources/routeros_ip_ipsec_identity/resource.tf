resource "routeros_ip_ipsec_identity" "identity_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auth_method = "rsa key"
  # generate_policy = "no"
  # match_by = "certificate"
  # notrack_chain = "replace-me"
  # peer = "replace-me"
  # policy_template_group = "replace-me"
}
