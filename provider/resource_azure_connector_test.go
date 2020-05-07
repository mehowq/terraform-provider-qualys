package provider

import (
	"testing"
	"fmt"
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mehowq/terraform-provider-qualys/api/client"
	guuid "github.com/google/uuid"
)

const testAccAzConnResType = "qualys_azure_connector"
const testAccAzConnResName = "connector_az_acctest"
const testAccAzConnResName2 = "connector_az_acctest2"
var testAccAzConnName = fmt.Sprintf("TF_AccTest_AzConn_%s", guuid.New().String())
var testAccAzConnDirId = guuid.New().String()
var testAccAzConnSubId = guuid.New().String()
var testAccAzConnAppId = guuid.New().String()
var testAccAzConnNamePostUpd = fmt.Sprintf("TF_AccTest_AzConn_%s", guuid.New().String())
var testAccAzConnDirIdPostUpd = guuid.New().String()
var testAccAzConnAppIdPostUpd = guuid.New().String()

func TestAccAzConnector_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzConnectorBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzConnectorExists(fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "name", testAccAzConnName),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "directory_id", testAccAzConnDirId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "subscription_id", testAccAzConnSubId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "application_id", testAccAzConnAppId),
					),
			},
		},
	})
}

func TestAccAzConnector_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzConnectorUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzConnectorExists(fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "name", testAccAzConnName),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "directory_id", testAccAzConnDirId),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "application_id", testAccAzConnAppId),
					),
			},
			{
				Config: testAccCheckAzConnectorUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzConnectorExists(fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName)),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "name", testAccAzConnNamePostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "directory_id", testAccAzConnDirIdPostUpd),
					resource.TestCheckResourceAttr(
						fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName), "application_id", testAccAzConnAppIdPostUpd),
					),
			},
		},
	})
}

func TestAccAzConnector_Multiple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAzConnectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAzConnectorMultiple(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzConnectorExists(fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName)),
					testAccCheckAzConnectorExists(fmt.Sprintf("%s.%s", testAccAzConnResType, testAccAzConnResName2)),
				),
			},
		},
	})
}

func testAccCheckAzConnectorBasic() string {	
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Azure Connector with random subscription and authentication details. Basic Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
	testAccAzConnResType,
	testAccAzConnResName,
	testAccAzConnName,
	testAccAzConnDirId,
	testAccAzConnSubId,
	testAccAzConnAppId,
	authKey))
}

func testAccCheckAzConnectorExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Azure Connector ID is set for %s", resource)
		}
		connectorId := rs.Primary.ID
		apiClient := testAccProvider.Meta().(*client.Client)
		_, err := apiClient.GetAzureConnector(connectorId)
		if err != nil {
			return fmt.Errorf("error fetching Azure Connector %s with ID %s: %s", resource, connectorId, err)
		}
		return nil
	}
}

func testAccCheckAzConnectorDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != testAccAzConnResType {
			continue
		}

		connectorId := rs.Primary.ID
		_, err := apiClient.GetAzureConnector(connectorId)
		if err == nil {
			return fmt.Errorf("Azure Connector %s still exists", connectorId)
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}

func testAccCheckAzConnectorUpdatePre() string {	
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Azure Connector with random subscription and authentication details. Pre Update Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
	testAccAzConnResType,
	testAccAzConnResName,
	testAccAzConnName,
	testAccAzConnDirId,
	testAccAzConnSubId,
	testAccAzConnAppId,
	authKey))
}

func testAccCheckAzConnectorUpdatePost() string {	
	authKey := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Azure Connector with random subscription and authentication details. Post Update Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
	testAccAzConnResType,
	testAccAzConnResName,
	testAccAzConnNamePostUpd,
	testAccAzConnDirIdPostUpd,
	testAccAzConnSubId,
	testAccAzConnAppIdPostUpd,
	authKey))
}

func testAccCheckAzConnectorMultiple() string {	
	authKey := guuid.New().String()
	conn2name := fmt.Sprintf("TF_AccTest_AzConn2_%s", guuid.New().String())
	conn2subId := guuid.New().String()

	return (fmt.Sprintf(`
resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Azure Connector with random subscription and authentication details. Multiple 1 Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}

resource "%s" "%s" {
	name = "%s"
	description = "Terraform Acceptance Test Azure Connector with random subscription and authentication details. Multiple 2 Test."
	directory_id = "%s"
	subscription_id = "%s"
	application_id = "%s"
	authentication_key = "%s"
	is_gov_cloud = false
}
	`,
	testAccAzConnResType,
	testAccAzConnResName,
	testAccAzConnName,
	testAccAzConnDirId,
	testAccAzConnSubId,
	testAccAzConnAppId,
	authKey,
	testAccAzConnResType,
	testAccAzConnResName2,
	conn2name,
	testAccAzConnDirId,
	conn2subId,
	testAccAzConnAppId,
	authKey))
}