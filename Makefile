-include Makefile.local.mk

build:
	go build

generate-docs:
	go generate ./...
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.25.0 validate --provider-name slr
