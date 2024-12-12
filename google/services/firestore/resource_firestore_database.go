// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package firestore

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceFirestoreDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirestoreDatabaseCreate,
		Read:   resourceFirestoreDatabaseRead,
		Update: resourceFirestoreDatabaseUpdate,
		Delete: resourceFirestoreDatabaseDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirestoreDatabaseImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The location of the database. Available locations are listed at
https://cloud.google.com/firestore/docs/locations.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the database, which will become the final
component of the database's resource name. This value should be 4-63
characters. Valid characters are /[a-z][0-9]-/ with first character
a letter and the last a letter or a number. Must not be
UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/.
"(default)" database id is also valid.`,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: verify.ValidateEnum([]string{"FIRESTORE_NATIVE", "DATASTORE_MODE"}),
				Description: `The type of the database.
See https://cloud.google.com/datastore/docs/firestore-or-datastore
for information about how to choose. Possible values: ["FIRESTORE_NATIVE", "DATASTORE_MODE"]`,
			},
			"app_engine_integration_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"ENABLED", "DISABLED", ""}),
				Description:  `The App Engine integration mode to use for this database. Possible values: ["ENABLED", "DISABLED"]`,
			},
			"cmek_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The CMEK (Customer Managed Encryption Key) configuration for a Firestore
database. If not present, the database is secured by the default Google
encryption key.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The resource ID of a Cloud KMS key. If set, the database created will
be a Customer-managed Encryption Key (CMEK) database encrypted with
this key. This feature is allowlist only in initial launch.

Only keys in the same location as this database are allowed to be used
for encryption. For Firestore's nam5 multi-region, this corresponds to Cloud KMS
multi-region us. For Firestore's eur3 multi-region, this corresponds to
Cloud KMS multi-region europe. See https://cloud.google.com/kms/docs/locations.

This value should be the KMS key resource ID in the format of
'projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}'.
How to retrieve this resource ID is listed at
https://cloud.google.com/kms/docs/getting-resource-ids#getting_the_id_for_a_key_and_version.`,
						},
						"active_key_version": {
							Type:     schema.TypeList,
							Computed: true,
							Description: `Currently in-use KMS key versions (https://cloud.google.com/kms/docs/resource-hierarchy#key_versions).
During key rotation (https://cloud.google.com/kms/docs/key-rotation), there can be
multiple in-use key versions.

The expected format is
'projects/{project_id}/locations/{kms_location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{key_version}'.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"concurrency_mode": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"OPTIMISTIC", "PESSIMISTIC", "OPTIMISTIC_WITH_ENTITY_GROUPS", ""}),
				Description:  `The concurrency control mode to use for this database. Possible values: ["OPTIMISTIC", "PESSIMISTIC", "OPTIMISTIC_WITH_ENTITY_GROUPS"]`,
			},
			"delete_protection_state": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED", ""}),
				Description: `State of delete protection for the database.
When delete protection is enabled, this database cannot be deleted.
The default value is 'DELETE_PROTECTION_STATE_UNSPECIFIED', which is currently equivalent to 'DELETE_PROTECTION_DISABLED'.
**Note:** Additionally, to delete this database using 'terraform destroy', 'deletion_policy' must be set to 'DELETE'. Possible values: ["DELETE_PROTECTION_STATE_UNSPECIFIED", "DELETE_PROTECTION_ENABLED", "DELETE_PROTECTION_DISABLED"]`,
			},
			"point_in_time_recovery_enablement": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"POINT_IN_TIME_RECOVERY_ENABLED", "POINT_IN_TIME_RECOVERY_DISABLED", ""}),
				Description: `Whether to enable the PITR feature on this database.
If 'POINT_IN_TIME_RECOVERY_ENABLED' is selected, reads are supported on selected versions of the data from within the past 7 days.
versionRetentionPeriod and earliestVersionTime can be used to determine the supported versions. These include reads against any timestamp within the past hour
and reads against 1-minute snapshots beyond 1 hour and within 7 days.
If 'POINT_IN_TIME_RECOVERY_DISABLED' is selected, reads are supported on any version of the data from within the past 1 hour. Default value: "POINT_IN_TIME_RECOVERY_DISABLED" Possible values: ["POINT_IN_TIME_RECOVERY_ENABLED", "POINT_IN_TIME_RECOVERY_DISABLED"]`,
				Default: "POINT_IN_TIME_RECOVERY_DISABLED",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The timestamp at which this database was created.`,
			},
			"earliest_version_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The earliest timestamp at which older versions of the data can be read from the database. See versionRetentionPeriod above; this field is populated with now - versionRetentionPeriod.
This value is continuously updated, and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. This checksum is computed by the server based on the value of other fields,
and may be sent on update and delete requests to ensure the client has an
up-to-date value before proceeding.`,
			},
			"key_prefix": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The keyPrefix for this database.
This keyPrefix is used, in combination with the project id ("~") to construct the application id
that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes.
This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The system-generated UUID4 for this Database.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The timestamp at which this database was most recently updated.`,
			},
			"version_retention_period": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The period during which past versions of data are retained in the database.
Any read or query can specify a readTime within this window, and will read the state of the database at that time.
If the PITR feature is enabled, the retention period is 7 days. Otherwise, the retention period is 1 hour.
A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".`,
			},
			"deletion_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Deletion behavior for this database.
If the deletion policy is 'ABANDON', the database will be removed from Terraform state but not deleted from Google Cloud upon destruction.
If the deletion policy is 'DELETE', the database will both be removed from Terraform state and deleted from Google Cloud upon destruction.
The default value is 'ABANDON'.
See also 'delete_protection'.`,
				Default: "ABANDON",
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirestoreDatabaseCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandFirestoreDatabaseName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	locationIdProp, err := expandFirestoreDatabaseLocationId(d.Get("location_id"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("location_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(locationIdProp)) && (ok || !reflect.DeepEqual(v, locationIdProp)) {
		obj["locationId"] = locationIdProp
	}
	typeProp, err := expandFirestoreDatabaseType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	concurrencyModeProp, err := expandFirestoreDatabaseConcurrencyMode(d.Get("concurrency_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("concurrency_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(concurrencyModeProp)) && (ok || !reflect.DeepEqual(v, concurrencyModeProp)) {
		obj["concurrencyMode"] = concurrencyModeProp
	}
	appEngineIntegrationModeProp, err := expandFirestoreDatabaseAppEngineIntegrationMode(d.Get("app_engine_integration_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_integration_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(appEngineIntegrationModeProp)) && (ok || !reflect.DeepEqual(v, appEngineIntegrationModeProp)) {
		obj["appEngineIntegrationMode"] = appEngineIntegrationModeProp
	}
	pointInTimeRecoveryEnablementProp, err := expandFirestoreDatabasePointInTimeRecoveryEnablement(d.Get("point_in_time_recovery_enablement"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("point_in_time_recovery_enablement"); !tpgresource.IsEmptyValue(reflect.ValueOf(pointInTimeRecoveryEnablementProp)) && (ok || !reflect.DeepEqual(v, pointInTimeRecoveryEnablementProp)) {
		obj["pointInTimeRecoveryEnablement"] = pointInTimeRecoveryEnablementProp
	}
	deleteProtectionStateProp, err := expandFirestoreDatabaseDeleteProtectionState(d.Get("delete_protection_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("delete_protection_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(deleteProtectionStateProp)) && (ok || !reflect.DeepEqual(v, deleteProtectionStateProp)) {
		obj["deleteProtectionState"] = deleteProtectionStateProp
	}
	etagProp, err := expandFirestoreDatabaseEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	cmekConfigProp, err := expandFirestoreDatabaseCmekConfig(d.Get("cmek_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cmek_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(cmekConfigProp)) && (ok || !reflect.DeepEqual(v, cmekConfigProp)) {
		obj["cmekConfig"] = cmekConfigProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases?databaseId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Database: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Database: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirestoreOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Database", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Database: %s", err)
	}

	if err := d.Set("name", flattenFirestoreDatabaseName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Database %q: %#v", d.Id(), res)

	return resourceFirestoreDatabaseRead(d, meta)
}

func resourceFirestoreDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirestoreDatabase %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("deletion_policy"); !ok {
		if err := d.Set("deletion_policy", "ABANDON"); err != nil {
			return fmt.Errorf("Error setting deletion_policy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	if err := d.Set("name", flattenFirestoreDatabaseName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("location_id", flattenFirestoreDatabaseLocationId(res["locationId"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("type", flattenFirestoreDatabaseType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("concurrency_mode", flattenFirestoreDatabaseConcurrencyMode(res["concurrencyMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("app_engine_integration_mode", flattenFirestoreDatabaseAppEngineIntegrationMode(res["appEngineIntegrationMode"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("point_in_time_recovery_enablement", flattenFirestoreDatabasePointInTimeRecoveryEnablement(res["pointInTimeRecoveryEnablement"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("key_prefix", flattenFirestoreDatabaseKeyPrefix(res["key_prefix"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("delete_protection_state", flattenFirestoreDatabaseDeleteProtectionState(res["deleteProtectionState"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("etag", flattenFirestoreDatabaseEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("create_time", flattenFirestoreDatabaseCreateTime(res["create_time"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("update_time", flattenFirestoreDatabaseUpdateTime(res["update_time"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("uid", flattenFirestoreDatabaseUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("version_retention_period", flattenFirestoreDatabaseVersionRetentionPeriod(res["versionRetentionPeriod"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("earliest_version_time", flattenFirestoreDatabaseEarliestVersionTime(res["earliestVersionTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}
	if err := d.Set("cmek_config", flattenFirestoreDatabaseCmekConfig(res["cmekConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Database: %s", err)
	}

	return nil
}

func resourceFirestoreDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	typeProp, err := expandFirestoreDatabaseType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	concurrencyModeProp, err := expandFirestoreDatabaseConcurrencyMode(d.Get("concurrency_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("concurrency_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, concurrencyModeProp)) {
		obj["concurrencyMode"] = concurrencyModeProp
	}
	appEngineIntegrationModeProp, err := expandFirestoreDatabaseAppEngineIntegrationMode(d.Get("app_engine_integration_mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("app_engine_integration_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, appEngineIntegrationModeProp)) {
		obj["appEngineIntegrationMode"] = appEngineIntegrationModeProp
	}
	pointInTimeRecoveryEnablementProp, err := expandFirestoreDatabasePointInTimeRecoveryEnablement(d.Get("point_in_time_recovery_enablement"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("point_in_time_recovery_enablement"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pointInTimeRecoveryEnablementProp)) {
		obj["pointInTimeRecoveryEnablement"] = pointInTimeRecoveryEnablementProp
	}
	deleteProtectionStateProp, err := expandFirestoreDatabaseDeleteProtectionState(d.Get("delete_protection_state"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("delete_protection_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, deleteProtectionStateProp)) {
		obj["deleteProtectionState"] = deleteProtectionStateProp
	}
	etagProp, err := expandFirestoreDatabaseEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Database %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("concurrency_mode") {
		updateMask = append(updateMask, "concurrencyMode")
	}

	if d.HasChange("app_engine_integration_mode") {
		updateMask = append(updateMask, "appEngineIntegrationMode")
	}

	if d.HasChange("point_in_time_recovery_enablement") {
		updateMask = append(updateMask, "pointInTimeRecoveryEnablement")
	}

	if d.HasChange("delete_protection_state") {
		updateMask = append(updateMask, "deleteProtectionState")
	}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating Database %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Database %q: %#v", d.Id(), res)
		}

		err = FirestoreOperationWaitTime(
			config, res, project, "Updating Database", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceFirestoreDatabaseRead(d, meta)
}

func resourceFirestoreDatabaseDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Database: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirestoreBasePath}}projects/{{project}}/databases/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	if deletionPolicy := d.Get("deletion_policy"); deletionPolicy != "DELETE" {
		log.Printf("[WARN] Firestore database %q deletion_policy is not set to 'DELETE', skipping deletion", d.Get("name").(string))
		return nil
	}
	if deleteProtection := d.Get("delete_protection_state"); deleteProtection == "DELETE_PROTECTION_ENABLED" {
		return fmt.Errorf("Cannot delete Firestore database %s: Delete Protection is enabled. Set delete_protection_state to DELETE_PROTECTION_DISABLED for this resource and run \"terraform apply\" before attempting to delete it.", d.Get("name").(string))
	}

	log.Printf("[DEBUG] Deleting Database %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Database")
	}

	log.Printf("[DEBUG] Finished deleting Database %q: %#v", d.Id(), res)
	return nil
}

func resourceFirestoreDatabaseImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/databases/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("deletion_policy", "ABANDON"); err != nil {
		return nil, fmt.Errorf("Error setting deletion_policy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenFirestoreDatabaseName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.NameFromSelfLinkStateFunc(v)
}

func flattenFirestoreDatabaseLocationId(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseConcurrencyMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseAppEngineIntegrationMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabasePointInTimeRecoveryEnablement(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseKeyPrefix(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseDeleteProtectionState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseVersionRetentionPeriod(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseEarliestVersionTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseCmekConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenFirestoreDatabaseCmekConfigKmsKeyName(original["kmsKeyName"], d, config)
	transformed["active_key_version"] =
		flattenFirestoreDatabaseCmekConfigActiveKeyVersion(original["activeKeyVersion"], d, config)
	return []interface{}{transformed}
}
func flattenFirestoreDatabaseCmekConfigKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirestoreDatabaseCmekConfigActiveKeyVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirestoreDatabaseName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/databases/{{name}}")
}

func expandFirestoreDatabaseLocationId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseConcurrencyMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseAppEngineIntegrationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabasePointInTimeRecoveryEnablement(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseDeleteProtectionState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseCmekConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandFirestoreDatabaseCmekConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	transformedActiveKeyVersion, err := expandFirestoreDatabaseCmekConfigActiveKeyVersion(original["active_key_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedActiveKeyVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["activeKeyVersion"] = transformedActiveKeyVersion
	}

	return transformed, nil
}

func expandFirestoreDatabaseCmekConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirestoreDatabaseCmekConfigActiveKeyVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
