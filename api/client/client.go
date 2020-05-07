package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"encoding/base64"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname    string
	api    		string
	port        int
	username    string
	password	string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate with API
func NewClient(hostname string, port int, api string, username string, password string) *Client {
	return &Client{
		hostname:   hostname,
		api:   		api,
		port:       port,
		username:	username,
		password:	password,
		httpClient: &http.Client{},
	}
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s%s", c.hostname, c.port, c.api, path)
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	creds := base64.StdEncoding.EncodeToString([]byte(c.username + ":" + c.password))
	req.Header.Add("Authorization", "Basic " + creds)
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
			respBody := new(bytes.Buffer)
			_, err := respBody.ReadFrom(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("got an unexpected status code: %v", resp.StatusCode)
			}
			return nil, fmt.Errorf("got an unexpected status code: %v - %s", resp.StatusCode, respBody.String())	
	}
	return resp.Body, nil
}