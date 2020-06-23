package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const apiAWSPath = "/aws/connectors"

// GetCloudViewAWSConnector gets an AWS Connector details with a specific Connector ID from the server
func (c *Client) GetCloudViewAWSConnector(id string) (*CloudViewAWSConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s/%s", apiAWSPath, id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	connector := &CloudViewAWSConnector{}
	err = json.NewDecoder(body).Decode(connector)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

// NewCloudViewAWSConnector creates new AWS Connector
func (c *Client) NewCloudViewAWSConnector(connector *CloudViewAWSConnector) (*CloudViewAWSConnector, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(apiAWSPath, "POST", buf)
	if err != nil {
		return nil, err
	}
	newConnector := &CloudViewAWSConnector{}
	err = json.NewDecoder(body).Decode(newConnector)
	if err != nil {
		return nil, err
	}
	return newConnector, nil
}

// UpdateCloudViewAWSConnector updates details of the given AWS Connector
func (c *Client) UpdateCloudViewAWSConnector(connector *CloudViewAWSConnector) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("%s/%s", apiAWSPath, connector.ConnectorId), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCloudViewAWSConnector removes CloudViewAWSConnector from the server
func (c *Client) DeleteCloudViewAWSConnector(connectorId string) error {
	body := fmt.Sprintf("[\"%s\"]", connectorId)
	buf := bytes.Buffer{}
	buf.WriteString(body)
	_, err := c.httpRequest(apiAWSPath, "DELETE", buf)
	if err != nil {
		return err
	}
	return nil
}
