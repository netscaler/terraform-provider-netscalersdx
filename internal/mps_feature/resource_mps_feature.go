package mps_feature

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*mpsFeatureResource)(nil)
var _ resource.ResourceWithConfigure = (*mpsFeatureResource)(nil)
var _ resource.ResourceWithImportState = (*mpsFeatureResource)(nil)

func MpsFeatureResource() resource.Resource {
	return &mpsFeatureResource{}
}

type mpsFeatureResource struct {
	client *service.NitroClient
}

func (r *mpsFeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *mpsFeatureResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mps_feature"
}

// Configure configures the client resource.
func (r *mpsFeatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *mpsFeatureResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = mpsFeatureResourceSchema()
}

func (r *mpsFeatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of mps_feature Resource")

	var data mpsFeatureModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	mpsFeatureReq := mpsFeatureGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "mps_feature"

	// Create the request
	_, err := r.client.UpdateResource(endpoint, mpsFeatureReq, "")

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	// Example data value setting
	data.Id = types.StringValue(data.FeatureName.ValueString())

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

func (r *mpsFeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of mps_feature Resource with Id: %s", resId))

	var data mpsFeatureModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "mps_feature"

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
			"GetAllResource returned empty list",
		)
		return
	}

	foundIndex := -1
	returnData := returnArr[endpoint].([]interface{})
	for i, v := range returnData {
		m := v.(map[string]interface{})
		if m["feature_name"] == resId.ValueString() {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		resp.Diagnostics.AddError(
			fmt.Sprintf("No Resource found: %s", endpoint),
			fmt.Sprintf("No mps_feature with %s", resId.ValueString()),
		)
		return
	}

	getResponseData := returnArr[endpoint].([]interface{})[foundIndex].(map[string]interface{})

	mpsFeatureSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *mpsFeatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of mps_feature Resource")

	var data mpsFeatureModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state mpsFeatureModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// resourceId := state.Id.ValueString()
	endpoint := "mps_feature"
	requestPayload := mpsFeatureGetThePayloadFromtheConfig(ctx, &data)
	data.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, "")

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

	rreq := resource.ReadRequest{
		State:        resp.State,
		ProviderMeta: req.ProviderMeta,
	}
	rresp := resource.ReadResponse{
		State:       resp.State,
		Diagnostics: resp.Diagnostics,
	}

	r.Read(ctx, rreq, &rresp)

	*resp = resource.UpdateResponse{
		State:       rresp.State,
		Diagnostics: rresp.Diagnostics,
	}
}

func (r *mpsFeatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of mps_feature Resource")
}
