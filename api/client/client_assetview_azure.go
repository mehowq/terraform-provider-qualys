package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const apiAzurePath = "/azure/connectors"

// GetAssetViewAzureConnector gets an Azure Connector details with a specific Connector ID from the server
func (c *Client) GetAssetViewAzureConnector(id string) (*AssetViewAzureConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s/%s", apiAzurePath, id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	connector := &AssetViewAzureConnector{}
	err = json.NewDecoder(body).Decode(connector)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

// GetAllAssetViewAzureConnectors gets Azure Connector details of all connectors currently configured
func (c *Client) GetAllAssetViewAzureConnectors(pageNo int, pageSize int) (*[]AssetViewAzureConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s?pageNo=%d&pageSize=%d", apiAzurePath, pageNo, pageSize), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	type ConnectorsContent struct {
		Connectors []AssetViewAzureConnector `json:"content"`
	}

	var content ConnectorsContent
	if err := json.Unmarshal(buf.Bytes(), &content); err != nil {
		return nil, err
	}

	return &content.Connectors, nil
}

// NewAssetViewAzureConnector creates new Azure Connector
func (c *Client) NewAssetViewAzureConnector(connector *AssetViewAzureConnector) (*AssetViewAzureConnector, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(apiAzurePath, "POST", buf)
	if err != nil {
		return nil, err
	}
	newConnector := &AssetViewAzureConnector{}
	err = json.NewDecoder(body).Decode(newConnector)
	if err != nil {
		return nil, err
	}
	return newConnector, nil
}

// UpdateAssetViewAzureConnector updates details of the given Azure Connector
func (c *Client) UpdateAssetViewAzureConnector(connector *AssetViewAzureConnector) error {
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

// DeleteAssetViewAzureConnector removes AssetViewAzureConnector from the server
func (c *Client) DeleteAssetViewAzureConnector(connectorId string) error {
	body := fmt.Sprintf("[\"%s\"]", connectorId)
	buf := bytes.Buffer{}
	buf.WriteString(body)
	_, err := c.httpRequest(apiAzurePath, "DELETE", buf)
	if err != nil {
		return err
	}
	return nil
}
