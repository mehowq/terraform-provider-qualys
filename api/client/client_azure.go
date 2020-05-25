package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

const apiAzurePath = "/azure/connectors"

// GetAzureConnector gets an Azure Connector details with a specific Connector ID from the server
func (c *Client) GetAzureConnector(id string) (*AzureConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s/%s", apiAzurePath, id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	connector := &AzureConnector{}
	err = json.NewDecoder(body).Decode(connector)
	if err != nil {
		return nil, err
	}
	return connector, nil
}

// GetAllAzureConnectors gets Azure Connector details of all connectors currently configured
func (c *Client) GetAllAzureConnectors(pageNo int, pageSize int) (*[]AzureConnector, error) {
	body, err := c.httpRequest(fmt.Sprintf("%s?pageNo=%d&pageSize=%d", apiAzurePath, pageNo, pageSize), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	type ConnectorsContent struct {
		connectors []AzureConnector `json:"content"`
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)

	var content ConnectorsContent
	if err := json.Unmarshal(buf.Bytes(), &content); err != nil {
		log.Fatal(err)
	}

	return &content.connectors, nil
}

// NewAzureConnector creates new Azure Connector
func (c *Client) NewAzureConnector(connector *AzureConnector) (*AzureConnector, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(connector)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(apiAzurePath, "POST", buf)
	if err != nil {
		return nil, err
	}
	newConnector := &AzureConnector{}
	err = json.NewDecoder(body).Decode(newConnector)
	if err != nil {
		return nil, err
	}
	return newConnector, nil
}

// UpdateAzureConnector updates details of the given Azure Connector
func (c *Client) UpdateAzureConnector(connector *AzureConnector) error {
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

// DeleteAzureConnector removes AzureConnector from the server
func (c *Client) DeleteAzureConnector(connectorId string) error {
	body := fmt.Sprintf("[\"%s\"]", connectorId)
	buf := bytes.Buffer{}
	buf.WriteString(body)
	_, err := c.httpRequest(apiAzurePath, "DELETE", buf)
	if err != nil {
		return err
	}
	return nil
}
