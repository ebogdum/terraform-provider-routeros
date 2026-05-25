resource "routeros_ip_cloud_back_to_home_user" "back_to_home_user_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # active = false
  # allow_lan = false
  # expires = "replace-me"
  # file_access_mode = ""
  # files = "replace-me"
  # name = "tf-example"
  # newe = "replace-me"
  # newfileman = "replace-me"
  # notnew = "replace-me"
  # oldfileman = "replace-me"
  # private_key = "REDACTED"
  # public_key = "REDACTED"
}
