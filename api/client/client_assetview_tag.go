package client

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
)

// GetAssetViewTag gets AssetView Tag details
func (c *Client) GetAssetViewDataTag(id int) (*AssetViewDataTag, error) {
	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/get/am/tag/%d", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	var tag *AssetViewDataTag
	if svcResp.Count == 1 {
		tag = &svcResp.AssetViewData.AssetViewTags[0]
	}

	return tag, nil
}

// // GetAllAssetViewAzureTags gets AssetView Azure Connector details of all connectors currently configured
// func (c *Client) GetAllAssetViewAzureConnectors() (*[]AssetViewAzureConnector, error) {
// 	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/search/am/azureassetdataconnector"), "POST", bytes.Buffer{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &svcResp.AssetViewData.AssetViewAzureConnectors, nil
// }

// // SearchAssetViewAzureConnectors searches for AssetView Azure Connectors by given criteria
// func (c *Client) SearchAssetViewAzureConnectors(criteria []AssetViewFiltersCriteria) (*[]AssetViewAzureConnector, error) {
// 	svcReq := new(AssetViewServiceRequest)
// 	filters := AssetViewFilters{
// 		Criteria: criteria,
// 	}
// 	svcReq.Filters = &filters
// 	buf := bytes.Buffer{}
// 	err := xml.NewEncoder(&buf).Encode(svcReq)
// 	//log.Print(string(buf.Bytes()))
// 	if err != nil {
// 		return nil, err
// 	}

// 	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/search/am/azureassetdataconnector"), "POST", buf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &svcResp.AssetViewData.AssetViewAzureConnectors, nil
// }

// NewAssetViewTag creates a new AssetView Tag
func (c *Client) NewAssetViewDataTag(tag *AssetViewDataTag) (*AssetViewDataTag, error) {
	svcReq := new(AssetViewServiceRequest)
	svcReq.AssetViewData.AssetViewTags = []AssetViewDataTag{*tag}
	buf := bytes.Buffer{}
	err := xml.NewEncoder(&buf).Encode(svcReq)
	if err != nil {
		return nil, err
	}
	log.Print("MKTEST111")
	log.Print(string(buf.Bytes()))

	svcResp, err := c.httpRequestAssetView("/create/am/tag", "POST", buf)
	if err != nil {
		return nil, err
	}

	var newTag *AssetViewDataTag
	if svcResp.Count == 1 {
		newTag = &svcResp.AssetViewData.AssetViewTags[0]
	}
	return newTag, nil
}

// UpdateAssetViewTag updates details of the given AssetView Tag
func (c *Client) UpdateAssetViewDataTag(tag *AssetViewDataTag) (*AssetViewDataTag, error) {
	svcReq := new(AssetViewServiceRequest)
	svcReq.AssetViewData.AssetViewTags = []AssetViewDataTag{*tag}
	buf := bytes.Buffer{}
	err := xml.NewEncoder(&buf).Encode(svcReq)
	if err != nil {
		return nil, err
	}

	//log.Println(string(buf.Bytes()))
	svcResp, err := c.httpRequestAssetView(fmt.Sprintf("/update/am/tag/%d", *tag.TagId), "POST", buf)
	if err != nil {
		return nil, err
	}

	var updatedTag *AssetViewDataTag
	if svcResp.Count == 1 {
		updatedTag = &svcResp.AssetViewData.AssetViewTags[0]
	}

	return updatedTag, nil
}

// DeleteAssetViewTag removes AssetView Tag from the server
func (c *Client) DeleteAssetViewDataTag(connectorId int) error {
	buf := bytes.Buffer{}
	_, err := c.httpRequestAssetView(fmt.Sprintf("/delete/am/tag/%d", connectorId), "POST", buf)
	if err != nil {
		return err
	}
	return nil
}
