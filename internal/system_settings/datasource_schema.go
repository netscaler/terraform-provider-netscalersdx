package system_settings

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func systemSettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to retrieve information about a system settings.",
		Attributes: map[string]schema.Attribute{
			"authorize_deviceapiproxy": schema.BoolAttribute{
				Computed:            true,
				Description:         "Authorize the DeviceAPIProxy request.",
				MarkdownDescription: "Authorize the DeviceAPIProxy request.",
			},
			"basicauth": schema.BoolAttribute{
				Computed:            true,
				Description:         "Allow Basic Authentication Protocol.",
				MarkdownDescription: "Allow Basic Authentication Protocol.",
			},
			"disable_agent_old_password_input": schema.BoolAttribute{
				Computed:            true,
				Description:         "Disable old password input requirement while changing ADM agent password.",
				MarkdownDescription: "Disable old password input requirement while changing ADM agent password.",
			},
			"disk_utilization_threshold": schema.Int64Attribute{
				Computed:            true,
				Description:         "Disk utilization threshold after which data processing it stopped.",
				MarkdownDescription: "Disk utilization threshold after which data processing it stopped.",
			},
			"enable_apiproxy_credentials": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable API Proxy Credentials.",
				MarkdownDescription: "Enable API Proxy Credentials.",
			},
			"enable_certificate_download": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable Certificate Download.",
				MarkdownDescription: "Enable Certificate Download.",
			},
			"enable_cuxip": schema.BoolAttribute{
				Computed:            true,
				Description:         "Used to enable/disable CUXIP(Customer User Experience Improvement Program).",
				MarkdownDescription: "Used to enable/disable CUXIP(Customer User Experience Improvement Program).",
			},
			"enable_delete_interface_on_adc": schema.BoolAttribute{
				Computed:            true,
				Description:         "Flag to enable/disable deleting interface from ADCs on SDX.",
				MarkdownDescription: "Flag to enable/disable deleting interface from ADCs on SDX.",
			},
			"enable_nsrecover_login": schema.BoolAttribute{
				Computed:            true,
				Description:         "This setting enalbes nsrecover login for SVM.",
				MarkdownDescription: "This setting enalbes nsrecover login for SVM.",
			},
			"enable_session_timeout": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enables session timeout feature.",
				MarkdownDescription: "Enables session timeout feature.",
			},
			"enable_shell_access": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable Shell access for non-nsroot User(s).",
				MarkdownDescription: "Enable Shell access for non-nsroot User(s).",
			},
			"is_metering_enabled": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable Metering for NetScaler VPX on SDX.",
				MarkdownDescription: "Enable Metering for NetScaler VPX on SDX.",
			},
			"keep_adc_image_count": schema.Int64Attribute{
				Computed:            true,
				Description:         "Count for number of NetScaler images to be saved in Agent.",
				MarkdownDescription: "Count for number of NetScaler images to be saved in Agent.",
			},
			"keep_alive_ping_interval": schema.Int64Attribute{
				Computed:            true,
				Description:         "Agent web socket keep alive ping interval for the system.",
				MarkdownDescription: "Agent web socket keep alive ping interval for the system.",
			},
			"prompt_creds_for_stylebooks": schema.BoolAttribute{
				Computed:            true,
				Description:         "Prompt Credentials for Stylebooks.",
				MarkdownDescription: "Prompt Credentials for Stylebooks.",
			},
			"secure_access_only": schema.BoolAttribute{
				Computed:            true,
				Description:         "Secure Access only.",
				MarkdownDescription: "Secure Access only.",
			},
			"session_timeout": schema.Int64Attribute{
				Computed:            true,
				Description:         "Session timeout for the system.",
				MarkdownDescription: "Session timeout for the system.",
			},
			"session_timeout_unit": schema.StringAttribute{
				Computed:            true,
				Description:         "Session timeout unit for the system.",
				MarkdownDescription: "Session timeout unit for the system.",
			},
			"svm_ns_comm": schema.StringAttribute{
				Computed:            true,
				Description:         "Communication with Instances. Minimum length =  1 Maximum length =  10",
				MarkdownDescription: "Communication with Instances. Minimum length =  1 Maximum length =  10",
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The ID of this data source. It is the unique randomstring.",
			},
		},
	}
}
