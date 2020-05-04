# qualys-provider

go build -o terraform-provider-qualys

https://www.qualys.com/docs/qualys-cloudview-api-user-guide.pdf
https://qualysguard.qualys.eu/cloudview-api/swagger-ui.html#!/

Can use env vars: QUALYS_API_USERNAME and QUALYS_API_PASSWORD
instead of passing them in directly to provider

Many features of the CloudView are available through REST APIs. Access support information at www.qualys.com/support. Permissions: User must have the CloudView module enabled and api access permission.

provider "qualys" {
    platform = "https://qualysguard.qualys.eu"
    api = "/cloudview-api/rest/v1"
    port = 443
    username = "myApiUserName"
    password = "myApiPassword"
}

resource "qualys_azure_connector" "az_connector" {
    name = "MyConnectorName"
    description = "MyConnectorDescription"
    directory_id = "TenantIdGuid"
    subscription_id = "SubscriptionIdGuid"
    application_id = "AppIdGuid"
    authentication_key = "AppIdAuthKey"
    is_gov_cloud = "false"
}

output "azConnectorId" {
    value = "${qualys_azure_connector.az_connector.id}"
}
output "azConnectorName" {
    value = "${qualys_azure_connector.az_connector.name}"
}
output "azConnectorLastSyncedOn" {
    value = "${qualys_azure_connector.az_connector.last_synced_on}"
}
output "azConnectorTotalAssets" {
    value = "${qualys_azure_connector.az_connector.total_assets}"
}
output "azConnectorState" {
    value = "${qualys_azure_connector.az_connector.state}"
}