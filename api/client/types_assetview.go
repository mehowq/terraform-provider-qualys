package client

import "encoding/xml"

type AssetViewServiceRequest struct {
	XMLName       xml.Name          `xml:"ServiceRequest"`
	Filters       *AssetViewFilters `xml:"filters"`
	AssetViewData AssetViewData     `xml:"data"`
}
type AssetViewFilters struct {
	XMLName  xml.Name                   `xml:"filters"`
	Criteria []AssetViewFiltersCriteria `xml:"Criteria"`
}
type AssetViewFiltersCriteria struct {
	XMLName  xml.Name `xml:"Criteria"`
	Value    string   `xml:",chardata"`
	Field    string   `xml:"field,attr"`
	Operator string   `xml:"operator,attr"`
}
type AssetViewServiceResponse struct {
	XMLName              xml.Name                      `xml:"ServiceResponse"`
	ResponseCode         string                        `xml:"responseCode"`
	ResponseErrorDetails AssetViewResponseErrorDetails `xml:"responseErrorDetails"`
	Count                int                           `xml:"count"`
	HasMoreRecords       *bool                         `xml:"hasMoreRecords"`
	AssetViewData        AssetViewData                 `xml:"data"`
}
type AssetViewResponseErrorDetails struct {
	XMLName      xml.Name `xml:"responseErrorDetails"`
	ErrorMessage string   `xml:"errorMessage"`
}
type AssetViewData struct {
	XMLName                  xml.Name                  `xml:"data"`
	AssetViewAzureConnectors []AssetViewAzureConnector `xml:"AzureAssetDataConnector"`
	AssetViewTags            []AssetViewDataTag        `xml:"Tag"`
}
type AssetViewAzureConnector struct {
	XMLName      xml.Name                   `xml:"AzureAssetDataConnector"`
	ConnectorId  *int                       `xml:"id"`
	Name         *string                    `xml:"name"`
	Description  string                     `xml:"description"`
	DefaultTags  *AssetViewDataTagsChildren `xml:"defaultTags"`
	IsGovCloud   bool                       `xml:"isGovCloudConfigured"`
	AuthRecord   AssetViewDataAuthRecord    `xml:"authRecord"`
	LastSyncedOn *string                    `xml:"lastSync"`
	TotalAssets  *int                       `xml:"totalAssets"`
	State        *string                    `xml:"connectorState"`
	Type         *string                    `xml:"type"`
	Disabled     bool                       `xml:"disabled"`
}
type AssetViewDataTagsChildren struct {
	XMLName  xml.Name               `xml:"children"`
	TagsSet  *AssetViewDataTagsSet  `xml:"set"`
	TagsList *AssetViewDataTagsList `xml:"list"`
}
type AssetViewDataTagsSet struct {
	XMLName   xml.Name                 `xml:"set"`
	TagSimple []AssetViewDataTagSimple `xml:"TagSimple"`
}
type AssetViewDataTagsList struct {
	XMLName   xml.Name                 `xml:"list"`
	TagSimple []AssetViewDataTagSimple `xml:"TagSimple"`
}
type AssetViewDataTagSimple struct {
	XMLName xml.Name `xml:"TagSimple"`
	Id      *int     `xml:"id"`
	Name    *string  `xml:"name"`
}
type AssetViewDataAuthRecord struct {
	XMLName           xml.Name `xml:"authRecord"`
	DirectoryId       *string  `xml:"directoryId"`
	SubscriptionId    *string  `xml:"subscriptionId"`
	ApplicationId     *string  `xml:"applicationId"`
	AuthenticationKey *string  `xml:"authenticationKey"`
}

type AssetViewAWSConnector struct {
	ConnectorId       string `xml:"id"`
	Name              string `xml:"name"`
	Description       string `xml:"description"`
	AWSAccountId      string `xml:"awsAccountId"`
	ARN               string `xml:"arn"`
	ExternalId        string `xml:"externalId"`
	IsGovCloud        bool   `xml:"isGovCloudConfigured"`
	IsChinaRegion     bool   `xml:"isChinaRegion"`
	IsPortalConnector bool   `xml:"isPortalConnector"`
	LastSyncedOn      string `xml:"lastSync"`
	TotalAssets       int    `xml:"totalAssets"`
	State             string `xml:"connectorState"`
	Type              string `xml:"type"`
	Disabled          string `xml:"disabled"`
}

type AssetViewDataTag struct {
	XMLName  xml.Name                   `xml:"Tag"`
	TagId    *int                       `xml:"id"`
	Name     *string                    `xml:"name"`
	RuleType *string                    `xml:"ruleType"`
	RuleText *string                    `xml:"ruleText"`
	Created  *string                    `xml:"created"`
	Modified *string                    `xml:"modified"`
	Color    *string                    `xml:"color"`
	Children *AssetViewDataTagsChildren `xml:"children"`
}
