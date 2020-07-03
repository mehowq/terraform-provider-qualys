package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cloudview_api": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_CLOUDVIEW_API", ""),
			},
			"assetview_api": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_ASSETVIEW_API", ""),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_API_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("QUALYS_API_PASSWORD", ""),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"qualys_cloudview_azure_connector": dataSourceCloudViewAzureConnector(),
			"qualys_cloudview_aws_connector":   dataSourceCloudViewAWSConnector(),
			"qualys_assetview_azure_connector": dataSourceAssetViewAzureConnector(),
			// TODO
			// "qualys_assetview_aws_connector":   dataSourceAssetViewAWSConnector(),
		},
		ResourcesMap: map[string]*schema.Resource{
			//Terraform gets confused and downloads azure provider if we name it simply azure_connector
			"qualys_cloudview_azure_connector": resourceCloudViewAzureConnector(),
			"qualys_cloudview_aws_connector":   resourceCloudViewAWSConnector(),
			"qualys_assetview_azure_connector": resourceAssetViewAzureConnector(),
			// TODO
			// "qualys_assetview_aws_connector":   resourceAssetViewAWSConnector(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	cloudview_api := d.Get("cloudview_api").(string)
	assetview_api := d.Get("assetview_api").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	return client.NewClient(cloudview_api, assetview_api, username, password), nil
}
