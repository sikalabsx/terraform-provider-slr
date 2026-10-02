package main

import "github.com/sikalabsx/terraform-provider-slr/internal/provider"

var version = "dev"

func main() {
	provider.Serve(version)
}
