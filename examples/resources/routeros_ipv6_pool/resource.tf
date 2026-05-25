resource "routeros_ipv6_pool" "pool_example" {
  # router = "my-router"  # which router to target; omit for the default
  name          = "example"
  prefix        = "fd00:db8::/56"
  prefix_length = 64

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # from_pool = "replace-me"
}
