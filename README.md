# qualys-provider

go build -o terraform-provider-qualys

TF_ACC=true

https://www.qualys.com/docs/qualys-cloudview-api-user-guide.pdf
https://qualysguard.qualys.eu/cloudview-api/swagger-ui.html#!/

Can use env vars: QUALYS_API_USERNAME and QUALYS_API_PASSWORD
instead of passing them in directly to provider

Many features of the CloudView are available through REST APIs. Access support information at www.qualys.com/support. Permissions: User must have the CloudView module enabled and api access permission.

provider "qualys" {
    platform = "https://qualysguard.qualys.eu"
    api = "cloudview-api/rest/v1"
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


resource "qualys_aws_connector" "aws_connector" {
    name = ""
    description = ""
    arn = ""
    external_id = ""
    is_gov_cloud = false
    is_china_region = false
    is_portal_connector = false
}

# output "awsConnectorName" {
#     value = "${qualys_aws_connector.aws_connector.name}"
# }
# output "awsConnectorAccountId" {
#     value = "${qualys_aws_connector.aws_connector.aws_account_id}"
# }
# output "awsConnectorArn" {
#     value = "${qualys_aws_connector.aws_connector.arn}"
# }
# output "awsConnectorExternalId" {
#     value = "${qualys_aws_connector.aws_connector.external_id}"
# }
# output "awsConnectorIsGovCloud" {
#     value = "${qualys_aws_connector.aws_connector.is_gov_cloud}"
# }
# output "awsConnectorIsChinaRegion" {
#     value = "${qualys_aws_connector.aws_connector.is_china_region}"
# }
# output "awsConnectorIsPortalConnector" {
#     value = "${qualys_aws_connector.aws_connector.is_portal_connector}"
# }
# output "awsConnectorLastSyncedOn" {
#     value = "${qualys_aws_connector.aws_connector.last_synced_on}"
# }
# output "awsConnectorTotalAssets" {
#     value = "${qualys_aws_connector.aws_connector.total_assets}"
# }
# output "awsConnectorState" {
#     value = "${qualys_aws_connector.aws_connector.state}"
# }