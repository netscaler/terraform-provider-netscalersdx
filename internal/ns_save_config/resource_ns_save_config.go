package ns_save_config

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*nsSaveConfigResource)(nil)
var _ resource.ResourceWithConfigure = (*nsSaveConfigResource)(nil)

func NsSaveConfigResource() resource.Resource {
	return &nsSaveConfigResource{}
}

type nsSaveConfigResource struct {
	client *service.NitroClient
}

func (r *nsSaveConfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ns_save_config"
}

// Configure configures the client resource.
func (r *nsSaveConfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *nsSaveConfigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = nsSaveConfigResourceSchema(ctx)
}

func (r *nsSaveConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "[DEBUG] In Create Method of nsSaveConfigResource")

	var data nsSaveConfigModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	nsSaveConfigPayload := make(map[string]interface{})
	stringList := utils.TypeListToUnmarshalStringList(data.NsIpAddressArr)

	nsSaveConfigPayload["ns_ip_address_arr"] = stringList

	endpoint := "ns_save_config"

	// Create the request
	_, err := r.client.AddResource(endpoint, nsSaveConfigPayload)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resId := utils.PrefixedUniqueId("ns_save_config-")
	data.Id = types.StringValue(resId)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *nsSaveConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *nsSaveConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *nsSaveConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
