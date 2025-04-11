package sdx_license

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func sdx_license(ctx context.Context, d licenseFileData, sessionId string) (string, error) {
	tflog.Debug(ctx, "In sdx_license Method of license_file Resource")

	url := fmt.Sprintf("%s/nitro/v1/config/sdx_license", d.Host)
	method := "POST"

	payload := strings.NewReader(`{"params":{"action":"start"},"sdx_license":{"sync_operation": false}}`)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Proxy: http.ProxyFromEnvironment,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Cookie", fmt.Sprintf("NITRO_AUTH_TOKEN=%s", sessionId))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Unmarshal the JSON response into a map
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "", err
	}

	// Extract the act_id
	actID, ok := response["sdx_license"].([]interface{})[0].(map[string]interface{})["act_id"].(string)
	if !ok {
		fmt.Println("Error extracting act_id from response")
		return "", fmt.Errorf("act_id not found in response")
	}

	return actID, nil
}
