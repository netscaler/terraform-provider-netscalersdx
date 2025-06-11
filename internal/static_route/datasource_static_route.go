package static_route

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*staticRouteDataSource)(nil)

func StaticRouteDataSource() datasource.DataSource {
	return &staticRouteDataSource{}
}

type staticRouteDataSource struct {
	client *service.NitroClient
}

func (d *staticRouteDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_static_route"
}

func (d *staticRouteDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *staticRouteDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = staticRouteDataSourceSchema()
}

func (d *staticRouteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data staticRouteModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "static_route"

	var responseData map[string]interface{}
	var err error

	if !data.Network.IsNull() {
		args := map[string]string{
			"network": data.Network.ValueString(),
		}

		responseData, err = d.client.GetResourceWithArgs(endpoint, args)
		if err != nil {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("%s: %v is not present in the remote", endpoint, data.Network.ValueString()),
			)
			return
		}
	} else {
		responseData, err = d.client.GetResource(endpoint, data.Id.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Resource Not Found",
				fmt.Sprintf("%s: %v is not present in the remote", endpoint, data.Id.ValueString()),
			)
			return
		}
	}

	if len(responseData[endpoint].([]interface{})) < 1 {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("No %s resource found in the remote for the given query.", endpoint),
		)
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	data.Gateway = types.StringValue(getResponseData["gateway"].(string))
	data.Netmask = types.StringValue(getResponseData["netmask"].(string))
	data.Network = types.StringValue(getResponseData["network"].(string))
	data.Id = types.StringValue(getResponseData["id"].(string))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
