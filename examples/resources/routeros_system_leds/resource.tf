resource "routeros_system_leds" "leds_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # interface = "ether1"
  # leds = "replace-me"
  # type = "replace-me"
}
