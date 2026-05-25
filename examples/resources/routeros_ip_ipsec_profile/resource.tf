resource "routeros_ip_ipsec_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  # Optional attributes (uncomment as needed):
  # dh_group = ["16388"]
  # dpd_interval = "disable-dpd"
  # dpd_maximum_failures = 4
  # enc_algorithm = []
  # encryption_algorithm = "12"
  # hash_algorithm = "replace-me"
  # hash_algorithms = "sha256"
  # lifebytes = 0
  # lifetime = "86400"
  # nat_traversal = true
  # ppk = "no"
  # prf_algorithms = "auto"
  # proposal_check = "obey"
}
