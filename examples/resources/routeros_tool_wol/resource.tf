resource "routeros_tool_wol" "wol_example" {
  # router = "my-router"  # which router to target; omit for the default
  mac = "02:00:00:00:00:01"
}
