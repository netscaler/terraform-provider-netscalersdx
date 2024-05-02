package snmp_mib

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

var _ resource.Resource = (*snmpMibResource)(nil)
var _ resource.ResourceWithConfigure = (*snmpMibResource)(nil)

func SnmpMibResource() resource.Resource {
	return &snmpMibResource{}
}

type snmpMibResource struct {
	client *service.NitroClient
}

func (r *snmpMibResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_mib"
}

// Configure configures the client resource.
func (r *snmpMibResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *snmpMibResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = snmpMibResourceSchema()
}

func (r *snmpMibResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of snmp_mib Resource")

	var data snmpMibModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	snmpMibReq := snmpMibGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "snmp_mib"

	// Create the request
	_, err := r.client.UpdateResource(endpoint, snmpMibReq, "")

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error creating resource: %s", endpoint),
			fmt.Sprintf("Error: %s", err.Error()),
		)
		return
	}

	// resID := returnData[endpoint].([]interface{})[0].(map[string]interface{})["id"].(string)
	resID := utils.PrefixedUniqueId("snmp_mib-")

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

func (r *snmpMibResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of snmp_mib Resource with Id: %s", resId))

	var data snmpMibModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "snmp_mib"

	responseData, err := r.client.GetAllResource(endpoint)
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource snmp_mib: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	snmpMibSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *snmpMibResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "In Update Method of snmp_mib Resource")

	var data snmpMibModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state snmpMibModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceId := state.Id.ValueString()
	endpoint := "snmp_mib"
	requestPayload := snmpMibGetThePayloadFromtheConfig(ctx, &data)
	data.Id = state.Id

	_, err := r.client.UpdateResource(endpoint, requestPayload, "")

	if err != nil {
		resp.Diagnostics.AddError(
			fmt.Sprintf("Error updating resource: %s", resourceId),
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

func (r *snmpMibResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
