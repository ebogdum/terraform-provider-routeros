resource "routeros_routing_table" "table_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # fib = "replace-me"
  # name = "tf-example"
}
