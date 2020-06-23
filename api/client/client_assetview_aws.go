package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const apiAWSPath = "/aws/connectors"

// GetAssetViewAWSConnector gets an AWS Connector details with a specific Connector ID from the server
func (c *Client) GetAssetViewAWSConnector(id string) (*AssetViewAWSConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s/%s", apiAWSPath, id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	connector := &AssetViewAWSConnector{}
	err = json.NewDecoder(body).Decode(connector)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

// NewAssetViewAWSConnector creates new AWS Connector
func (c *Client) NewAssetViewAWSConnector(connector *AssetViewAWSConnector) (*AssetViewAWSConnector, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(apiAWSPath, "POST", buf)
	if err != nil {
		return nil, err
	}
	newConnector := &AssetViewAWSConnector{}
	err = json.NewDecoder(body).Decode(newConnector)
	if err != nil {
		return nil, err
	}
	return newConnector, nil
}

// UpdateAssetViewAWSConnector updates details of the given AWS Connector
func (c *Client) UpdateAssetViewAWSConnector(connector *AssetViewAWSConnector) error {
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

// DeleteAssetViewAWSConnector removes AssetViewAWSConnector from the server
func (c *Client) DeleteAssetViewAWSConnector(connectorId string) error {
	body := fmt.Sprintf("[\"%s\"]", connectorId)
	buf := bytes.Buffer{}
	buf.WriteString(body)
	_, err := c.httpRequest(apiAWSPath, "DELETE", buf)
	if err != nil {
		return err
	}
	return nil
}
