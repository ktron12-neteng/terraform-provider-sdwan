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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanNetworkHierarchyCflowdProfileParcel(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "active_flow_timeout", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "inactive_flow_timeout", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "flow_refresh", "600"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "sampling_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collect_tloc_loopback", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "protocol", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "tos", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "remarked_dscp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collectors.0.vpn_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collectors.0.ip_address", "10.10.101.2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collectors.0.port", "4739"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collectors.0.export_spreading", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collectors.0.bfd_metrics_exporting", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_cflowd_feature.test", "collectors.0.exporting_interval", "60"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchyCflowdProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanNetworkHierarchyCflowdProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanNetworkHierarchyCflowdProfileParcelConfig_minimum() string {
	config := `resource "sdwan_network_hierarchy_cflowd_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	network_hierarchy_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"` + "\n"
	config += `	active_flow_timeout = 600` + "\n"
	config += `	inactive_flow_timeout = 60` + "\n"
	config += `	flow_refresh = 600` + "\n"
	config += `	sampling_interval = 1` + "\n"
	config += `	collect_tloc_loopback = false` + "\n"
	config += `	protocol = "ipv4"` + "\n"
	config += `	tos = false` + "\n"
	config += `	remarked_dscp = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanNetworkHierarchyCflowdProfileParcelConfig_all() string {
	config := `resource "sdwan_network_hierarchy_cflowd_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	network_hierarchy_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"` + "\n"
	config += `	active_flow_timeout = 600` + "\n"
	config += `	inactive_flow_timeout = 60` + "\n"
	config += `	flow_refresh = 600` + "\n"
	config += `	sampling_interval = 1` + "\n"
	config += `	collect_tloc_loopback = false` + "\n"
	config += `	protocol = "ipv4"` + "\n"
	config += `	tos = false` + "\n"
	config += `	remarked_dscp = false` + "\n"
	config += `	collectors = [{` + "\n"
	config += `	  vpn_id = 10` + "\n"
	config += `	  ip_address = "10.10.101.2"` + "\n"
	config += `	  port = 4739` + "\n"
	config += `	  export_spreading = false` + "\n"
	config += `	  bfd_metrics_exporting = false` + "\n"
	config += `	  exporting_interval = 60` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
