package provider

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/sikalabsx/terraform-provider-slr/internal/resources/dogsay"
	"github.com/sikalabsx/terraform-provider-slr/internal/resources/hello"
)

func Serve(version string) {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/sikalabsx/slr",
		Debug:   debug,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}

var _ provider.Provider = &slrProvider{}

type slrProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &slrProvider{version: version}
	}
}

func (p *slrProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "slr"
	resp.Version = p.version
}

func (p *slrProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Terraform provider `slr` by SikaLabs with simple example resources.",
	}
}

func (p *slrProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *slrProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		hello.NewResource,
		dogsay.NewResource,
	}
}

func (p *slrProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}
