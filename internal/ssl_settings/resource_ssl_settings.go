package ssl_settings

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

var _ resource.Resource = (*sslSettingsResource)(nil)
var _ resource.ResourceWithConfigure = (*sslSettingsResource)(nil)

func SslSettingsResource() resource.Resource {
	return &sslSettingsResource{}
}

type sslSettingsResource struct {
	client *service.NitroClient
}

func (r *sslSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_settings"
}

// Configure configures the client resource.
func (r *sslSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *sslSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = sslSettingsResourceSchema()
}

func (r *sslSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "In Create Method of ssl_settings Resource")

	var data sslSettingsModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	sslSettingsReq := sslSettingsGetThePayloadFromtheConfig(ctx, &data)

	endpoint := "ssl_settings"

	_, _ = r.client.UpdateResource(endpoint, sslSettingsReq, "")

	// Wait for the resource to come up
	for {
		time.Sleep(5 * time.Second)

		n := service.NitroRequestParams{
			Resource:           "ssl_settings",
			ResourcePath:       "nitro/v1/config/ssl_settings",
			Method:             "GET",
			SuccessStatusCodes: []int{200},
		}

		_, err := r.client.MakeNitroRequest(n)

		if err == nil {
			time.Sleep(5 * time.Second)
			break
		}
	}

	resID := utils.PrefixedUniqueId("ssl_settings-")

	// Example data value setting
	data.Id = types.StringValue(resID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	// time.Sleep(20 * time.Second)

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

func (r *sslSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var resId types.String
	req.State.GetAttribute(ctx, path.Root("id"), &resId)
	tflog.Debug(ctx, fmt.Sprintf("In Read Method of ssl_settings Resource with Id: %s", resId))

	var data sslSettingsModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ssl_settings"

	responseData, err := r.client.GetAllResource(endpoint)
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ssl_settings: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	sslSettingsSetAttrFromGet(ctx, &data, getResponseData)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *sslSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *sslSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
