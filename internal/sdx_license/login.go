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

func getSessionID(ctx context.Context, d licenseFileData) (string, error) {
	tflog.Debug(ctx, "In getSessionID Method of license_file Resource")

	url := fmt.Sprintf("%s/nitro/v2/config/login", d.Host)
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"login":{"username":"%s","password":"%s"}}`, d.Username, d.Password))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the JSON response into a map
	var loginResponse map[string]interface{}
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		return "", err
	}

	// Access the sessionid field
	if login, ok := loginResponse["login"].([]interface{}); ok && len(login) > 0 {
		if sessionData, ok := login[0].(map[string]interface{}); ok {
			if sessionID, ok := sessionData["sessionid"].(string); ok {
				return sessionID, nil
			} else {
				return "", fmt.Errorf("Session ID not found")
			}
		} else {
			return "", fmt.Errorf("Invalid session data format")
		}
	} else {
		return "", fmt.Errorf("No login data found")
	}
}
