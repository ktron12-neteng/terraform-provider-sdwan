// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanTopologyCustomControlProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "default_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "target_level", "SITE"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.name", "Backup_TLOC"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.type", "route"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.actions.0.set.0.preference", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.actions.0.set.0.tloc_action", "backup"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTopologyCustomControlProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTopologyCustomControlProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTopologyCustomControlProfileParcelConfig_minimum() string {
	config := `resource "sdwan_topology_custom_control_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTopologyCustomControlProfileParcelConfig_all() string {
	config := `resource "sdwan_topology_custom_control_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	default_action = "reject"` + "\n"
	config += `	target_level = "SITE"` + "\n"
	config += `	target_inbound_site_ids = ["e00b80d7-826f-40e1-a7fe-7a542d42e059"]` + "\n"
	config += `	target_outbound_site_ids = ["e00b80d7-826f-40e1-a7fe-7a542d42e059"]` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 10` + "\n"
	config += `	  name = "Backup_TLOC"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  type = "route"` + "\n"
	config += `	  ip_type = "ipv4"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		site_ids = ["e00b80d7-826f-40e1-a7fe-7a542d42e059"]` + "\n"
	config += `	}]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `      set = [{` + "\n"
	config += `			preference = 5000` + "\n"
	config += `			tloc_list_id = "33923a1b-0e38-49dd-b833-1ff47661311d"` + "\n"
	config += `			tloc_action = "backup"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
