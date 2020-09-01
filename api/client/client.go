package client

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Client holds all of the information required to connect to a server
type Client struct {
	cloudview_api string
	assetview_api string
	username      string
	password      string
	httpClient    *http.Client
}

// NewClient returns a new client configured to communicate with API
func NewClient(cloudview_api string, assetview_api string, username string, password string) *Client {
	return &Client{
		cloudview_api: cloudview_api,
		assetview_api: assetview_api,
		username:      username,
		password:      password,
		httpClient:    &http.Client{},
	}
}

func (c *Client) requestPath(apiUrl string, path string) string {
	return fmt.Sprintf("%s%s", apiUrl, path)
}

func (c *Client) httpRequestCloudView(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(c.cloudview_api, path), &body)
	if err != nil {
		return nil, err
	}
	creds := base64.StdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
	req.Header.Add("Authorization", "Basic "+creds)

	switch method {
	case "GET":
	case "DELETE":
		req.Header.Add("Content-Type", "application/json")
	case "PUT":
		req.Header.Add("Content-Type", "application/json")
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case
		http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent:
	default:
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Got an unexpected HTTP status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("Got an unexpected status HTTP code: %v - %s", resp.StatusCode, buf.String())
	}
	fmt.Print("BODY:")
	fmt.Print(resp.Body)
	return resp.Body, nil
}

func (c *Client) httpRequestAssetView(path, method string, body bytes.Buffer) (*AssetViewServiceResponse, error) {
	req, err := http.NewRequest(method, c.requestPath(c.assetview_api, path), &body)
	if err != nil {
		return nil, err
	}
	creds := base64.StdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
	req.Header.Add("Authorization", "Basic "+creds)
	req.Header.Add("Content-Type", "application/xml")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var svcResponse *AssetViewServiceResponse
	switch resp.StatusCode {
	case
		http.StatusOK:
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(resp.Body)
		//log.Print(string(buf.Bytes()))

		if err != nil {
			return nil, err
		}
		err = xml.Unmarshal(buf.Bytes(), &svcResponse)
		if err != nil {
			return nil, err
		}
		if !strings.EqualFold(svcResponse.ResponseCode, "SUCCESS") {
			return nil, fmt.Errorf("Got an unexpected XML status code: %v - %s", svcResponse.ResponseCode, svcResponse.ResponseErrorDetails.ErrorMessage)
		}
	default:
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Got an unexpected HTTP status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("Got an unexpected HTTP status code: %v - %s", resp.StatusCode, buf.String())
	}
	return svcResponse, nil
}
