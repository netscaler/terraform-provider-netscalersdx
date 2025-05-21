package snmp_mib

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func snmpMibResourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Configuration for SNMP MIB Information resource.",
		Attributes: map[string]schema.Attribute{
			"contact": schema.StringAttribute{
				Required:            true,
				Description:         "Name of the administrator for appliance.. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Name of the administrator for appliance.. Minimum length =  1 Maximum length =  127",
			},
			"custom_id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Custom identification number for appliance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Custom identification number for appliance. Minimum length =  1 Maximum length =  127",
			},
			"location": schema.StringAttribute{
				Required:            true,
				Description:         "Physical location of appliance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Physical location of appliance. Minimum length =  1 Maximum length =  127",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Name for appliance. Minimum length =  1 Maximum length =  127",
				MarkdownDescription: "Name for appliance. Minimum length =  1 Maximum length =  127",
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
		},
	}
}

type snmpMibModel struct {
	Contact  types.String `tfsdk:"contact"`
	CustomId types.String `tfsdk:"custom_id"`
	Location types.String `tfsdk:"location"`
	Name     types.String `tfsdk:"name"`
	Id       types.String `tfsdk:"id"`
}

func snmpMibGetThePayloadFromtheConfig(ctx context.Context, data *snmpMibModel) snmpMibReq {
	tflog.Debug(ctx, "In snmpMibGetThePayloadFromtheConfig Function")
	snmpMibReqPayload := snmpMibReq{
		Contact:  data.Contact.ValueString(),
		CustomId: data.CustomId.ValueString(),
		Location: data.Location.ValueString(),
		Name:     data.Name.ValueString(),
	}
	return snmpMibReqPayload
}
func snmpMibSetAttrFromGet(ctx context.Context, data *snmpMibModel, getResponseData map[string]interface{}) *snmpMibModel {
	tflog.Debug(ctx, "In snmpMibSetAttrFromGet Function")

	data.Contact = types.StringValue(getResponseData["contact"].(string))
	data.CustomId = types.StringValue(getResponseData["custom_id"].(string))
	data.Location = types.StringValue(getResponseData["location"].(string))
	data.Name = types.StringValue(getResponseData["name"].(string))
	
	return data
}

type snmpMibReq struct {
	Contact  string `json:"contact,omitempty"`
	CustomId string `json:"custom_id,omitempty"`
	Location string `json:"location,omitempty"`
	Name     string `json:"name,omitempty"`
}
