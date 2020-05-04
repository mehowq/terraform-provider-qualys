package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/mehowq/terraform-provider-qualys/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
