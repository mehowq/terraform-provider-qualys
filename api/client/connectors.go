package client

type AzureConnector struct {
	ConnectorId string			`json:"connectorId"`
	Name        string   		`json:"name"`
	Description string   		`json:"description"`
	DirectoryId string			`json:"directoryId"`
	SubscriptionId string		`json:"subscriptionId"`
	ApplicationId string		`json:"applicationId"`
	AuthenticationKey string	`json:"authenticationKey"`
	IsGovCloud bool				`json:"isGovCloud"`
	LastSyncedOn string			`json:"lastSyncedOn"`
	TotalAssets int				`json:"totalAssets"`
	State string				`json:"state"`
}
