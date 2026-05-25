resource "routeros_ip_ipsec_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  # Optional attributes (uncomment as needed):
  # dh_group = ["16388"]
  # dpd_interval = "disable DPD"
  # dpd_maximum_failures = 4
  # enc_algorithm = []
  # hash_algorithm = "replace-me"
  # lifebytes = 0
  # lifetime = "86400"
  # nat_traversal = true
  # ppk = "No"
  # proposal_check = "obey"
}
