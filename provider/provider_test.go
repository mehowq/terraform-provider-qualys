package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"qualys": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("QUALYS_API_PLATFROM"); v == "" {
		t.Fatal("QUALYS_API_PLATFROM environment variable must be set for acceptance tests")
	}
	if v := os.Getenv("QUALYS_API"); v == "" {
		t.Fatal("QUALYS_API environment variable must be set for acceptance tests")
	}
	if v := os.Getenv("QUALYS_API_PORT"); v == "" {
		t.Fatal("QUALYS_API_PORT environment variable must be set for acceptance tests")
	}
	if v := os.Getenv("QUALYS_API_USERNAME"); v == "" {
		t.Fatal("QUALYS_API_USERNAME environment variable must be set for acceptance tests")
	}
	if v := os.Getenv("QUALYS_API_PASSWORD"); v == "" {
		t.Fatal("QUALYS_API_PASSWORD environment variable must be set for acceptance tests")
	}
}