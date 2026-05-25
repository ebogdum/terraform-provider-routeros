resource "routeros_ip_kid_control_device" "device_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # mac_address = "10.99.0.0/24"
  # name = "tf-example"
  # user = "myuser"
}
