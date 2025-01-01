package mps_ssl_certkey

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

var _ resource.Resource = (*mpsSslCertkeyResource)(nil)
var _ resource.ResourceWithConfigure = (*mpsSslCertkeyResource)(nil)
var _ resource.ResourceWithImportState = (*mpsSslCertkeyResource)(nil)

func MpsSslCertkeyResource() resource.Resource {
	return &mpsSslCertkeyResource{}
}

type mpsSslCertkeyResource struct {
	client *service.NitroClient
}

// ImportState implements resource.ResourceWithImportState.
func (r *mpsSslCertkeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *mpsSslCertkeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mps_ssl_certkey"
}

// Configure configures the client resource.
func (r *mpsSslCertkeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *mpsSslCertkeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = mpsSslCertkeyResourceSchema()
}

func (r *mpsSslCertkeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of mps_ssl_certkey Resource")

	var data mpsSslCertkeyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	mpsSslCertkeyReq := mpsSslCertkeyGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "mps_ssl_certkey"

	_, _ = r.client.AddResource(endpoint, mpsSslCertkeyReq)

	// Wait for the resource to reboot
	for {
		time.Sleep(5 * time.Second)

		n := service.NitroRequestParams{
			Resource:           "mps_ssl_certkey",
			ResourcePath:       "nitro/v1/config/mps_ssl_certkey",
			Method:             "GET",
			SuccessStatusCodes: []int{200},
		}

		_, err := r.client.MakeNitroRequest(n)
		tflog.Debug(ctx, fmt.Sprintf("mps_ssl_certkey resource is not yet available: %v", err))

		if err == nil {
			time.Sleep(10 * time.Second)
			break
		}
	}

	resID := utils.PrefixedUniqueId("mps_ssl_certkey-")

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

func (r *mpsSslCertkeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of mps_ssl_certkey Resource with Id: %s", resId))

	var data mpsSslCertkeyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "mps_ssl_certkey"

	responseData, err := r.client.GetAllResource(endpoint)
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource mps_ssl_certkey: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	mpsSslCertkeySetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *mpsSslCertkeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *mpsSslCertkeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
