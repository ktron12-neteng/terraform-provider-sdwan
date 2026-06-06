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
func TestAccSdwanTopologyHubAndSpokeProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_hub_and_spoke_feature.test", "spokes.0.name", "Spokes"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_hub_and_spoke_feature.test", "spokes.0.hubs.0.preference", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTopologyHubAndSpokeProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTopologyHubAndSpokeProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTopologyHubAndSpokeProfileParcelConfig_minimum() string {
	config := `resource "sdwan_topology_hub_and_spoke_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	vpns = ["VPN10"]` + "\n"
	config += `	hubs = ["dab0b0a9-32bb-4c1b-a32b-3472c1b1c992"]` + "\n"
	config += `	spokes = [{` + "\n"
	config += `	  name = "Spokes"` + "\n"
	config += `	  site_ids = ["ad8729e6-d454-4008-ab1f-1fc08d82706e"]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTopologyHubAndSpokeProfileParcelConfig_all() string {
	config := `resource "sdwan_topology_hub_and_spoke_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	vpns = ["VPN10"]` + "\n"
	config += `	hubs = ["dab0b0a9-32bb-4c1b-a32b-3472c1b1c992"]` + "\n"
	config += `	spokes = [{` + "\n"
	config += `	  name = "Spokes"` + "\n"
	config += `	  site_ids = ["ad8729e6-d454-4008-ab1f-1fc08d82706e"]` + "\n"
	config += `	  hubs = [{` + "\n"
	config += `		site_ids = ["dab0b0a9-32bb-4c1b-a32b-3472c1b1c992"]` + "\n"
	config += `		preference = 10` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
