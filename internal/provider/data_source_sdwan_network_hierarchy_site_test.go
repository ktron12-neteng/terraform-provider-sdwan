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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanNetworkHierarchySite(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_site.test", "name", "DC1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_site.test", "description", "My site"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_site.test", "parent_id", "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_network_hierarchy_site.test", "site_id", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanNetworkHierarchySiteConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanNetworkHierarchySiteConfig() string {
	config := ""
	config += `resource "sdwan_network_hierarchy_site" "test" {` + "\n"
	config += `	name = "DC1"` + "\n"
	config += `	description = "My site"` + "\n"
	config += `	parent_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"` + "\n"
	config += `	site_id = 100` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_network_hierarchy_site" "test" {
			id = sdwan_network_hierarchy_site.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
