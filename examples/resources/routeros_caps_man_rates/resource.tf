resource "routeros_caps_man_rates" "rates_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"
}
