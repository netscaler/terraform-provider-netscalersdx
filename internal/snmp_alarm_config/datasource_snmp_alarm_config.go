package snmp_alarm_config

import (
	"context"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*snmpAlarmConfigDataSource)(nil)

func SnmpAlarmConfigDataSource() datasource.DataSource {
	return &snmpAlarmConfigDataSource{}
}

type snmpAlarmConfigDataSource struct {
	client *service.NitroClient
}

func (d *snmpAlarmConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmp_alarm_config"
}

func (d *snmpAlarmConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *snmpAlarmConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = snmpAlarmConfigDataSourceSchema()
}

func (d *snmpAlarmConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data snmpAlarmConfigModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	endpoint := "snmp_alarm_config"

	var resid string
	if !data.Name.IsNull() {
		resid = data.Name.ValueString()
	} else if !data.Id.IsNull() {
		resid = data.Id.ValueString()
	}

	if resid == "" {
		resp.Diagnostics.AddError(
			"Missing snmp alarm config ID",
			"Either 'name' or 'id' must be provided to identify the snmp alarm config.",
		)
		return
	}

	responseData, err := d.client.GetResource(endpoint, resid)
	if err != nil {
		resp.Diagnostics.AddError(
			"Resource Not Found",
			fmt.Sprintf("%s: %v is not present in the remote", endpoint, data.Name.ValueString()),
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

	snmpAlarmConfigSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
