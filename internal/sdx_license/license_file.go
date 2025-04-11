package sdx_license

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func licenseFileMultipart(ctx context.Context, d licenseFileData, sessionId string) error {
	tflog.Debug(ctx, "In licenseFileMultipart Method of license_file Resource")

	url := fmt.Sprintf("%s/nitro/v1/upload/license_file", d.Host)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(d.FileName)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("license_files", filepath.Base(d.FileName))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		tflog.Debug(ctx, fmt.Sprint(errFile1))
		return errFile1
	}
	err := writer.Close()
	if err != nil {
		tflog.Debug(ctx, fmt.Sprint(errFile1))
		return err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Proxy: http.ProxyFromEnvironment,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		tflog.Debug(ctx, fmt.Sprint(errFile1))
		return err
	}
	req.Header.Add("Cookie", fmt.Sprintf("NITRO_AUTH_TOKEN=%s", sessionId))

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprint(errFile1))
		return err
	}
	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprint(errFile1))
		return err
	}
	return nil
}
