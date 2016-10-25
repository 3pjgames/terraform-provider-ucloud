package ucloud

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ucloud": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	publicKey := os.Getenv("UCLOUD_PUBLIC_KEY")
	privateKey := os.Getenv("UCLOUD_PRIVATE_KEY")
	region := os.Getenv("UCLOUD_REGION")
	zone := os.Getenv("UCLOUD_ZONE")

	if publicKey == "" {
		t.Fatal("UCLOUD_PUBLIC_KEY is not set")
	}
	if privateKey == "" {
		t.Fatal("UCLOUD_PRIVATE_KEY is not set")
	}
	if region == "" {
		t.Fatal("UCLOUD_REGION is not set")
	}
	if zone == "" {
		t.Fatal("UCLOUD_ZONE is not set")
	}
}
