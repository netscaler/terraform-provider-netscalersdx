package service

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/go-hclog"
	tftypes "github.com/hashicorp/terraform-plugin-framework/types"
)

// HTTP Headers to be masked and not shown in logs
var headersToBeMasked = []string{"X-NITRO-USER", "X-NITRO-PASS", "Set-Cookie"}

// NitroRequestParams is a struct to hold the parameters for a Nitro request
type NitroRequestParams struct {
	ResourcePath       string
	Method             string
	Headers            map[string]string
	Resource           string
	ResourceData       interface{}
	SuccessStatusCodes []int
	ActionParams       string
}

// NitroParams encapsulates options to create a NitroClient
type NitroParams struct {
	Host          tftypes.String `tfsdk:"host"`
	Username      tftypes.String `tfsdk:"username"`
	Password      tftypes.String `tfsdk:"password"`
	SslVerify     tftypes.Bool   `tfsdk:"ssl_verify"`
	RootCAPath    tftypes.String `tfsdk:"root_ca_path"`
	ServerName    tftypes.String `tfsdk:"server_name"`
	Headers       tftypes.Map    `tfsdk:"headers"`
	LogLevel      tftypes.String `tfsdk:"log_level"`
	JSONLogFormat tftypes.Bool   `tfsdk:"json_log_format"`
}

type NitroParamsapi struct {
	Host          string
	Username      string
	Password      string
	SslVerify     bool
	RootCAPath    string
	ServerName    string
	Headers       map[string]string
	LogLevel      string
	JSONLogFormat bool
}

// NitroClient has methods to configure the NetScaler
// It abstracts the REST operations of the NITRO API
type NitroClient struct {
	host      string
	username  string
	password  string
	sslVerify bool
	client    *http.Client
	headers   map[string]string
	logger    hclog.Logger
}

// Getter methods for private fields
func (c *NitroClient) Host() string {
	return c.host
}

func (c *NitroClient) Username() string {
	return c.username
}

func (c *NitroClient) Password() string {
	return c.password
}

// NewNitroClientFromParams returns a usable NitroClient. Does not check validity of supplied parameters
func NewNitroClientFromParams(params NitroParamsapi) (*NitroClient, error) {
	u, err := url.Parse(params.Host)
	if err != nil {
		return nil, fmt.Errorf("Supplied URL %s is not a URL", params.Host)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, fmt.Errorf("Supplied Host %s does not have a HTTP/HTTPS scheme", params.Host)
	}
	c := new(NitroClient)
	c.host = params.Host
	c.username = params.Username
	c.password = params.Password
	c.sslVerify = params.SslVerify
	log.Println("NewNitroClientFromParams", "host", c.host, "username", c.username, "sslVerify", c.sslVerify)
	if params.SslVerify {
		if len(params.RootCAPath) > 0 {
			caCert, err := os.ReadFile(params.RootCAPath)
			if err != nil {
				return nil, fmt.Errorf("Unable to read certificate file: %v", err)
			}
			caCertPool := x509.NewCertPool()
			if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
				return nil, fmt.Errorf("could not Append CA certificate")
			}
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:    caCertPool,
					ServerName: params.ServerName,
				},
				Proxy: http.ProxyFromEnvironment,
			}
			c.client = &http.Client{Transport: tr}

		} else {
			c.client = &http.Client{}
		}
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Proxy: http.ProxyFromEnvironment,
		}
		c.client = &http.Client{Transport: tr}
	}

	c.logger = hclog.New(&hclog.LoggerOptions{
		Name:            "netscalersdx-client",
		Level:           hclog.LevelFromString(params.LogLevel),
		Color:           hclog.AutoColor,
		JSONFormat:      params.JSONLogFormat,
		IncludeLocation: true,
	})
	return c, nil
}

func maskHeaders(headers http.Header) http.Header {
	maskedHeaders := make(http.Header, len(headers))
	for k, v := range headers {
		upperKey := strings.ToUpper(k)
		if Contains(headersToBeMasked, upperKey) {
			maskedHeaders[k] = []string{"*********"}
		} else {
			maskedHeaders[k] = v
		}
	}
	return maskedHeaders
}

// MakeNitroRequest makes a API request to the NetScaler
func (c *NitroClient) MakeNitroRequest(n NitroRequestParams) ([]byte, error) {
	var buff []byte
	var err error

	if n.Method == "POST" || n.Method == "PUT" {
		payload := map[string]interface{}{n.Resource: n.ResourceData}
		buff, err = JSONMarshal(payload)
		if err != nil {
			return nil, err
		}
		if n.Resource != "login" {
			log.Println("MakeNitroRequest payload", toJSONIndent(payload)) // print json converted payload
		}
	} else if n.Method == "GET" || n.Method == "DELETE" {
		buff = []byte{}
	}

	urlstr := fmt.Sprintf("%s/%s", c.host, n.ResourcePath)

	req, err := http.NewRequest(n.Method, urlstr, bytes.NewBuffer(buff))
	if err != nil {
		return nil, err
	}

	// Standard headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	if n.Resource != "login" {
		req.Header.Set("X-NITRO-USER", c.username)
		req.Header.Set("X-NITRO-PASS", c.password)
	}
	// User defined headers may overwrite previous headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}
	log.Println("MakeNitroRequest HTTP method", "method", n.Method, "url", urlstr, "headers", maskHeaders(req.Header))

	resp, err := c.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	var body []byte

	if statusCodeSuccess(n.SuccessStatusCodes, resp.StatusCode) {
		body, _ = io.ReadAll(resp.Body)
		log.Println("MakeNitroRequest response", n.Method, "url:", urlstr, "status:", resp.StatusCode)
		return body, nil
	}
	log.Println("MakeNitroRequest resopnse", n.Method, "url:", urlstr, "status:", resp.StatusCode)
	body, _ = io.ReadAll(resp.Body)
	return []byte{}, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
}

func statusCodeSuccess(slice []int, val int) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}

// JSONMarshal https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225#28596225
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// WaitForActivityCompletion waits for the activity to complete
func (c *NitroClient) WaitForActivityCompletion(activityID string) error {
	for {
		time.Sleep(time.Second * 5)

		returnData, err := c.GetResource("activity_status", activityID)
		if err != nil {
			return err
		}
		activityStatus := returnData["activity_status"].([]interface{})
		log.Println("Activity Status", toJSONIndent(activityStatus))

		// check for "is_last" key in activityStatus array and if it is true, then check for "status" key. And if the value of "status" key is "Completed" or "Failed" then return the activityStatus
		for _, activity := range activityStatus {
			if activity.(map[string]interface{})["is_last"].(string) == "true" {
				if activity.(map[string]interface{})["status"].(string) == "Completed" {
					return nil
				}
				if activity.(map[string]interface{})["status"].(string) == "Failed" {
					return fmt.Errorf(activity.(map[string]interface{})["status"].(string) + ": " + activity.(map[string]interface{})["message"].(string))
				}
			}
		}
	}
}

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.EqualFold(item, val) {
			return true
		}
	}
	return false
}

func toJSONIndent(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "    ")
	return string(b)
}

// GetResource returns a resource
func (c *NitroClient) GetResource(resource string, resourceID string) (map[string]interface{}, error) {
	log.Println("GetResource method:", resource, resourceID)
	var returnData map[string]interface{}

	var resourcePath string

	resourcePath = fmt.Sprintf("nitro/v1/config/%s/%s", resource, resourceID)

	n := NitroRequestParams{
		Resource:           resource,
		ResourcePath:       resourcePath,
		Method:             "GET",
		SuccessStatusCodes: []int{200},
	}

	body, err := c.MakeNitroRequest(n)
	if err != nil {
		return returnData, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &returnData)
		if err != nil {
			return returnData, err
		}
		log.Printf("GetResource response %v", toJSONIndent(returnData))
	}
	return returnData, nil
}

// GetAllResource returns all resources
func (c *NitroClient) GetAllResource(resource string) (map[string]interface{}, error) {
	log.Println("GetAllResource method:", resource)
	var returnData map[string]interface{}

	var resourcePath string
	resourcePath = fmt.Sprintf("nitro/v1/config/%s", resource)

	n := NitroRequestParams{
		Resource:           resource,
		ResourcePath:       resourcePath,
		Method:             "GET",
		SuccessStatusCodes: []int{200},
	}

	body, err := c.MakeNitroRequest(n)
	if err != nil {
		return returnData, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &returnData)
		if err != nil {
			return returnData, err
		}
		log.Printf("GetAllResource response %v", toJSONIndent(returnData))
	}
	return returnData, nil
}

// AddResource adds a resource
func (c *NitroClient) AddResource(resource string, resourceData interface{}) (map[string]interface{}, error) {
	// For security reasons, we don't want to print the resourceData for login resource
	if resource != "login" {
		log.Println("AddResource method:", resource, resourceData)
	}
	var returnData map[string]interface{}

	var resourcePath string
	resourcePath = fmt.Sprintf("nitro/v1/config/%s", resource)

	n := NitroRequestParams{
		Resource:           resource,
		ResourcePath:       resourcePath,
		ResourceData:       resourceData,
		Method:             "POST",
		SuccessStatusCodes: []int{200, 201, 202},
	}

	body, err := c.MakeNitroRequest(n)
	if err != nil {
		return returnData, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &returnData)
		if err != nil {
			return returnData, err
		}
		if resource != "login" {
			log.Printf("AddResource response %v", toJSONIndent(returnData))
		}
	}
	return returnData, nil
}

// AddResourceWithActionParams adds a resource with action params
func (c *NitroClient) AddResourceWithActionParams(resource string, resourceData interface{}, actionParam string, resourceID string) (map[string]interface{}, error) {
	log.Println("AddResourceWithActionParams method:", resource, resourceData, actionParam)
	var returnData map[string]interface{}

	var resourcePath string
	if resourceID == "" {
		resourcePath = fmt.Sprintf("nitro/v1/config/%s?action=%s", resource, actionParam)
	} else {
		resourcePath = fmt.Sprintf("nitro/v1/config/%s/%s?action=%s", resource, resourceID, actionParam)
	}

	n := NitroRequestParams{
		Resource:           resource,
		ResourcePath:       resourcePath,
		ResourceData:       resourceData,
		ActionParams:       actionParam,
		Method:             "POST",
		SuccessStatusCodes: []int{200, 201},
	}

	body, err := c.MakeNitroRequest(n)
	if err != nil {
		return returnData, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &returnData)
		if err != nil {
			return returnData, err
		}
		log.Printf("AddResourceWithActionParams response %v", toJSONIndent(returnData))
	}
	return returnData, nil
}

// UpdateResource updates a resource
func (c *NitroClient) UpdateResource(resource string, resourceData interface{}, resourceID string) (map[string]interface{}, error) {
	log.Println("UpdateResource method:", resource, resourceData, resourceID)
	var returnData map[string]interface{}

	var resourcePath string

	if resourceID == "" {
		resourcePath = fmt.Sprintf("nitro/v1/config/%s", resource)
	} else {
		resourcePath = fmt.Sprintf("nitro/v1/config/%s/%s", resource, resourceID)
	}

	n := NitroRequestParams{
		Resource:           resource,
		ResourcePath:       resourcePath,
		ResourceData:       resourceData,
		Method:             "PUT",
		SuccessStatusCodes: []int{200, 201, 202},
	}

	body, err := c.MakeNitroRequest(n)
	if err != nil {
		return returnData, err
	}

	if len(body) == 0 {
		err = json.Unmarshal(body, &returnData)
		if err != nil {
			return returnData, err
		}
		log.Printf("UpdateResource response %v", toJSONIndent(returnData))
	}
	return returnData, nil
}

// DeleteResource deletes a resource
func (c *NitroClient) DeleteResource(resource string, resourceID string) (map[string]interface{}, error) {
	log.Println("DeleteResource method:", resource, resourceID)
	var returnData map[string]interface{}

	var resourcePath string
	resourcePath = fmt.Sprintf("nitro/v1/config/%s/%s", resource, resourceID)

	n := NitroRequestParams{
		Resource:           resource,
		ResourcePath:       resourcePath,
		Method:             "DELETE",
		SuccessStatusCodes: []int{200, 202, 204},
	}

	body, err := c.MakeNitroRequest(n)
	if err != nil {
		return returnData, err
	}

	if len(body) > 0 {
		err = json.Unmarshal(body, &returnData)
		if err != nil {
			return returnData, err
		}
		log.Printf("DeleteResource response %v", toJSONIndent(returnData))
	}
	return returnData, nil
}
