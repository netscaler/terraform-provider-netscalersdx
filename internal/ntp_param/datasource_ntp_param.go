package ntp_param

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"
	"terraform-provider-netscalersdx/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*ntpParamDataSource)(nil)

func NtpParamDataSource() datasource.DataSource {
	return &ntpParamDataSource{}
}

type ntpParamDataSource struct {
	client *service.NitroClient
}

func (d *ntpParamDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntp_param"
}

func (d *ntpParamDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ntpParamDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ntpParamDataSourceSchema()
}

func (d *ntpParamDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ntpParamModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "ntp_param"

	responseData, err := d.client.GetAllResource(endpoint)
	if err != nil {
		resp.State.RemoveResource(ctx)
		tflog.Warn(ctx, fmt.Sprintf("removing resource ntp_param: %v from state because it is not present in the remote", data.Id.ValueString()))
		return
	}

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	data.Authentication = types.BoolValue(utils.StringToBool(getResponseData["authentication"].(string)))
	data.AutomaxLogsec = types.Int64Value(utils.StringToInt(getResponseData["automax_logsec"].(string)))
	data.RevokeLogsec = types.Int64Value(utils.StringToInt(getResponseData["revoke_logsec"].(string)))
	data.TrustedKeyList = utils.StringListToTypeInt64List(utils.ToStringList(getResponseData["trusted_key_list"].([]interface{})))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
