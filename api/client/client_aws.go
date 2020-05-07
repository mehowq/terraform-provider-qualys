package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const apiAWSPath = "/aws/connectors"

// GetAWSConnector gets an AWS Connector details with a specific Connector ID from the server
func (c *Client) GetAWSConnector(id string) (*AWSConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s/%s", apiAWSPath, id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	connector := &AWSConnector{}
	err = json.NewDecoder(body).Decode(connector)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

// NewAWSConnector creates new AWS Connector
func (c *Client) NewAWSConnector(connector *AWSConnector) (*AWSConnector, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(apiAWSPath, "POST", buf)
	if err != nil {
		return nil, err
	}
	newConnector := &AWSConnector{}
	err = json.NewDecoder(body).Decode(newConnector)
	if err != nil {
		return nil, err
	}
	return newConnector, nil
}

// UpdateAWSConnector updates details of the given AWS Connector
func (c *Client) UpdateAWSConnector(connector *AWSConnector) error {
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

// DeleteAWSConnector removes AWSConnector from the server
func (c *Client) DeleteAWSConnector(connectorId string) error {
	body := fmt.Sprintf("[\"%s\"]", connectorId)
	buf := bytes.Buffer{}
	buf.WriteString(body)
	_, err := c.httpRequest(apiAWSPath, "DELETE", buf)
	if err != nil {
		return err
	}
	return nil
}