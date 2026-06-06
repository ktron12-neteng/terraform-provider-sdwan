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
type NetworkHierarchy struct {
	Id    types.String            `tfsdk:"id"`
	Nodes []NetworkHierarchyNodes `tfsdk:"nodes"`
}

type NetworkHierarchyNodes struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Label         types.String `tfsdk:"label"`
	ParentId      types.String `tfsdk:"parent_id"`
	HierarchyPath types.String `tfsdk:"hierarchy_path"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NetworkHierarchy) getPath() string {
	return "/v1/network-hierarchy/"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data NetworkHierarchy) toBody(ctx context.Context) string {
	body := ""
	if true {
		body, _ = sjson.Set(body, "@this", []interface{}{})
		for _, item := range data.Nodes {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Label.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "data.label", item.Label.ValueString())
			}
			if !item.ParentId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "data.parentUuid", item.ParentId.ValueString())
			}
			if !item.HierarchyPath.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hierarchyPath", item.HierarchyPath.ValueString())
			}
			body, _ = sjson.SetRaw(body, "@this.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *NetworkHierarchy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("@this"); value.Exists() && len(value.Array()) > 0 {
		data.Nodes = make([]NetworkHierarchyNodes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchyNodes{}
			if cValue := v.Get("id"); cValue.Exists() {
				item.Id = types.StringValue(cValue.String())
			} else {
				item.Id = types.StringNull()
			}
			if cValue := v.Get("name"); cValue.Exists() {
				item.Name = types.StringValue(cValue.String())
			} else {
				item.Name = types.StringNull()
			}
			if cValue := v.Get("data.label"); cValue.Exists() {
				item.Label = types.StringValue(cValue.String())
			} else {
				item.Label = types.StringNull()
			}
			if cValue := v.Get("data.parentUuid"); cValue.Exists() {
				item.ParentId = types.StringValue(cValue.String())
			} else {
				item.ParentId = types.StringNull()
			}
			if cValue := v.Get("hierarchyPath"); cValue.Exists() {
				item.HierarchyPath = types.StringValue(cValue.String())
			} else {
				item.HierarchyPath = types.StringNull()
			}
			data.Nodes = append(data.Nodes, item)
			return true
		})
	} else {
		if len(data.Nodes) > 0 {
			data.Nodes = []NetworkHierarchyNodes{}
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *NetworkHierarchy) hasChanges(ctx context.Context, state *NetworkHierarchy) bool {
	hasChanges := false
	if len(data.Nodes) != len(state.Nodes) {
		hasChanges = true
	} else {
		for i := range data.Nodes {
			if !data.Nodes[i].Id.Equal(state.Nodes[i].Id) {
				hasChanges = true
			}
			if !data.Nodes[i].Name.Equal(state.Nodes[i].Name) {
				hasChanges = true
			}
			if !data.Nodes[i].Label.Equal(state.Nodes[i].Label) {
				hasChanges = true
			}
			if !data.Nodes[i].ParentId.Equal(state.Nodes[i].ParentId) {
				hasChanges = true
			}
			if !data.Nodes[i].HierarchyPath.Equal(state.Nodes[i].HierarchyPath) {
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
func (data *NetworkHierarchy) processImport(ctx context.Context) {
}

// End of section. //template:end processImport

// Section below is generated&owned by "gen/generator.go". //template:begin applyFilters
// End of section. //template:end applyFilters
