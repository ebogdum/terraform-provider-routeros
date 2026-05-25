resource "routeros_ip_ipsec_policy" "policy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "encrypt"
  # active = false
  # dst_address = "10.99.0.0/24"
  # dst_port = "443"
  # group = "replace-me"
  # ipsec_protocols = "esp"
  # level = "unique"
  # nopeer = "replace-me"
  # notemplate = "replace-me"
  # peer = "replace-me"
  # proposal = "replace-me"
  # protocol = "icmp"
  # src_address = "10.99.0.0/24"
  # src_port = "443"
  # template = false
  # tunnel = false
}
