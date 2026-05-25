package provider

import "github.com/hashicorp/terraform-plugin-framework/path"

var (
	pathHost     = path.Root("host")
	pathUsername = path.Root("username")
	pathRouters  = path.Root("routers")
)

var _ = pathHost
var _ = pathUsername
