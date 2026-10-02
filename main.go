package main

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.25.0 generate --provider-name slr

import "github.com/sikalabsx/terraform-provider-slr/internal/provider"

var version = "dev"

func main() {
	provider.Serve(version)
}
