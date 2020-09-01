package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// GetAssetViewAzureConnector gets an AssetView Azure Connector details with a specific Connector ID from the server
func (c *Client) GetAssetViewAzureConnector(id int) (*AssetViewAzureConnector, error) {
	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/get/am/azureassetdataconnector/%d", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	var connector *AssetViewAzureConnector
	if svcResp.Count == 1 {
		connector = &svcResp.AssetViewData.AssetViewAzureConnectors[0]
	}

	return connector, nil
}

// GetAllAssetViewAzureConnectors gets AssetView Azure Connector details of all connectors currently configured
func (c *Client) GetAllAssetViewAzureConnectors() (*[]AssetViewAzureConnector, error) {
	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/search/am/azureassetdataconnector"), "POST", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	return &svcResp.AssetViewData.AssetViewAzureConnectors, nil
}

// SearchAssetViewAzureConnectors searches for AssetView Azure Connectors by given criteria
func (c *Client) SearchAssetViewAzureConnectors(criteria []AssetViewFiltersCriteria) (*[]AssetViewAzureConnector, error) {
	svcReq := new(AssetViewServiceRequest)
	filters := AssetViewFilters{
		Criteria: criteria,
	}
	svcReq.Filters = &filters
	buf := bytes.Buffer{}
	err := xml.NewEncoder(&buf).Encode(svcReq)
	//log.Print(string(buf.Bytes()))
	if err != nil {
		return nil, err
	}

	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/search/am/azureassetdataconnector"), "POST", buf)
	if err != nil {
		return nil, err
	}
	return &svcResp.AssetViewData.AssetViewAzureConnectors, nil
}

// NewAssetViewAzureConnector creates a new AssetView Azure Connector
func (c *Client) NewAssetViewAzureConnector(connector *AssetViewAzureConnector) (*AssetViewAzureConnector, error) {
	svcReq := new(AssetViewServiceRequest)
	svcReq.AssetViewData.AssetViewAzureConnectors = []AssetViewAzureConnector{*connector}
	buf := bytes.Buffer{}
	err := xml.NewEncoder(&buf).Encode(svcReq)
	if err != nil {
		return nil, err
	}

	svcResp, err := c.httpRequestAssetView("/create/am/azureassetdataconnector", "POST", buf)
	if err != nil {
		return nil, err
	}

	var newConnector *AssetViewAzureConnector
	if svcResp.Count == 1 {
		newConnector = &svcResp.AssetViewData.AssetViewAzureConnectors[0]
	}
	return newConnector, nil
}

// UpdateAssetViewAzureConnector updates details of the given AssetView Azure Connector
func (c *Client) UpdateAssetViewAzureConnector(connector *AssetViewAzureConnector) (*AssetViewAzureConnector, error) {
	svcReq := new(AssetViewServiceRequest)
	svcReq.AssetViewData.AssetViewAzureConnectors = []AssetViewAzureConnector{*connector}
	buf := bytes.Buffer{}
	err := xml.NewEncoder(&buf).Encode(svcReq)
	if err != nil {
		return nil, err
	}

	//log.Println(string(buf.Bytes()))
	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/update/am/azureassetdataconnector/%d", *connector.ConnectorId), "POST", buf)
	if err != nil {
		return nil, err
	}

	var updatedConnector *AssetViewAzureConnector
	if svcResp.Count == 1 {
		updatedConnector = &svcResp.AssetViewData.AssetViewAzureConnectors[0]
	}

	return updatedConnector, nil
}

// DeleteAssetViewAzureConnector removes AssetView Azure Connector from the server
func (c *Client) DeleteAssetViewAzureConnector(connectorId int) error {
	buf := bytes.Buffer{}
	_, err := c.httpRequestAssetView(fmt.Sprintf("/delete/am/azureassetdataconnector/%d", connectorId), "POST", buf)
	if err != nil {
		return err
	}
	return nil
}
