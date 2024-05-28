package ntp_sync

import (
	"context"
	"fmt"
	"time"

	"terraform-provider-netscalersdx/internal/service"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = (*ntpSyncResource)(nil)
var _ resource.ResourceWithConfigure = (*ntpSyncResource)(nil)

func NtpSyncResource() resource.Resource {
	return &ntpSyncResource{}
}

type ntpSyncResource struct {
	client *service.NitroClient
}

func (r *ntpSyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntp_sync"
}

// Configure configures the client resource.
func (r *ntpSyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ntpSyncResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ntpSyncResourceSchema()
}

func (r *ntpSyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of ntp_sync Resource")

	var data ntpSyncModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	ntpSyncReq := ntpSyncGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "ntp_sync"

	// Create the request
	_, _ = r.client.UpdateResource(endpoint, ntpSyncReq, "")


	// Wait for the resource to reboot
	for {
		time.Sleep(5 * time.Second)

		n := service.NitroRequestParams{
			Resource:           "ntp_sync",
			ResourcePath:       "nitro/v1/config/ntp_sync",
			Method:             "GET",
			SuccessStatusCodes: []int{200},
		}

		_, err := r.client.MakeNitroRequest(n)

		if err == nil {
			time.Sleep(5 * time.Second)
			break
		}
	}

	resID := utils.PrefixedUniqueId("ntp_sync-")

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

func (r *ntpSyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of ntp_sync Resource with Id: %s", resId))

	var data ntpSyncModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ntp_sync"

	responseData, err := r.client.GetAllResource(endpoint)
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ntp_sync: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	ntpSyncSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ntpSyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *ntpSyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
