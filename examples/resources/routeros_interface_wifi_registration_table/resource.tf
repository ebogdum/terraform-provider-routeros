resource "routeros_interface_wifi_registration_table" "registration_table_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false
}
