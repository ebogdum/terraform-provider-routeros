resource "routeros_snmp_community" "community_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # addresses = "10.99.0.0/24"
  # authentication_password = "REDACTED"
  # authentication_protocol = "MD5"
  # encryption_password = "REDACTED"
  # encryption_protocol = "DES"
  # read_access = true
  # security = "none"
  # write_access = false
}
