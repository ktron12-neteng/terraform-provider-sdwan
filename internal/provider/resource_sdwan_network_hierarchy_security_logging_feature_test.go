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
func TestAccSdwanNetworkHierarchySecurityLoggingProfileParcel(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging_feature.test", "high_speed_logging.0.name", "server1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging_feature.test", "high_speed_logging.0.vrf", "VPN10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging_feature.test", "high_speed_logging.0.server_ip", "10.1.200.253"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging_feature.test", "high_speed_logging.0.port", "2055"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging_feature.test", "utd_syslog.0.vpn", "VPN10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_network_hierarchy_security_logging_feature.test", "utd_syslog.0.server_ip", "10.1.200.253"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanNetworkHierarchySecurityLoggingProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanNetworkHierarchySecurityLoggingProfileParcelConfig_minimum() string {
	config := `resource "sdwan_network_hierarchy_security_logging_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	network_hierarchy_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanNetworkHierarchySecurityLoggingProfileParcelConfig_all() string {
	config := `resource "sdwan_network_hierarchy_security_logging_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	network_hierarchy_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"` + "\n"
	config += `	high_speed_logging = [{` + "\n"
	config += `	  name = "server1"` + "\n"
	config += `	  vrf = "VPN10"` + "\n"
	config += `	  server_ip = "10.1.200.253"` + "\n"
	config += `	  port = 2055` + "\n"
	config += `	}]` + "\n"
	config += `	utd_syslog = [{` + "\n"
	config += `	  vpn = "VPN10"` + "\n"
	config += `	  server_ip = "10.1.200.253"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
