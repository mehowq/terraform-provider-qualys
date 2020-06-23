package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const apiAzurePath = "/azure/connectors"

// GetCloudViewAzureConnector gets an Azure Connector details with a specific Connector ID from the server
func (c *Client) GetCloudViewAzureConnector(id string) (*CloudViewAzureConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s/%s", apiAzurePath, id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	connector := &CloudViewAzureConnector{}
	err = json.NewDecoder(body).Decode(connector)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

// GetAllCloudViewAzureConnectors gets Azure Connector details of all connectors currently configured
func (c *Client) GetAllCloudViewAzureConnectors(pageNo int, pageSize int) (*[]CloudViewAzureConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s?pageNo=%d&pageSize=%d", apiAzurePath, pageNo, pageSize), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	type ConnectorsContent struct {
		Connectors []CloudViewAzureConnector `json:"content"`
	}

	var content ConnectorsContent
	if err := json.Unmarshal(buf.Bytes(), &content); err != nil {
		return nil, err
	}

	return &content.Connectors, nil
}

// NewCloudViewAzureConnector creates new Azure Connector
func (c *Client) NewCloudViewAzureConnector(connector *CloudViewAzureConnector) (*CloudViewAzureConnector, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(apiAzurePath, "POST", buf)
	if err != nil {
		return nil, err
	}
	newConnector := &CloudViewAzureConnector{}
	err = json.NewDecoder(body).Decode(newConnector)
	if err != nil {
		return nil, err
	}
	return newConnector, nil
}

// UpdateCloudViewAzureConnector updates details of the given Azure Connector
func (c *Client) UpdateCloudViewAzureConnector(connector *CloudViewAzureConnector) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("%s/%s", apiAzurePath, connector.ConnectorId), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCloudViewAzureConnector removes CloudViewAzureConnector from the server
func (c *Client) DeleteCloudViewAzureConnector(connectorId string) error {
	body := fmt.Sprintf("[\"%s\"]", connectorId)
	buf := bytes.Buffer{}
	buf.WriteString(body)
	_, err := c.httpRequest(apiAzurePath, "DELETE", buf)
	if err != nil {
		return err
	}
	return nil
}
