
# Terraform Qualys Cloud Connector Provider

This provider allows managing Azure and AWS Qualys Cloud Connectors configured in Qualys Portal Cloud View.

https://www.qualys.com/docs/qualys-cloudview-api-user-guide.pdf

https://qualysguard.qualys.eu/cloudview-api/swagger-ui.html#!/


## Compiling

For Windows
```
go build -o terraform-provider-qualys_v0.1.0.exe
```

For Linux
```
GOOS=linux GOARCH=amd64 go build -o terraform-provider-qualys_v0.1.0
# make it executable
git update-index --chmod=+x .terraform/plugins/linux_amd64/terraform-provider-qualys_v0.1.0
git commit -m "updated permissions"
git push
```

Place the compiled files into:
.terrafrom/plugins/linux_amd64/terraform-provider-qualys_v0.1.0
.terrafrom/plugins/windows_amd64/terraform-provider-qualys_v0.1.0.exe

## Testing

The code contains basic acceptance tests that you can run:
```
TF_ACC=true go test -v provider/*
```

## Debugging
Use `log.Print()` and set environment variables:
```
export TF_LOG=1
export TF_LOG_PATH="C:\Temp\mylog.txt"
```

## Environment variables

To run the tests or prevent storing the senstive information please use the following environment variables:

- QUALYS_API_PLATFORM
- QUALYS_API
- QUALYS_API_PORT
- QUALYS_API_USERNAME
- QUALYS_API_PASSWORD

## Qualys Azure Cloud Connector resource example

Please note modifying `subscription_id` forces redeployment as the API doesn't allow to update this property.

```
provider "qualys" {
  platform = "https://qualysguard.qualys.eu"
  api      = "cloudview-api/rest/v1"
  port     = 443
  username = "<the API username>"
  password = "<the API password>"
}

resource "qualys_azure_connector" "example" {
  name               = "AzConnector"
  description        = "Description"
  directory_id       = "<Azure AppID Directory ID>"
  subscription_id    = "<Azure AppID Subscription ID>"
  application_id     = "<Azure AppID>"
  authentication_key = "<Azure AppID Authentication Key>"
  is_gov_cloud       = false
}
```

## Qualys Azure Cloud Connector data source example

```
provider "qualys" {
  platform = "https://qualysguard.qualys.eu"
  api      = "cloudview-api/rest/v1"
  port     = 443
  username = "<the API username>"
  password = "<the API password>"
}

data "qualys_azure_connector" "ds_by_id_example" {
     connector_id = "<Qualys Cloud Connector ID>"
}
data "qualys_azure_connector" "ds_by_subid_example" {
     subscription_id = "<Subscription ID>"
}
```

## Attributes Reference
The following attributes are exported:
- id - Autogenerated Qualys Cloud Connector ID (UUID)
- name
- description
- directory_id
- subscription_id
- application_id
- is_gov_cloud
- last_synced_on
- total_assets
- state