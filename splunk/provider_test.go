package splunk

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
		"splunk": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("SPLUNK_URL"); v == "" {
		t.Fatal("SPLUNK_URL must be set for acceptance tests")
	}

	if v := os.Getenv("SPLUNK_USERNAME"); v == "" {
		t.Fatal("SPLUNK_USERNAME must be set for acceptance tests. The domain is used to create and destroy record against.")
	}

	if v := os.Getenv("SPLUNK_PASSWORD"); v == "" {
		t.Fatal("SPLUNK_PASSWORD must be set for acceptance tests")
	}
}
