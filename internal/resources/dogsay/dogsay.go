package dogsay

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/sikalabs/dogsay/pkg/dogsay"
)

func ResourceDogSay() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"text": {
				Type:     schema.TypeString,
				Required: true,
			},
			"output": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: false,
			},
		},

		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			d.SetId(uuid.NewString())
			text := d.Get("text").(string)
			_ = d.Set("output", dogsay.DogSay(text))
			return nil
		},
		UpdateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			text := d.Get("text").(string)
			_ = d.Set("output", dogsay.DogSay(text))
			return nil
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			d.SetId("")
			return nil
		},
	}
}
