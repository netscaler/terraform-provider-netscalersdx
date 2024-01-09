package tacacs_server

import (
	"context"
	"fmt"
	"strconv"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*tacacsServerResource)(nil)
var _ resource.ResourceWithConfigure = (*tacacsServerResource)(nil)

func TacacsServerResource() resource.Resource {
	return &tacacsServerResource{}
}

type tacacsServerResource struct {
	client *service.NitroClient
}

func (r *tacacsServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tacacs_server"
}

// Configure configures the client resource.
func (r *tacacsServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *tacacsServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = tacacsServerResourceSchema(ctx)
}

func (r *tacacsServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of tacacs_server Resource")

	var data tacacsServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tacacsServerReq := tacacsServerGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "tacacs_server"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, tacacsServerReq)

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

func (r *tacacsServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of tacacs_server Resource with Id: %s", resId))

	var data tacacsServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "tacacs_server"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource tacacs_server: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.Accounting.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["accounting"].(string))
		data.Accounting = types.BoolValue(val)
	}
	if !data.AuthTimeout.IsNull() {
		val, _ := strconv.Atoi(getResponseData["auth_timeout"].(string))
		data.AuthTimeout = types.Int64Value(int64(val))
	}
	if !data.GroupAttrName.IsNull() {
		data.GroupAttrName = types.StringValue(getResponseData["group_attr_name"].(string))
	}
	if !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.Port.IsNull() {
		val, _ := strconv.Atoi(getResponseData["port"].(string))
		data.Port = types.Int64Value(int64(val))
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tacacsServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of tacacs_server Resource")

	var data tacacsServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state tacacsServerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "tacacs_server"
	requestPayload := tacacsServerGetThePayloadFromtheConfig(ctx, &data)
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

func (r *tacacsServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In CrDeleteeate Method of tacacs_server Resource")

	var data tacacsServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "tacacs_server"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
