package device_group

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

var _ resource.Resource = (*deviceGroupResource)(nil)
var _ resource.ResourceWithConfigure = (*deviceGroupResource)(nil)

func DeviceGroupResource() resource.Resource {
	return &deviceGroupResource{}
}

type deviceGroupResource struct {
	client *service.NitroClient
}

func (r *deviceGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_group"
}

// Configure configures the client resource.
func (r *deviceGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *deviceGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = deviceGroupResourceSchema(ctx)
}

func (r *deviceGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of device_group Resource")

	var data deviceGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	deviceGroupReq := deviceGroupGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "device_group"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, deviceGroupReq)

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

func (r *deviceGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of device_group Resource with Id: %s", resId))

	var data deviceGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "device_group"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource device_group: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.Category.IsNull() {
		data.Category = types.StringValue(getResponseData["category"].(string))
	}
	if !data.CriteriaCondn.IsNull() {
		data.CriteriaCondn = types.StringValue(getResponseData["criteria_condn"].(string))
	}
	if !data.CriteriaType.IsNull() {
		data.CriteriaType = types.StringValue(getResponseData["criteria_type"].(string))
	}
	if !data.CriteriaValue.IsNull() {
		data.CriteriaValue = types.StringValue(getResponseData["criteria_value"].(string))
	}
	if !data.DeviceFamily.IsNull() {
		data.DeviceFamily = types.StringValue(getResponseData["device_family"].(string))
	}
	if !data.DisableUpgrade.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["disable_upgrade"].(string))
		data.DisableUpgrade = types.BoolValue(val)
	}
	if !data.Duration.IsNull() {
		val, _ := strconv.Atoi(getResponseData["duration"].(string))
		data.Duration = types.Int64Value(int64(val))
	}
	if !data.LockAcquireTime.IsNull() {
		data.LockAcquireTime = types.StringValue(getResponseData["lock_acquire_time"].(string))
	}
	if !data.LockAcquiringDevice.IsNull() {
		data.LockAcquiringDevice = types.StringValue(getResponseData["lock_acquiring_device"].(string))
	}
	if !data.MaintenanceWindowStart.IsNull() {
		data.MaintenanceWindowStart = types.StringValue(getResponseData["maintenance_window_start"].(string))
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.StaticDeviceList.IsNull() {
		data.StaticDeviceList = types.StringValue(getResponseData["static_device_list"].(string))
	}
	if !data.UpgradeLock.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["upgrade_lock"].(string))
		data.UpgradeLock = types.BoolValue(val)
	}
	if !data.UpgradeVersion.IsNull() {
		data.UpgradeVersion = types.StringValue(getResponseData["upgrade_version"].(string))
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *deviceGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of device_group Resource")

	var data deviceGroupModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state deviceGroupModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "device_group"
	requestPayload := deviceGroupGetThePayloadFromtheConfig(ctx, &data)
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

func (r *deviceGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of device_group Resource")

	var data deviceGroupModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "device_group"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
