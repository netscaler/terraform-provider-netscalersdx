package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// PEMEntitlementMapping maps PEM codes to entitlement names
var PEMEntitlementMapping = map[string]string{
	"CNS_8905_SERVER":   "MPX 8905",
	"CNS_8910_SERVER":   "MPX 8910",
	"CNS_8920_SERVER":   "MPX 8920",
	"CNS_8930_SERVER":   "MPX 8930",
	"CNS_9110_SERVER":   "MPX 9110",
	"CNS_9120_SERVER":   "MPX 9120",
	"CNS_9130_SERVER":   "MPX 9130",
	"CNS_5901_SERVER":   "MPX 5901",
	"CNS_5905_SERVER":   "MPX 5905",
	"CNS_5910_SERVER":   "MPX 5910",
	"CNS_14020_SERVER":  "FIPS MPX 14020",
	"CNS_14030_SERVER":  "FIPS MPX 14030",
	"CNS_14060_SERVER":  "FIPS MPX 14060",
	"CNS_14080_SERVER":  "FIPS MPX 14080",
	"CNS_14500_SERVER":  "FIPS MPX 14500",
	"CNS_16030_SERVER":  "MPX 16030",
	"CNS_16040_SERVER":  "MPX 16040",
	"CNS_16060_SERVER":  "MPX 16060",
	"CNS_16120_SERVER":  "MPX 16120",
	"CNS_16200_SERVER":  "MPX 16200",
	"CNS_15120_SERVER":  "MPX 15120 / 15120-50G",
	"CNS_26200_SERVER":  "MPX 26200 / 26200-50S / 26200-100G Premium",
	"CNS_9205_SERVER":   "MPX 9205",
	"CNS_9210_SERVER":   "MPX 9210",
	"CNS_9220_SERVER":   "MPX 9220",
	"CNS_9240_SERVER":   "MPX 9240",
	"CNS_9260_SERVER":   "MPX 9260",
	"CNS_9280_SERVER":   "MPX 9280",
	"CNS_9295_SERVER":   "MPX 9295",
	"CNS_9299_SERVER":   "MPX 9299",
	"CNS_17020_SERVER":  "MPX 17020",
	"CNS_17050_SERVER":  "MPX 17050",
	"CNS_17100_SERVER":  "MPX 17100",
	"CNS_17150_SERVER":  "MPX 17150",
	"CNS_17200_SERVER":  "MPX 17200",
	"CNS_17250_SERVER":  "MPX 17250",
	"CNS_17300_SERVER":  "MPX 17300",
	"CNS_17400_SERVER":  "MPX 17400",
	"CNS_17500_SERVER":  "MPX 17500",
	"CNS_M8920_SERVER":  "SDX 8920",
	"CNS_M8930_SERVER":  "SDX 8930",
	"CNS_M9120_SERVER":  "SDX 9120",
	"CNS_M9130_SERVER":  "SDX 9130",
	"CNS_M9140_SERVER":  "SDX 9140",
	"CNS_M9160_SERVER":  "SDX 9160",
	"CNS_M9180_SERVER":  "SDX 9180",
	"CNS_M9195_SERVER":  "SDX 9195",
	"CNS_M9220_SERVER":  "SDX 9220",
	"CNS_M9280_SERVER":  "SDX 9280",
	"CNS_M9295_SERVER":  "SDX 9299",
	"CNS_M14020_SERVER": "SDX 14020",
	"CNS_M14030_SERVER": "SDX 14030",
	"CNS_M14040_SERVER": "SDX 14040",
	"CNS_M14060_SERVER": "SDX 14060",
	"CNS_M14080_SERVER": "SDX 14080",
	"CNS_M14100_SERVER": "SDX 14100",
	"CNS_M16030_SERVER": "SDX 16030",
	"CNS_M16040_SERVER": "SDX 16040",
	"CNS_M16060_SERVER": "SDX 16060",
	"CNS_M16080_SERVER": "SDX 16080",
	"CNS_M16100_SERVER": "SDX 16100",
	"CNS_M16120_SERVER": "SDX 16120",
	"CNS_M16160_SERVER": "SDX 16160",
	"CNS_M16200_SERVER": "SDX 16200",
	"CNS_M15020_SERVER": "SDX 15020 / 15020-50G",
	"CNS_M15030_SERVER": "SDX 15030 / 15030-50G",
	"CNS_M15040_SERVER": "SDX 15040 / 15040-50G",
	"CNS_M15060_SERVER": "SDX 15060 / 15060-50G",
	"CNS_M15080_SERVER": "SDX 15080 / 15080-50G",
	"CNS_M15100_SERVER": "SDX 15100 / 15100-50G",
	"CNS_M15120_SERVER": "SDX 15120 / 15120-50G",
	"CNS_M26100_SERVER": "SDX 26100 / 26100-50S / 26100-100G",
	"CNS_M26160_SERVER": "SDX 26160 / 26160-50S / 26160-100G",
	"CNS_M26200_SERVER": "SDX 26200 / 26200-50S / 26200-100G",
	"CNS_M17020_SERVER": "SDX 17020",
	"CNS_M17050_SERVER": "SDX 17050",
	"CNS_M17100_SERVER": "SDX 17100",
	"CNS_M17300_SERVER": "SDX 17300",
	"CNS_M17400_SERVER": "SDX 17400",
	"CNS_M17500_SERVER": "SDX 17500",
	"CNS_V25000_SERVER": "VPX 25000",
	"CNS_V10000_SERVER": "VPX 10000",
	"CNS_V5000_SERVER":  "VPX 5000",
	"CNS_V3000_SERVER":  "VPX 3000",
	"CNS_V1000_SERVER":  "VPX 1000",
	"CNS_V200_SERVER":   "VPX 200",
	"CNS_V25_SERVER":    "VPX 25",
	"CNS_V10_SERVER":    "VPX 10",
}

// LASTokenGenerator handles LAS token generation
type LASTokenGenerator struct {
	Endpoint      string
	LSGUID        string
	CCID          string
	SecretClient  string
	SecretPwd     string
	BaseURL       string
	CCTokenURL    string
	BearerCache   string
	BearerToken   string
	HTTPClient    *http.Client
	InsecureHTTPS bool
}

// NewLASTokenGenerator creates a new LAS token generator
func NewLASTokenGenerator(endpoint, lsguid, ccid, client, password, baseURL, ccTokenURL string) *LASTokenGenerator {
	return &LASTokenGenerator{
		Endpoint:      endpoint,
		LSGUID:        lsguid,
		CCID:          ccid,
		SecretClient:  client,
		SecretPwd:     password,
		BaseURL:       baseURL,
		CCTokenURL:    ccTokenURL,
		BearerCache:   "/tmp/las_bearer_cache",
		InsecureHTTPS: true,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// RunCurlHTTPSFallback tries HTTPS first, falls back to HTTP
func RunCurlHTTPSFallback(ctx context.Context, url, method string, auth *BasicAuth, body []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Try HTTPS first
	httpsURL := strings.Replace(url, "http://", "https://", 1)
	resp, err := makeHTTPRequest(ctx, client, httpsURL, method, auth, body, headers)
	if err == nil {
		return resp, nil
	}

	tflog.Debug(ctx, "HTTPS request failed, falling back to HTTP", map[string]interface{}{"error": err.Error()})

	// Fallback to HTTP
	httpURL := strings.Replace(httpsURL, "https://", "http://", 1)
	return makeHTTPRequest(ctx, client, httpURL, method, auth, body, headers)
}

// BasicAuth holds basic authentication credentials
type BasicAuth struct {
	Username string
	Password string
}

func makeHTTPRequest(ctx context.Context, client *http.Client, url, method string, auth *BasicAuth, body []byte, headers map[string]string) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	// Log request details
	tflog.Debug(ctx, "HTTP Request", map[string]interface{}{
		"method": method,
		"url":    url,
	})
	if body != nil && len(body) > 0 {
		tflog.Debug(ctx, "Request Body", map[string]interface{}{
			"body": string(body),
		})
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if auth != nil {
		req.SetBasicAuth(auth.Username, auth.Password)
		tflog.Debug(ctx, "Using Basic Auth", map[string]interface{}{"username": auth.Username})
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Log response details
	tflog.Debug(ctx, "HTTP Response", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if len(respBody) > 0 {
		tflog.Debug(ctx, "Response Body", map[string]interface{}{
			"body": string(respBody),
		})
	}

	if resp.StatusCode >= 400 {
		return respBody, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// GetOfflineRequestPackage generates and retrieves offline activation request package
func GetOfflineRequestPackage(ctx context.Context, product, ip, hostname, username, password string, useHostname bool) (string, []byte, error) {
	auth := &BasicAuth{Username: username, Password: password}
	var url, srcDir string

	switch product {
	case "NS":
		if useHostname {
			url = fmt.Sprintf("http://%s/nitro/v1/config/nslicenseactivationdata?args=usehostname:true", ip)
		} else {
			url = fmt.Sprintf("http://%s/nitro/v1/config/nslicenseactivationdata", ip)
		}
		srcDir = "/nsconfig/license/"
	case "SDX":
		url = fmt.Sprintf("http://%s/nitro/v1/config/las_activation_request", ip)
		srcDir = "/mpsconfig/license/"
	case "ADM":
		url = fmt.Sprintf("http://%s/nitro/v1/config/lic_darksite", ip)
		srcDir = "/mpsconfig/license/"
	default:
		return "", nil, fmt.Errorf("unsupported product: %s", product)
	}

	// Make API call to generate request package
	var body []byte
	var headers map[string]string
	method := "GET"

	if useHostname && product != "NS" {
		// For SDX and ADM with useHostname, use POST with form-encoded body
		method = "POST"
		headers = map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
		if product == "SDX" {
			payload := map[string]interface{}{
				"las_activation_request": map[string]bool{
					"use_hostname": true,
				},
			}
			jsonPayload, _ := json.Marshal(payload)
			body = []byte(fmt.Sprintf("object=%s", string(jsonPayload)))
		} else if product == "ADM" {
			payload := map[string]interface{}{
				"lic_darksite": map[string]bool{
					"use_hostname": true,
				},
			}
			jsonPayload, _ := json.Marshal(payload)
			body = []byte(fmt.Sprintf("object=%s", string(jsonPayload)))
		}
	}
	// For NS with useHostname, keep GET method with usehostname:true in URL params

	respBody, err := RunCurlHTTPSFallback(ctx, url, method, auth, body, headers)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate request package: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Extract filename
	var filename string
	switch product {
	case "NS":
		if data, ok := result["nslicenseactivationdata"].(map[string]interface{}); ok {
			filename, _ = data["filename"].(string)
		}
	case "SDX":
		if data, ok := result["las_activation_request"].([]interface{}); ok && len(data) > 0 {
			if obj, ok := data[0].(map[string]interface{}); ok {
				filename, _ = obj["package_name"].(string)
			}
		}
	case "ADM":
		if data, ok := result["lic_darksite"].([]interface{}); ok && len(data) > 0 {
			if obj, ok := data[0].(map[string]interface{}); ok {
				filename, _ = obj["package_name"].(string)
			}
		}
	}

	if filename == "" {
		return "", nil, fmt.Errorf("failed to extract filename from response")
	}

	// Download the file via SCP
	remotePath := srcDir + filename
	fileContent, err := SCPDownload(ctx, ip, username, password, remotePath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to download request package: %w", err)
	}

	tflog.Info(ctx, "Generated request package", map[string]interface{}{"ip": ip, "filename": filename})
	return filename, fileContent, nil
}

// SCPDownload downloads a file via SCP
func SCPDownload(ctx context.Context, ip, username, password, remotePath string) ([]byte, error) {
	tflog.Debug(ctx, "Starting SFTP download", map[string]interface{}{
		"ip":         ip,
		"remotePath": remotePath,
	})

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				answers := make([]string, len(questions))
				for i := range answers {
					answers[i] = password
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		tflog.Error(ctx, "Failed to dial SSH", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	defer client.Close()

	// Create SFTP client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		tflog.Error(ctx, "Failed to create SFTP client", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to create SFTP client: %w", err)
	}
	defer sftpClient.Close()

	tflog.Debug(ctx, "Opening remote file via SFTP", map[string]interface{}{"path": remotePath})

	// Open the remote file
	remoteFile, err := sftpClient.Open(remotePath)
	if err != nil {
		tflog.Error(ctx, "Failed to open remote file", map[string]interface{}{
			"error": err.Error(),
			"path":  remotePath,
		})
		return nil, fmt.Errorf("failed to open remote file: %w", err)
	}
	defer remoteFile.Close()

	// Read the file content
	var buf bytes.Buffer
	bytesRead, err := io.Copy(&buf, remoteFile)
	if err != nil {
		tflog.Error(ctx, "Failed to read remote file", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to read remote file: %w", err)
	}

	tflog.Debug(ctx, "SFTP download successful", map[string]interface{}{
		"bytesDownloaded": bytesRead,
	})
	return buf.Bytes(), nil
}

// SCPUpload uploads a file via SFTP
func SCPUpload(ctx context.Context, ip, username, password, remotePath string, content []byte) error {
	tflog.Debug(ctx, "Starting SFTP upload", map[string]interface{}{
		"ip":         ip,
		"remotePath": remotePath,
		"size":       len(content),
	})

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			ssh.KeyboardInteractive(func(user, instruction string, questions []string, echos []bool) ([]string, error) {
				answers := make([]string, len(questions))
				for i := range answers {
					answers[i] = password
				}
				return answers, nil
			}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		tflog.Error(ctx, "Failed to dial SSH", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to dial: %w", err)
	}
	defer client.Close()

	// Create SFTP client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		tflog.Error(ctx, "Failed to create SFTP client", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to create SFTP client: %w", err)
	}
	defer sftpClient.Close()

	tflog.Debug(ctx, "Creating remote file via SFTP", map[string]interface{}{"path": remotePath})

	// Create the remote file
	remoteFile, err := sftpClient.Create(remotePath)
	if err != nil {
		tflog.Error(ctx, "Failed to create remote file", map[string]interface{}{
			"error": err.Error(),
			"path":  remotePath,
		})
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer remoteFile.Close()

	// Write content to the file
	bytesWritten, err := remoteFile.Write(content)
	if err != nil {
		tflog.Error(ctx, "Failed to write to remote file", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to write to remote file: %w", err)
	}

	tflog.Info(ctx, "Uploaded file via SFTP", map[string]interface{}{
		"remotePath":   remotePath,
		"bytesWritten": bytesWritten,
	})
	return nil
}

// ExtractLSGUIDFromPackage extracts LSGUID from request package
func ExtractLSGUIDFromPackage(ctx context.Context, product string, packageData []byte) (string, error) {
	var jsonFile string
	switch product {
	case "ADM":
		jsonFile = "console_offline_activation_request.json"
	case "SDX":
		jsonFile = "svm_offline_activation_request.json"
	case "NS":
		jsonFile = "ns_offline_activation_request.json"
	default:
		return "", fmt.Errorf("unsupported product: %s", product)
	}

	// Decompress and extract
	gzReader, err := gzip.NewReader(bytes.NewReader(packageData))
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read tar: %w", err)
		}

		if header.Name == jsonFile || strings.HasSuffix(header.Name, "/"+jsonFile) {
			jsonData, err := io.ReadAll(tarReader)
			if err != nil {
				return "", fmt.Errorf("failed to read JSON file: %w", err)
			}

			var result map[string]interface{}
			if err := json.Unmarshal(jsonData, &result); err != nil {
				return "", fmt.Errorf("failed to parse JSON: %w", err)
			}

			lsguid, ok := result["lsguid"].(string)
			if !ok || lsguid == "" {
				return "", fmt.Errorf("lsguid not found in JSON")
			}

			tflog.Info(ctx, "Extracted LSGUID", map[string]interface{}{"lsguid": lsguid})
			return lsguid, nil
		}
	}

	return "", fmt.Errorf("JSON file %s not found in package", jsonFile)
}

// GenerateBearerToken generates a bearer token for LAS API
func (ltg *LASTokenGenerator) GenerateBearerToken(ctx context.Context) (string, error) {
	payload := map[string]string{
		"clientId":     ltg.SecretClient,
		"clientSecret": ltg.SecretPwd,
	}
	body, _ := json.Marshal(payload)

	// Log request (sanitize sensitive data)
	tflog.Debug(ctx, "API Call: GenerateBearerToken", map[string]interface{}{
		"url":    ltg.CCTokenURL,
		"method": "POST",
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"clientId":     ltg.SecretClient,
		"clientSecret": "***REDACTED***",
	})

	req, err := http.NewRequestWithContext(ctx, "POST", ltg.CCTokenURL, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to generate bearer token: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Bearer Token Generation Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("invalid JSON response: %w", err)
	}

	token, ok := result["token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("token not found in response")
	}

	ltg.BearerToken = token
	tflog.Info(ctx, "Generated bearer token")
	tflog.Debug(ctx, "Response Payload", map[string]interface{}{
		"token": "***REDACTED***",
	})
	return token, nil
}

// GetFingerprintForLSGUID gets fingerprint for LSGUID and deregisters if exists
func (ltg *LASTokenGenerator) GetFingerprintForLSGUID(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s/support/%s/%s/listls", ltg.BaseURL, ltg.CCID, ltg.Endpoint)
	payload := map[string]string{"ver": "1.0"}
	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: GetFingerprintForLSGUID", map[string]interface{}{
		"url":    url,
		"method": "POST",
		"lsguid": ltg.LSGUID,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to list license servers: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("invalid JSON response: %w", err)
	}

	// Find and deregister if exists
	if lstlasactivatedls, ok := result["lstlasactivatedls"].([]interface{}); ok {
		for _, ls := range lstlasactivatedls {
			if lsObj, ok := ls.(map[string]interface{}); ok {
				if lsguid, _ := lsObj["lsguid"].(string); lsguid == ltg.LSGUID {
					fingerprint, _ := lsObj["lsfingerprint"].(string)
					// if fingerprint != "" {
					// 	tflog.Info(ctx, "Found existing license server, deregistering", map[string]interface{}{"fingerprint": fingerprint})
					// 	ltg.DeregisterLicense(ctx, fingerprint)
					// }
					return fingerprint, nil
				}
			}
		}
	}

	return "", nil
}

// DeregisterLicense deregisters a license by fingerprint
func (ltg *LASTokenGenerator) DeregisterLicense(ctx context.Context, fingerprint string) error {
	url := fmt.Sprintf("%s/%s/%s/deregisterls", ltg.BaseURL, ltg.CCID, ltg.Endpoint)
	payload := map[string]string{
		"lsfingerprint": fingerprint,
		"ver":           "1.0",
	}
	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: DeregisterLicense", map[string]interface{}{
		"url":         url,
		"method":      "POST",
		"fingerprint": fingerprint,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to deregister: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	if resp.StatusCode >= 400 {
		respBody, _ := io.ReadAll(resp.Body)
		tflog.Error(ctx, "Deregister Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	tflog.Info(ctx, "Deregistered license", map[string]interface{}{"fingerprint": fingerprint})
	return nil
}

// ImportOfflineActivationRequest imports the offline activation request
func (ltg *LASTokenGenerator) ImportOfflineActivationRequest(ctx context.Context, requestPackage []byte, fingerprint string) (string, error) {
	url := fmt.Sprintf("%s/support/%s/%s/importofflineactivationrequest", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	// Create multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Write file part
	part, err := writer.CreateFormFile("file", "request.tgz")
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := part.Write(requestPackage); err != nil {
		return "", fmt.Errorf("failed to write file part: %w", err)
	}

	// Write data part
	dataJSON, _ := json.Marshal(map[string]string{
		"ver":           "1.0",
		"lsfingerprint": fingerprint,
	})
	if err := writer.WriteField("data", string(dataJSON)); err != nil {
		return "", fmt.Errorf("failed to write data field: %w", err)
	}

	writer.Close()

	// Log request
	tflog.Debug(ctx, "API Call: ImportOfflineActivationRequest", map[string]interface{}{
		"url":          url,
		"method":       "POST",
		"fingerprint":  fingerprint,
		"package_size": len(requestPackage),
	})
	tflog.Debug(ctx, "Request Data Field", map[string]interface{}{
		"data": string(dataJSON),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to import request: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Import Request Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("invalid JSON response: %w", err)
	}

	importToken, ok := result["importrequesttoken"].(string)
	if !ok || importToken == "" {
		return "", fmt.Errorf("importrequesttoken not found in response")
	}

	tflog.Info(ctx, "Imported offline activation request", map[string]interface{}{"importToken": importToken})
	return importToken, nil
}

// GenerateOfflineActivation generates the offline activation
func (ltg *LASTokenGenerator) GenerateOfflineActivation(ctx context.Context, importToken, entitlementName string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/%s/generateofflineactivation", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	payload := map[string]interface{}{
		"ver":                "1.0",
		"importrequesttoken": importToken,
	}

	if ltg.Endpoint == "netscalerfixedbw" && entitlementName != "" {
		payload["entitlementname"] = entitlementName
	}

	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: GenerateOfflineActivation", map[string]interface{}{
		"url":    url,
		"method": "POST",
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate activation: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Generate Activation Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	tflog.Info(ctx, "Generated offline activation", map[string]interface{}{"result": result})
	return result, nil
}

// ExportOfflineActivationResponse exports the license blob
func (ltg *LASTokenGenerator) ExportOfflineActivationResponse(ctx context.Context, activationID, fingerprint string) ([]byte, error) {
	url := fmt.Sprintf("%s/support/%s/%s/exportofflineactivationresponse", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	payload := map[string]string{
		"ver":             "1.0",
		"lsfingerprint":   fingerprint,
		"newactivationid": activationID,
	}
	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: ExportOfflineActivationResponse", map[string]interface{}{
		"url":          url,
		"method":       "POST",
		"activationID": activationID,
		"fingerprint":  fingerprint,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to export blob: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body Size", map[string]interface{}{
		"size_bytes": len(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Export Activation Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	tflog.Info(ctx, "Exported offline activation response")
	return respBody, nil
}

// ApplyLicenseBlobADM applies license blob to ADM/SDX
func ApplyLicenseBlobADM(ctx context.Context, ip, username, password string, blobContent []byte) error {
	// Create temp filename
	filename := fmt.Sprintf("offline_token_%s_activation.blob.tgz", ip)

	// Upload blob to device
	remotePath := "/mpsconfig/license/" + filename
	if err := SCPUpload(ctx, ip, username, password, remotePath, blobContent); err != nil {
		return fmt.Errorf("failed to upload license blob: %w", err)
	}

	// Apply license - SDX API expects form-encoded data with "object" key
	auth := &BasicAuth{Username: username, Password: password}
	url := fmt.Sprintf("http://%s/nitro/v1/config/las_lic_apply", ip)

	payload := map[string]interface{}{
		"las_lic_apply": map[string]string{
			"las_token_file_name": filename,
		},
	}
	payloadJSON, _ := json.Marshal(payload)

	// Form-encode with "object" key
	formData := fmt.Sprintf("object=%s", string(payloadJSON))
	body := []byte(formData)
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	// Log request
	tflog.Debug(ctx, "API Call: ApplyLicenseBlobADM", map[string]interface{}{
		"url":      url,
		"method":   "POST",
		"ip":       ip,
		"filename": filename,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": formData,
	})

	respBody, err := RunCurlHTTPSFallback(ctx, url, "POST", auth, body, headers)
	if err != nil {
		return fmt.Errorf("failed to apply license: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	tflog.Info(ctx, "Applied license blob", map[string]interface{}{"ip": ip})
	return nil
}

// GetEntitlementNameForFixedBW gets the entitlement name for fixed bandwidth
func GetEntitlementNameForFixedBW(product, requestPEM, requestED string, isFIPS bool) (string, error) {
	baseEntString, ok := PEMEntitlementMapping[requestPEM]
	if !ok {
		return "", fmt.Errorf("PEM not found: %s", requestPEM)
	}

	// FIPS validation
	if isFIPS {
		fipsPEMs := []string{"CNS_8910_SERVER", "CNS_8920_SERVER", "CNS_9130_SERVER", "CNS_15120_SERVER",
			"CNS_V5000_SERVER", "CNS_V3000_SERVER", "CNS_V1000_SERVER", "CNS_V200_SERVER", "CNS_V25_SERVER"}
		found := false
		for _, pem := range fipsPEMs {
			if requestPEM == pem {
				found = true
				break
			}
		}
		if !found {
			return "", fmt.Errorf("FIPS not supported for PEM: %s", requestPEM)
		}

		mpxFipsPEMs := []string{"CNS_8910_SERVER", "CNS_8920_SERVER", "CNS_9130_SERVER", "CNS_15120_SERVER"}
		for _, pem := range mpxFipsPEMs {
			if requestPEM == pem && requestED != "Premium" {
				return "", fmt.Errorf("MPX FIPS only supported for Premium edition")
			}
		}

		if requestPEM == "CNS_15120_SERVER" {
			baseEntString = "FIPS MPX 15120-50G"
		} else {
			baseEntString = "FIPS " + baseEntString
		}
	}

	// Product-specific formatting
	if product == "SDX" {
		if (strings.Contains(requestPEM, "CNS_M15") && requestED == "50G") ||
			(strings.Contains(requestPEM, "CNS_M26") && (requestED == "50S" || requestED == "100G")) {
			return baseEntString + "-" + requestED + " Premium", nil
		}
		return baseEntString + " Premium", nil
	} else if requestED == "Advanced" || requestED == "Standard" || requestED == "Premium" {
		return baseEntString + " " + requestED, nil
	}

	return "", fmt.Errorf("invalid edition: %s", requestED)
}

// GetMPSVersion retrieves version and build information from SDX/ADM
func GetMPSVersion(ctx context.Context, ip, username, password string) (string, string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Try HTTPS first
	url := fmt.Sprintf("https://%s/nitro/v1/config/mps", ip)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		// Fallback to HTTP
		url = fmt.Sprintf("http://%s/nitro/v1/config/mps", ip)
		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return "", "", fmt.Errorf("failed to create HTTP request: %w", err)
		}
		req.SetBasicAuth(username, password)
		resp, err = client.Do(req)
		if err != nil {
			return "", "", fmt.Errorf("failed to get MPS version: %w", err)
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("failed to read response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	mpsData, ok := result["mps"].([]interface{})
	if !ok || len(mpsData) == 0 {
		return "", "", fmt.Errorf("invalid mps response structure")
	}

	mpsObj, ok := mpsData[0].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid mps object structure")
	}

	buildNumber, ok := mpsObj["product_build_number"].(string)
	if !ok || buildNumber == "" {
		return "", "", fmt.Errorf("product_build_number not found")
	}

	// Parse build_number format: "NetScaler SDX 14.1: Build 4.51"
	// Extract release (e.g., "14.1")
	releaseRegex := regexp.MustCompile(`(\d+\.\d+):`)
	releaseMatch := releaseRegex.FindStringSubmatch(buildNumber)
	if len(releaseMatch) < 2 {
		return "", "", fmt.Errorf("failed to parse release from: %s", buildNumber)
	}
	release := releaseMatch[1]

	// Extract build (e.g., "4.51")
	buildRegex := regexp.MustCompile(`Build\s+(\d+\.\d+)`)
	buildMatch := buildRegex.FindStringSubmatch(buildNumber)
	if len(buildMatch) < 2 {
		return "", "", fmt.Errorf("failed to parse build from: %s", buildNumber)
	}
	build := buildMatch[1]

	tflog.Debug(ctx, "Retrieved MPS version", map[string]interface{}{
		"ip":      ip,
		"release": release,
		"build":   build,
	})

	return release, build, nil
}

// DetermineNewAPINeeded checks if new API is needed based on version
func DetermineNewAPINeeded(product, version, build string) bool {
	parts := strings.Split(build, ".")
	if len(parts) < 2 {
		return false
	}
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])

	// Version 14.1 checks
	if version == "14.1" {
		if major > 68 {
			return true
		}
		if major == 68 && minor >= 3 {
			return true
		}
		if major == 60 && minor >= 55 {
			return true
		}
		if major == 66 && minor >= 32 {
			return true
		}
	}

	// Version 13.1 checks for NS/SDX
	if version == "13.1" {
		if product == "NS" || product == "SDX" {
			if major > 62 {
				return true
			}
			if major == 62 && minor >= 6 {
				return true
			}
			if major == 61 && minor >= 26 {
				return true
			}
		}
		// FIPS
		if major == 37 && minor >= 256 {
			return true
		}
	}

	return false
}
