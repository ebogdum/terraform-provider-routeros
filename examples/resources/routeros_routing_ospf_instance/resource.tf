resource "routeros_routing_ospf_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment  = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # domain_id = "replace-me"
  # domain_tag = "replace-me"
  # in_filter = "replace-me"
  # mpls_te_address = "10.99.0.0/24"
  # mpls_te_area = "replace-me"
  # name = "tf-example"
  # originate_default = "replace-me"
  # out_filter = "replace-me"
  # out_filter_select = "replace-me"
  # redistribute = "replace-me"
  # router_id = "replace-me"
  # routing_table = "main"
  # version = "2"
  # vrf = "main"
}
