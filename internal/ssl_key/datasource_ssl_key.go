package ssl_key

import (
	"context"
	"fmt"
	"net/url"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*sslKeyDataSource)(nil)

func SslKeyDataSource() datasource.DataSource {
	return &sslKeyDataSource{}
}

type sslKeyDataSource struct {
	client *service.NitroClient
}

func (d *sslKeyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ssl_key"
}

func (d *sslKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *sslKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = sslKeyDataSourceSchema()
}

func (d *sslKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data sslKeyModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ssl_key"

	if data.Id.IsNull() || data.Id.ValueString() == "" {
		resp.Diagnostics.AddError(
			"Missing ssl_key ID",
			"'id' (the SSL key file_name on the SDX appliance) must be provided to look up the ssl_key data source.",
		)
		return
	}

	resid := data.Id.ValueString()

	responseData, err := d.client.GetResourceV2(endpoint, url.PathEscape(resid))
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s: %v is not present in the remote", endpoint, resid),
		)
		return
	}

	data.Id = types.StringValue(resid)
	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	sslKeySetAttrFromGet(ctx, &data, getResponseData)

	// For the data source, file_location_path is the server-side path the SDX
	// reports — that's what the consumer of the data source actually wants to
	// see (unlike the resource, where overwriting the operator's local-path
	// input would cause drift). Populate it explicitly here.
	if v, ok := getResponseData["file_location_path"]; ok && v != nil {
		data.FileLocationPath = types.StringValue(toString(v))
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
