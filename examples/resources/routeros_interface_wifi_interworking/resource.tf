resource "routeros_interface_wifi_interworking" "interworking_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # x3gpp_info = "replace-me"
  # x3gpp_info_raw = "replace-me"
  # authentication_types = "replace-me"
  # connection_capabilities = "replace-me"
  # dgaf = "replace-me"
  # domain_names = "replace-me"
  # esr = "replace-me"
  # hessid = "replace-me"
  # hotspot_2_0 = "replace-me"
  # internet = "replace-me"
  # ipv4_availability = "replace-me"
  # ipv6_availability = "replace-me"
  # network_type = "replace-me"
  # operational_classes = "replace-me"
  # operator_names = "replace-me"
  # realms = "replace-me"
  # realms_raw = "replace-me"
  # roaming_ois = "replace-me"
  # uesa = "replace-me"
  # venue = "replace-me"
  # venue_names = "replace-me"
  # wan_at_capacity = "replace-me"
  # wan_downlink = "replace-me"
  # wan_downlink_load = "replace-me"
  # wan_measurement_duration = "replace-me"
  # wan_status = "replace-me"
  # wan_symmetric = "replace-me"
  # wan_uplink = "replace-me"
  # wan_uplink_load = "replace-me"
}
