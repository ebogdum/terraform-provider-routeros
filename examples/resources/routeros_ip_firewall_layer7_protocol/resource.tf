resource "routeros_ip_firewall_layer7_protocol" "layer7_protocol_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
}
