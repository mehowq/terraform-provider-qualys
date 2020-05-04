package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mehowq/terraform-provider-qualys/api/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"platform": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"api": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:	 true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"qualys_azure_connector": dataSourceAzureConnector(),
		},
		ResourcesMap: map[string]*schema.Resource{
			//Terraform gets confused and downloads azure provider if we name it simply azure_connector
			"qualys_azure_connector": resourceAzureConnector(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	platform := d.Get("platform").(string)
	api := d.Get("api").(string)
	port := d.Get("port").(int)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	return client.NewClient(platform, port, api, username, password), nil
}