package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"gitlab.com/sport24ru/terraform-provider-dex/private/provider"
	"log"
)

func main() {
	if err := providerserver.Serve(context.Background(), provider.NewDexProvider, providerserver.ServeOpts{
		Address: "registry.terraform.io/sport24ru/dex",
	}); err != nil {
		log.Fatal(err)
	}
}
