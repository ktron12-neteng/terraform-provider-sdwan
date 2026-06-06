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
	"fmt"
	"net/url"
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type TopologyHubAndSpoke struct {
	Id               types.String                `tfsdk:"id"`
	Version          types.Int64                 `tfsdk:"version"`
	Name             types.String                `tfsdk:"name"`
	Description      types.String                `tfsdk:"description"`
	FeatureProfileId types.String                `tfsdk:"feature_profile_id"`
	Vpns             types.Set                   `tfsdk:"vpns"`
	Hubs             types.Set                   `tfsdk:"hubs"`
	Spokes           []TopologyHubAndSpokeSpokes `tfsdk:"spokes"`
}

type TopologyHubAndSpokeSpokes struct {
	Name    types.String                    `tfsdk:"name"`
	SiteIds types.Set                       `tfsdk:"site_ids"`
	Hubs    []TopologyHubAndSpokeSpokesHubs `tfsdk:"hubs"`
}

type TopologyHubAndSpokeSpokesHubs struct {
	SiteIds    types.Set   `tfsdk:"site_ids"`
	Preference types.Int64 `tfsdk:"preference"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TopologyHubAndSpoke) getModel() string {
	return "topology_hub_and_spoke"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyHubAndSpoke) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/topology/%v/hubspoke", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyHubAndSpoke) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.Vpns.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.vpn.optionType", "global")
			var values []string
			data.Vpns.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.vpn.value", values)
		}
	}
	if !data.Hubs.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"selectedHierarchyHubs.optionType", "global")
			var values []string
			data.Hubs.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"selectedHierarchyHubs.value", values)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"spokes", []interface{}{})
		for _, item := range data.Spokes {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "name.value", item.Name.ValueString())
				}
			}
			if !item.SiteIds.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "spokeHierarchyUuid.optionType", "global")
					var values []string
					item.SiteIds.ElementsAs(ctx, &values, false)
					itemBody, _ = sjson.Set(itemBody, "spokeHierarchyUuid.value", values)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "hubSites", []interface{}{})
				for _, childItem := range item.Hubs {
					itemChildBody := ""
					if !childItem.SiteIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "hierarchyUuid.optionType", "global")
							var values []string
							childItem.SiteIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "hierarchyUuid.value", values)
						}
					}
					if !childItem.Preference.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.optionType", "global")
							itemChildBody, _ = sjson.Set(itemChildBody, "preference.value", childItem.Preference.ValueInt64())
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "hubSites.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"spokes.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TopologyHubAndSpoke) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Vpns = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpns = helpers.GetStringSet(va.Array())
		}
	}
	data.Hubs = types.SetNull(types.StringType)

	if t := res.Get(path + "selectedHierarchyHubs.optionType"); t.Exists() {
		va := res.Get(path + "selectedHierarchyHubs.value")
		if t.String() == "global" {
			data.Hubs = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "spokes"); value.Exists() && len(value.Array()) > 0 {
		data.Spokes = make([]TopologyHubAndSpokeSpokes, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyHubAndSpokeSpokes{}
			item.Name = types.StringNull()

			if t := v.Get("name.optionType"); t.Exists() {
				va := v.Get("name.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.SiteIds = types.SetNull(types.StringType)

			if t := v.Get("spokeHierarchyUuid.optionType"); t.Exists() {
				va := v.Get("spokeHierarchyUuid.value")
				if t.String() == "global" {
					item.SiteIds = helpers.GetStringSet(va.Array())
				}
			}
			if cValue := v.Get("hubSites"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Hubs = make([]TopologyHubAndSpokeSpokesHubs, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyHubAndSpokeSpokesHubs{}
					cItem.SiteIds = types.SetNull(types.StringType)

					if t := cv.Get("hierarchyUuid.optionType"); t.Exists() {
						va := cv.Get("hierarchyUuid.value")
						if t.String() == "global" {
							cItem.SiteIds = helpers.GetStringSet(va.Array())
						}
					}
					cItem.Preference = types.Int64Null()

					if t := cv.Get("preference.optionType"); t.Exists() {
						va := cv.Get("preference.value")
						if t.String() == "global" {
							cItem.Preference = types.Int64Value(va.Int())
						}
					}
					item.Hubs = append(item.Hubs, cItem)
					return true
				})
			}
			data.Spokes = append(data.Spokes, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *TopologyHubAndSpoke) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Vpns = types.SetNull(types.StringType)

	if t := res.Get(path + "target.vpn.optionType"); t.Exists() {
		va := res.Get(path + "target.vpn.value")
		if t.String() == "global" {
			data.Vpns = helpers.GetStringSet(va.Array())
		}
	}
	data.Hubs = types.SetNull(types.StringType)

	if t := res.Get(path + "selectedHierarchyHubs.optionType"); t.Exists() {
		va := res.Get(path + "selectedHierarchyHubs.value")
		if t.String() == "global" {
			data.Hubs = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.Spokes {
		keys := [...]string{"name"}
		keyValues := [...]string{data.Spokes[i].Name.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "spokes").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						} else if tt.String() == "default" {
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.Spokes[i].Name = types.StringNull()

		if t := r.Get("name.optionType"); t.Exists() {
			va := r.Get("name.value")
			if t.String() == "global" {
				data.Spokes[i].Name = types.StringValue(va.String())
			}
		}
		data.Spokes[i].SiteIds = types.SetNull(types.StringType)

		if t := r.Get("spokeHierarchyUuid.optionType"); t.Exists() {
			va := r.Get("spokeHierarchyUuid.value")
			if t.String() == "global" {
				data.Spokes[i].SiteIds = helpers.GetStringSet(va.Array())
			}
		}
		for ci := range data.Spokes[i].Hubs {
			keys := [...]string{"hierarchyUuid", "preference"}
			keyValues := [...]string{helpers.GetStringFromSet(data.Spokes[i].Hubs[ci].SiteIds).ValueString(), strconv.FormatInt(data.Spokes[i].Hubs[ci].Preference.ValueInt64(), 10)}
			keyValuesVariables := [...]string{"", ""}

			var cr gjson.Result
			r.Get("hubSites").ForEach(
				func(_, v gjson.Result) bool {
					found := false
					for ik := range keys {
						tt := v.Get(keys[ik] + ".optionType")
						vv := v.Get(keys[ik] + ".value")
						if tt.Exists() && vv.Exists() {
							if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
								found = true
								continue
							} else if tt.String() == "default" {
								continue
							}
							found = false
							break
						}
						continue
					}
					if found {
						cr = v
						return false
					}
					return true
				},
			)
			data.Spokes[i].Hubs[ci].SiteIds = types.SetNull(types.StringType)

			if t := cr.Get("hierarchyUuid.optionType"); t.Exists() {
				va := cr.Get("hierarchyUuid.value")
				if t.String() == "global" {
					data.Spokes[i].Hubs[ci].SiteIds = helpers.GetStringSet(va.Array())
				}
			}
			data.Spokes[i].Hubs[ci].Preference = types.Int64Null()

			if t := cr.Get("preference.optionType"); t.Exists() {
				va := cr.Get("preference.value")
				if t.String() == "global" {
					data.Spokes[i].Hubs[ci].Preference = types.Int64Value(va.Int())
				}
			}
		}
	}
}

// End of section. //template:end updateFromBody
