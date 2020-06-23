package client

type CloudViewAzureConnector struct {
	ConnectorId       string `json:"connectorId"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DirectoryId       string `json:"directoryId"`
	SubscriptionId    string `json:"subscriptionId"`
	ApplicationId     string `json:"applicationId"`
	AuthenticationKey string `json:"authenticationKey"`
	IsGovCloud        bool   `json:"isGovCloud"`
	LastSyncedOn      string `json:"lastSyncedOn"`
	TotalAssets       int    `json:"totalAssets"`
	State             string `json:"state"`
}

type CloudViewAWSConnector struct {
	ConnectorId       string `json:"connectorId"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AWSAccountId      string `json:"awsAccountId"`
	ARN               string `json:"arn"`
	ExternalId        string `json:"externalId"`
	IsGovCloud        bool   `json:"isGovCloud"`
	IsChinaRegion     bool   `json:"isChinaRegion"`
	IsPortalConnector bool   `json:"isPortalConnector"`
	LastSyncedOn      string `json:"lastSyncedOn"`
	TotalAssets       int    `json:"totalAssets"`
	State             string `json:"state"`
}

type AssetViewAzureConnector struct {
	ConnectorId       string `json:"connectorId"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DirectoryId       string `json:"directoryId"`
	SubscriptionId    string `json:"subscriptionId"`
	ApplicationId     string `json:"applicationId"`
	AuthenticationKey string `json:"authenticationKey"`
	IsGovCloud        bool   `json:"isGovCloud"`
	LastSyncedOn      string `json:"lastSyncedOn"`
	TotalAssets       int    `json:"totalAssets"`
	State             string `json:"state"`
}

type AssetViewAWSConnector struct {
	ConnectorId       string `json:"connectorId"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AWSAccountId      string `json:"awsAccountId"`
	ARN               string `json:"arn"`
	ExternalId        string `json:"externalId"`
	IsGovCloud        bool   `json:"isGovCloud"`
	IsChinaRegion     bool   `json:"isChinaRegion"`
	IsPortalConnector bool   `json:"isPortalConnector"`
	LastSyncedOn      string `json:"lastSyncedOn"`
	TotalAssets       int    `json:"totalAssets"`
	State             string `json:"state"`
}
