package syslog_server

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

var _ resource.Resource = (*syslogServerResource)(nil)
var _ resource.ResourceWithConfigure = (*syslogServerResource)(nil)

func SyslogServerResource() resource.Resource {
	return &syslogServerResource{}
}

type syslogServerResource struct {
	client *service.NitroClient
}

func (r *syslogServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_syslog_server"
}

// Configure configures the client resource.
func (r *syslogServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *syslogServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = syslogServerResourceSchema(ctx)
}

func (r *syslogServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of syslog_server Resource")

	var data syslogServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	syslogServerReq := syslogServerGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "syslog_server"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, syslogServerReq)

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["name"].(string)

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

func (r *syslogServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of syslog_server Resource with Id: %s", resId))

	var data syslogServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "syslog_server"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource syslog_server: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	}
	if !data.LogLevelAll.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["log_level_all"].(string))
		data.LogLevelAll = types.BoolValue(val)
	}
	if !data.LogLevelCritical.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["log_level_critical"].(string))
		data.LogLevelCritical = types.BoolValue(val)
	}
	if !data.LogLevelError.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["log_level_error"].(string))
		data.LogLevelError = types.BoolValue(val)
	}
	if !data.LogLevelInfo.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["log_level_info"].(string))
		data.LogLevelInfo = types.BoolValue(val)
	}
	if !data.LogLevelNone.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["log_level_none"].(string))
		data.LogLevelNone = types.BoolValue(val)
	}
	if !data.LogLevelWarning.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["log_level_warning"].(string))
		data.LogLevelWarning = types.BoolValue(val)
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

func (r *syslogServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of syslog_server Resource")

	var data syslogServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state syslogServerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "syslog_server"
	requestPayload := syslogServerGetThePayloadFromtheConfig(ctx, &data)
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

func (r *syslogServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In CrDeleteeate Method of syslog_server Resource")

	var data syslogServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "syslog_server"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
