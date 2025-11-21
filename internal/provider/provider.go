package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/sikalabsx/terraform-provider-slr/internal/resources/dogsay"
	"github.com/sikalabsx/terraform-provider-slr/internal/resources/hello"
)

func Serve() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return Provider()
		},
	})
}

type Config struct{}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		ResourcesMap: map[string]*schema.Resource{
			"slr_hello":  hello.ResourceHello(),
			"slr_dogsay": dogsay.ResourceDogSay(),
		},

		ConfigureContextFunc: func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
			var diags diag.Diagnostics
			cfg := &Config{}
			return cfg, diags
		},
	}
}
