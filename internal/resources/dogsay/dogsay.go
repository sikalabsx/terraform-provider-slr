package dogsay

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/sikalabs/dogsay/pkg/dogsay"
)

var _ resource.Resource = &dogSayResource{}

type dogSayResource struct{}

type dogSayModel struct {
	ID     types.String `tfsdk:"id"`
	Text   types.String `tfsdk:"text"`
	Output types.String `tfsdk:"output"`
}

func NewResource() resource.Resource {
	return &dogSayResource{}
}

func (r *dogSayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dogsay"
}

func (r *dogSayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Renders text as ASCII art of a dog saying it (like cowsay) using [dogsay](https://github.com/sikalabs/dogsay).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Unique identifier (UUID) of the resource.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"text": schema.StringAttribute{
				MarkdownDescription: "Text the dog should say.",
				Required:            true,
			},
			"output": schema.StringAttribute{
				MarkdownDescription: "Rendered ASCII art output.",
				Computed:            true,
			},
		},
	}
}

func (r *dogSayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data dogSayModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(uuid.NewString())
	data.Output = types.StringValue(dogsay.DogSay(data.Text.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *dogSayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *dogSayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data dogSayModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Output = types.StringValue(dogsay.DogSay(data.Text.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *dogSayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
