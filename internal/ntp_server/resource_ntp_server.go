package ntp_server

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

var _ resource.Resource = (*ntpServerResource)(nil)
var _ resource.ResourceWithConfigure = (*ntpServerResource)(nil)

func NtpServerResource() resource.Resource {
	return &ntpServerResource{}
}

type ntpServerResource struct {
	client *service.NitroClient
}

func (r *ntpServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntp_server"
}

// Configure configures the client resource.
func (r *ntpServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ntpServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ntpServerResourceSchema(ctx)
}

func (r *ntpServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of ntp_server Resource")

	var data ntpServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ntpServerReq := ntpServerGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "ntp_server"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, ntpServerReq)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["server"].(string)

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

func (r *ntpServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of ntp_server Resource with Id: %s", resId))

	var data ntpServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ntp_server"

	dataArr, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ntp_server: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := dataArr[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.Autokey.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["autokey"].(string))
		data.Autokey = types.BoolValue(val)
	}
	if !data.KeyId.IsNull() {
		val, _ := strconv.Atoi(getResponseData["key_id"].(string))
		data.KeyId = types.Int64Value(int64(val))
	}
	if !data.Maxpoll.IsNull() {
		val, _ := strconv.Atoi(getResponseData["maxpoll"].(string))
		data.Maxpoll = types.Int64Value(int64(val))
	}
	if !data.Minpoll.IsNull() {
		val, _ := strconv.Atoi(getResponseData["minpoll"].(string))
		data.Minpoll = types.Int64Value(int64(val))
	}
	if !data.PreferredServer.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["preferred_server"].(string))
		data.PreferredServer = types.BoolValue(val)
	}
	if !data.Server.IsNull() {
		data.Server = types.StringValue(getResponseData["server"].(string))
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ntpServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of ntp_server Resource")

	var data ntpServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state ntpServerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	endpoint := "ntp_server"
	_, err := r.client.UpdateResource(endpoint, ntpServerGetThePayloadFromtheConfig(ctx, &data), state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error updating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	data.Id = state.Id

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ntpServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In CrDeleteeate Method of ntp_server Resource")

	var data ntpServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "ntp_server"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
