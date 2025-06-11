package mps_ssl_certkey

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*mpsSslCertkeyDataSource)(nil)

func MpsSslCertkeyDataSource() datasource.DataSource {
	return &mpsSslCertkeyDataSource{}
}

type mpsSslCertkeyDataSource struct {
	client *service.NitroClient
}

func (d *mpsSslCertkeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mps_ssl_certkey"
}

func (d *mpsSslCertkeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *mpsSslCertkeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = mpsSslCertkeyDataSourceSchema()
}

func (d *mpsSslCertkeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data mpsSslCertkeyModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "mps_ssl_certkey"

	responseData, err := d.client.GetAllResource(endpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s is not present in the remote", endpoint),
		)
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	mpsSslCertkeySetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
