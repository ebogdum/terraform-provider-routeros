resource "routeros_tool_fetch" "fetch_example" {
  # router = "my-router"  # which router to target; omit for the default
  url = "https://example.com"

  # Optional attributes (uncomment as needed):
  # mode = "http"
}
