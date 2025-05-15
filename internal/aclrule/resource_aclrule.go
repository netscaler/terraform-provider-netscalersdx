package aclrule

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*aclruleResource)(nil)
var _ resource.ResourceWithConfigure = (*aclruleResource)(nil)
var _ resource.ResourceWithImportState = &aclruleResource{}

// ImportState implements resource.ResourceWithImportState.
func (r *aclruleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func AclruleResource() resource.Resource {
	return &aclruleResource{}
}

type aclruleResource struct {
	client *service.NitroClient
}

func (r *aclruleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aclrule"
}

// Configure configures the client resource.
func (r *aclruleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *aclruleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = aclruleResourceSchema()
}

func (r *aclruleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of aclrule Resource")

	var data aclruleModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	aclruleReq := aclruleGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "aclrule"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, aclruleReq)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)

	// Example data value setting
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

func (r *aclruleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of aclrule Resource with Id: %s", resId))

	var data aclruleModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "aclrule"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource aclrule: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	aclruleSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *aclruleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of aclrule Resource")

	var data aclruleModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state aclruleModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "aclrule"
	requestPayload := aclruleGetThePayloadFromtheConfig(ctx, &data)
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

func (r *aclruleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of aclrule Resource")

	var data aclruleModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "aclrule"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
