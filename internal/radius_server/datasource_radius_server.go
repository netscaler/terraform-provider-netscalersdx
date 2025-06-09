package radius_server

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*radiusServerDataSource)(nil)

func RadiusServerDataSource() datasource.DataSource {
	return &radiusServerDataSource{}
}

type radiusServerDataSource struct {
	client *service.NitroClient
}

func (d *radiusServerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_server"
}

func (d *radiusServerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *radiusServerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = radiusServerDataSourceSchema()
}

func (d *radiusServerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data radiusServerModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var responseData map[string]interface{}
	var err error
	// Read API call logic

	endpoint := "radius_server"
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
		data.Id = data.Name

	} else {
		responseData, err = d.client.GetResource(endpoint, data.Id.ValueString())
		if err != nil {
			resp.State.RemoveResource(ctx)
			tflog.Warn(ctx, fmt.Sprintf("removing resource radius_server: %v from state because it is not present in the remote", data.Id.ValueString()))
			return
		}
		data.Id = data.Id
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	radiusServerSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
