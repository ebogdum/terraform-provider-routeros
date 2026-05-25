resource "routeros_ip_ipsec_proposal" "proposal_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auth_algorithms = "md5"
  # enc_algorithms = []
  # lifetime = "1800"
  # name = "example"
  # pfs_group = "replace-me"
}
