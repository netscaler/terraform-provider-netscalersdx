package device_group

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*deviceGroupDataSource)(nil)

func DeviceGroupDataSource() datasource.DataSource {
	return &deviceGroupDataSource{}
}

type deviceGroupDataSource struct {
	client *service.NitroClient
}

func (d *deviceGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_group"
}

func (d *deviceGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *deviceGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = deviceGroupDataSourceSchema()
}

func (d *deviceGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceGroupModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "device_group"
	var responseData map[string]interface{}
	var err error
	// Read API call logic

	if !data.Name.IsNull() {
		args := map[string]string{
			"name": data.Name.ValueString(),
		}

		responseData, err = d.client.GetResourceWithArgs(endpoint, args)
		if err != nil {
			resp.State.RemoveResource(ctx)
			tflog.Warn(ctx, fmt.Sprintf("removing resource radius_server: %v from state because it is not present in the remote", data.Id.ValueString()))
			return
		}
	} else {
		responseData, err = d.client.GetResource(endpoint, data.Id.ValueString())
		if err != nil {
			resp.State.RemoveResource(ctx)
			tflog.Warn(ctx, fmt.Sprintf("removing resource radius_server: %v from state because it is not present in the remote", data.Id.ValueString()))
			return
		}
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	data.Id = types.StringValue(getResponseData["id"].(string))

	deviceGroupSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
