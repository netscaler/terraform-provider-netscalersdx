package ns

import (
	"context"
	"errors"
	"fmt"

	"terraform-provider-netscalersdx/internal/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = &vpxDataSource{}

func VpxDataSource() datasource.DataSource {
	return &vpxDataSource{}
}

type vpxDataSource struct {
	client *service.NitroClient
}

type vpxDataSourceModel struct {
	IpAddress types.String `tfsdk:"ip_address"`
	Id        types.String `tfsdk:"id"`
}

func (d *vpxDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpx"
}

func (d *vpxDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Get a VPX device ID by IP address",
		Attributes: map[string]schema.Attribute{
			"ip_address": schema.StringAttribute{
				Required:    true,
				Description: "IP Address for this VPX device",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "ID of the resource",
			},
		},
	}
}

func (d *vpxDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *vpxDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "In Read Method of vpxDataSource")
	var data vpxDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceID, err := getVpxID(d.client, data.IpAddress.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to get Managed Device ID",
			fmt.Sprintf("Unable to get Managed Device ID: %s", err.Error()),
		)
		return
	}
	data.Id = types.StringValue(resourceID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func getVpxID(c *service.NitroClient, ipAddress string) (string, error) {
	endpoint := "ns"
	returnData, err := c.GetAllResource(endpoint)
	if err != nil {
		return "", err
	}

	for _, v := range returnData[endpoint].([]interface{}) {
		if v.(map[string]interface{})["ip_address"].(string) == ipAddress {
			return v.(map[string]interface{})["id"].(string), nil
		}
	}
	return "", errors.New("Failed to find VPX instance ID with IP: " + ipAddress)
}
