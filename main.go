package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/ebogdum/terraform-provider-routeros/internal/provider"
)

// Overridden via -ldflags at build/release time.
var (
	version = "dev"
	commit  = "none"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "run as debug server (attach Terraform via TF_REATTACH_PROVIDERS)")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/ebogdum/routeros",
		Debug:   debug,
	}

	_ = commit // surfaced via -ldflags; reserved for `provider --version` output

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err)
	}
}
