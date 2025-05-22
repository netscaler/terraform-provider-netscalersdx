package cipher_group

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*cipherGroupResource)(nil)
var _ resource.ResourceWithConfigure = (*cipherGroupResource)(nil)
var _ resource.ResourceWithImportState = (*cipherGroupResource)(nil)

func CipherGroupResource() resource.Resource {
	return &cipherGroupResource{}
}

type cipherGroupResource struct {
	client *service.NitroClient
}

func (r *cipherGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *cipherGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cipher_group"
}

// Configure configures the client resource.
func (r *cipherGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *cipherGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = cipherGroupResourceSchema(ctx)
}

func (r *cipherGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of cipher_group Resource")

	var data cipherGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	cipherGroupReq := cipherGroupGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "cipher_group"

	// Create the request
	_, err := r.client.AddResource(endpoint, cipherGroupReq)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resID := data.CipherGroupName.ValueString()

	data.Id = types.StringValue(resID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	rreq := resource.ReadRequest{
		State:        resp.State,
		ProviderMeta: req.ProviderMeta,
	}
	rresp := resource.ReadResponse{
		State:       resp.State,
		Diagnostics: resp.Diagnostics,
	}

	r.Read(ctx, rreq, &rresp)

	*resp = resource.CreateResponse{
		State:       rresp.State,
		Diagnostics: rresp.Diagnostics,
	}

}

func (r *cipherGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of cipher_group Resource with Id: %s", resId))

	var data cipherGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "cipher_group"

	returnArr, err := r.client.GetAllResource(endpoint)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error getting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	if len(returnArr) == 0 {
		resp.Diagnostics.AddError(
			fmt.Sprintf("No Resource found: %s", endpoint),
			fmt.Sprintf("GetAllResource returned empty list"),
		)
		return
	}

	foundIndex := -1

	returnData := returnArr[endpoint].([]interface{})
	for i, v := range returnData {
		m := v.(map[string]interface{})
		if m["cipher_group_name"] == resId.ValueString() {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		resp.Diagnostics.AddError(
			fmt.Sprintf("No Resource found: %s", endpoint),
			fmt.Sprintf("No cipher_group with %s cipher_group_name", resId),
		)
		return
	}

	resourceData := returnArr[endpoint].([]interface{})[foundIndex].(map[string]interface{})

	data.CipherGroupDescription = types.StringValue(resourceData["cipher_group_description"].(string))
	data.CipherGroupName = types.StringValue(resourceData["cipher_group_name"].(string))
	data.CipherNameListArray = utils.StringListToTypeSet(utils.ToStringList(resourceData["cipher_name_list_array"].([]interface{})))

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *cipherGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of cipher_group Resource")

	var data cipherGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state cipherGroupModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "cipher_group"
	requestPayload := cipherGroupGetThePayloadFromtheConfig(ctx, &data)
	data.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, resourceId)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Resource",
			fmt.Sprintf("Error updating resource: %s", err.Error()),
		)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *cipherGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of cipher_group Resource")

	var data cipherGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := "cipher_group"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

}
