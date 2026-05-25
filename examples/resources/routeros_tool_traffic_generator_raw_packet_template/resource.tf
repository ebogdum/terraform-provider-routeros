resource "routeros_tool_traffic_generator_raw_packet_template" "raw_packet_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # data = "uninitialized"
  # data_byte = 0
  # header = "replace-me"
  # ip_header_offset = "replace-me"
  # ipv6_header_offset = "replace-me"
  # name = "example"
  # port = "443"
  # random_byte_offsets_and_masks = "replace-me"
  # random_ranges = "replace-me"
  # special_footer = false
  # udp_header_offset = "replace-me"
}
