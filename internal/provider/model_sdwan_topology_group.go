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
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TopologyGroup struct {
	Id          types.String            `tfsdk:"id"`
	Name        types.String            `tfsdk:"name"`
	Description types.String            `tfsdk:"description"`
	Solution    types.String            `tfsdk:"solution"`
	Profiles    []TopologyGroupProfiles `tfsdk:"profiles"`
	Activate    types.Bool              `tfsdk:"activate"`
	// RedeployTrigger is a pure trigger: it is NOT sent to the API (omitted from
	// toBody) and NOT read back (untouched by fromBody), so changing its value
	// forces Terraform to call Update, which re-deploys/re-activates the group.
	// Set it to a hash of upstream policy/topology content (or timestamp()) so the
	// control policy re-pushes to the vSmart when a referenced object changes.
	RedeployTrigger types.String `tfsdk:"redeploy_trigger"`
}

type TopologyGroupProfiles struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyGroup) getPath() string {
	return "/v1/topology-group/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyGroup) toBody(ctx context.Context) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Solution.IsNull() {
		body, _ = sjson.Set(body, "solution", data.Solution.ValueString())
	}
	if true {
		body, _ = sjson.Set(body, "profiles", []interface{}{})
		for _, item := range data.Profiles {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "profiles.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TopologyGroup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("solution"); value.Exists() {
		data.Solution = types.StringValue(value.String())
	} else {
		data.Solution = types.StringNull()
	}
	if value := res.Get("profiles"); value.Exists() && len(value.Array()) > 0 {
		data.Profiles = make([]TopologyGroupProfiles, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyGroupProfiles{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			data.Profiles = append(data.Profiles, item)
			return true
		})
	} else {
		if len(data.Profiles) > 0 {
			data.Profiles = []TopologyGroupProfiles{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *TopologyGroup) hasChanges(ctx context.Context, state *TopologyGroup) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if !data.Solution.Equal(state.Solution) {
		hasChanges = true
	}
	if len(data.Profiles) != len(state.Profiles) {
		hasChanges = true
	} else {
		for i := range data.Profiles {
			if !data.Profiles[i].Id.Equal(state.Profiles[i].Id) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions

// End of section. //template:end updateVersions

// Section below is generated&owned by "gen/generator.go". //template:begin processImport
func (data *TopologyGroup) processImport(ctx context.Context) {
}

// End of section. //template:end processImport

// Section below is generated&owned by "gen/generator.go". //template:begin applyFilters
// End of section. //template:end applyFilters
