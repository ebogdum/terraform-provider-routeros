resource "routeros_ip_pool" "pool_example" {
  # router = "my-router"  # which router to target; omit for the default
  name   = "tf-example"
  ranges = "10.99.0.100-10.99.0.200"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # next_pool = "replace-me"
}
