package snmp_trap

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*snmpTrapDataSource)(nil)

func SnmpTrapDataSource() datasource.DataSource {
	return &snmpTrapDataSource{}
}

type snmpTrapDataSource struct {
	client *service.NitroClient
}

func (d *snmpTrapDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_trap"
}

func (d *snmpTrapDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *snmpTrapDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = snmpTrapDataSourceSchema()
}

func (d *snmpTrapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data snmpTrapModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "snmp_trap"

	var resid string
	if !data.DestServer.IsNull() {
		resid = data.DestServer.ValueString()
	} else if !data.Id.IsNull() {
		resid = data.Id.ValueString()
	}

	if resid == "" {
		resp.Diagnostics.AddError(
			"Missing NTP Server ID",
			"Either 'dest_server' or 'id' must be provided to identify the NTP server.",
		)
		return
	}

	responseData, err := d.client.GetResource(endpoint, resid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s: %v is not present in the remote", endpoint, resid),
		)
		return
	}

	if len(responseData[endpoint].([]interface{})) < 1 {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("No %s resource found in the remote for the given query.", endpoint),
		)
		return
	}

	data.Id = types.StringValue(resid)

	getResponseData := responseData[endpoint].([]interface{})[0].(map[string]interface{})

	snmpTrapSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
