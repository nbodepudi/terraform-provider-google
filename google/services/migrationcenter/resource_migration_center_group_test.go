// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package migrationcenter_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccMigrationCenterGroup_migrationGroupUpdate(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckMigrationCenterGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMigrationCenterGroup_migrationGroupBasicExample(context),
			},
			{
				ResourceName:            "google_migration_center_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "group_id", "labels", "terraform_labels"},
			},
			{
				Config: testAccMigrationCenterGroup_migrationGroupUpdate(context),
			},
			{
				ResourceName:            "google_migration_center_group.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "group_id", "labels", "terraform_labels"},
			},
		},
	})
}

func testAccMigrationCenterGroup_migrationGroupUpdate(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_migration_center_group" "default" {
  location     = "us-central1"
  group_id     = "tf-test-group-test%{random_suffix}"
  description  = "Updated Terraform integration test description"
  display_name = "Updated  integration test display"
  labels       = {
    key2 = "value2"
    key = "value"
  }
}
`, context)
}
