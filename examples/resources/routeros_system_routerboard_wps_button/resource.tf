resource "routeros_system_routerboard_wps_button" "this" {
  # router = "my-router"  # which router to target; omit for the default
  enabled   = true
  hold_time = "0s..1m"
  on_event  = "wps-accept"
}
