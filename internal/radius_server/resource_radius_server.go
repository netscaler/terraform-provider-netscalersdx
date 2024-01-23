package radius_server

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

var _ resource.Resource = (*radiusServerResource)(nil)
var _ resource.ResourceWithConfigure = (*radiusServerResource)(nil)

func RadiusServerResource() resource.Resource {
	return &radiusServerResource{}
}

type radiusServerResource struct {
	client *service.NitroClient
}

func (r *radiusServerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_server"
}

// Configure configures the client resource.
func (r *radiusServerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *radiusServerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = radiusServerResourceSchema(ctx)
}

func (r *radiusServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of radius_server Resource")

	var data radiusServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	radiusServerReq := radiusServerGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "radius_server"

	// Create the request
	returnData, err := r.client.AddResource(endpoint, radiusServerReq)

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

func (r *radiusServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of radius_server Resource with Id: %s", resId))

	var data radiusServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "radius_server"

	responseData, err := r.client.GetResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource radius_server: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	if !data.Accounting.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["accounting"].(string))
		data.Accounting = types.BoolValue(val)
	}
	if !data.AddressType.IsNull() {
		val, _ := strconv.Atoi(getResponseData["address_type"].(string))
		data.AddressType = types.Int64Value(int64(val))
	}
	if !data.AttributeType.IsNull() {
		val, _ := strconv.Atoi(getResponseData["attribute_type"].(string))
		data.AttributeType = types.Int64Value(int64(val))
	}
	if !data.AuthTimeout.IsNull() {
		val, _ := strconv.Atoi(getResponseData["auth_timeout"].(string))
		data.AuthTimeout = types.Int64Value(int64(val))
	}
	if !data.DefaultAuthenticationGroup.IsNull() {
		data.DefaultAuthenticationGroup = types.StringValue(getResponseData["default_authentication_group"].(string))
	}
	if !data.EnableNasIp.IsNull() {
		val, _ := strconv.ParseBool(getResponseData["enable_nas_ip"].(string))
		data.EnableNasIp = types.BoolValue(val)
	}
	if !data.GroupSeparator.IsNull() {
		data.GroupSeparator = types.StringValue(getResponseData["group_separator"].(string))
	}
	if !data.GroupsPrefix.IsNull() {
		data.GroupsPrefix = types.StringValue(getResponseData["groups_prefix"].(string))
	}
	if !data.IpAddress.IsNull() {
		data.IpAddress = types.StringValue(getResponseData["ip_address"].(string))
	}
	if !data.IpAttributeType.IsNull() {
		val, _ := strconv.Atoi(getResponseData["ip_attribute_type"].(string))
		data.IpAttributeType = types.Int64Value(int64(val))
	}
	if !data.IpVendorId.IsNull() {
		val, _ := strconv.Atoi(getResponseData["ip_vendor_id"].(string))
		data.IpVendorId = types.Int64Value(int64(val))
	}
	if !data.Name.IsNull() {
		data.Name = types.StringValue(getResponseData["name"].(string))
	}
	if !data.NasId.IsNull() {
		data.NasId = types.StringValue(getResponseData["nas_id"].(string))
	}
	if !data.PassEncoding.IsNull() {
		data.PassEncoding = types.StringValue(getResponseData["pass_encoding"].(string))
	}
	if !data.Port.IsNull() {
		val, _ := strconv.Atoi(getResponseData["port"].(string))
		data.Port = types.Int64Value(int64(val))
	}
	if !data.PwdAttributeType.IsNull() {
		val, _ := strconv.Atoi(getResponseData["pwd_attribute_type"].(string))
		data.PwdAttributeType = types.Int64Value(int64(val))
	}
	if !data.PwdVendorId.IsNull() {
		val, _ := strconv.Atoi(getResponseData["pwd_vendor_id"].(string))
		data.PwdVendorId = types.Int64Value(int64(val))
	}
	if !data.VendorId.IsNull() {
		val, _ := strconv.Atoi(getResponseData["vendor_id"].(string))
		data.VendorId = types.Int64Value(int64(val))
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *radiusServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of radius_server Resource")

	var data radiusServerModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state radiusServerModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "radius_server"
	requestPayload := radiusServerGetThePayloadFromtheConfig(ctx, &data)
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

func (r *radiusServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "In Delete Method of radius_server Resource")

	var data radiusServerModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	endpoint := "radius_server"
	_, err := r.client.DeleteResource(endpoint, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error deleting resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}
}
