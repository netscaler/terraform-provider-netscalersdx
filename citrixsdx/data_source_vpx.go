package citrixsdx

import (
	"context"
	"errors"
	"log"

	"terraform-provider-citrixsdx/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVpx() *schema.Resource {
	return &schema.Resource{
		Description: "Get a VPX device ID by IP address",
		ReadContext: dataSourceVpxRead,
		Schema: map[string]*schema.Schema{
			"ip_address": {
				Description: "IP Address for this VPX device",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceVpxRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("In dataSourceVpxRead")
	var diags diag.Diagnostics
	c := m.(*service.NitroClient)

	resourceID, err := getVpxID(c, d.Get("ip_address").(string))

	if err != nil {
		return diag.Errorf("unable to get Managed Device ID: %s", err.Error())
	}
	d.SetId(resourceID)
	d.Set("ip_address", d.Get("ip_address").(string))

	return diags
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
