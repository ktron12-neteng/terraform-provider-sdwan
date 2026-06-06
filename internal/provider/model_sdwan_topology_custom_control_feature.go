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
type TopologyCustomControl struct {
	Id                    types.String                     `tfsdk:"id"`
	Version               types.Int64                      `tfsdk:"version"`
	Name                  types.String                     `tfsdk:"name"`
	Description           types.String                     `tfsdk:"description"`
	FeatureProfileId      types.String                     `tfsdk:"feature_profile_id"`
	DefaultAction         types.String                     `tfsdk:"default_action"`
	TargetLevel           types.String                     `tfsdk:"target_level"`
	TargetInboundSiteIds  types.Set                        `tfsdk:"target_inbound_site_ids"`
	TargetOutboundSiteIds types.Set                        `tfsdk:"target_outbound_site_ids"`
	Sequences             []TopologyCustomControlSequences `tfsdk:"sequences"`
}

type TopologyCustomControlSequences struct {
	Id           types.Int64                                  `tfsdk:"id"`
	Name         types.String                                 `tfsdk:"name"`
	BaseAction   types.String                                 `tfsdk:"base_action"`
	Type         types.String                                 `tfsdk:"type"`
	IpType       types.String                                 `tfsdk:"ip_type"`
	MatchEntries []TopologyCustomControlSequencesMatchEntries `tfsdk:"match_entries"`
	Actions      []TopologyCustomControlSequencesActions      `tfsdk:"actions"`
}

type TopologyCustomControlSequencesMatchEntries struct {
	SiteIds types.Set `tfsdk:"site_ids"`
}
type TopologyCustomControlSequencesActions struct {
	Set []TopologyCustomControlSequencesActionsSet `tfsdk:"set"`
}

type TopologyCustomControlSequencesActionsSet struct {
	Preference types.Int64  `tfsdk:"preference"`
	TlocListId types.String `tfsdk:"tloc_list_id"`
	TlocAction types.String `tfsdk:"tloc_action"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data TopologyCustomControl) getModel() string {
	return "topology_custom_control"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data TopologyCustomControl) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/topology/%v/custom-control", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data TopologyCustomControl) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."
	if !data.DefaultAction.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"defaultAction.optionType", "global")
			body, _ = sjson.Set(body, path+"defaultAction.value", data.DefaultAction.ValueString())
		}
	}
	if !data.TargetLevel.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.level.optionType", "global")
			body, _ = sjson.Set(body, path+"target.level.value", data.TargetLevel.ValueString())
		}
	}
	if !data.TargetInboundSiteIds.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.inboundHierarchyUuid.optionType", "global")
			var values []string
			data.TargetInboundSiteIds.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.inboundHierarchyUuid.value", values)
		}
	}
	if !data.TargetOutboundSiteIds.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"target.outboundHierarchyUuid.optionType", "global")
			var values []string
			data.TargetOutboundSiteIds.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"target.outboundHierarchyUuid.value", values)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"sequences", []interface{}{})
		for _, item := range data.Sequences {
			itemBody := ""
			if !item.Id.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceId.value", item.Id.ValueInt64())
				}
			}
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceName.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceName.value", item.Name.ValueString())
				}
			}
			if !item.BaseAction.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "baseAction.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "baseAction.value", item.BaseAction.ValueString())
				}
			}
			if !item.Type.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceType.value", item.Type.ValueString())
				}
			}
			if !item.IpType.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "sequenceIpType.value", item.IpType.ValueString())
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "match.entries", []interface{}{})
				for _, childItem := range item.MatchEntries {
					itemChildBody := ""
					if !childItem.SiteIds.IsNull() {
						if true {
							itemChildBody, _ = sjson.Set(itemChildBody, "hierarchyUuid.optionType", "global")
							var values []string
							childItem.SiteIds.ElementsAs(ctx, &values, false)
							itemChildBody, _ = sjson.Set(itemChildBody, "hierarchyUuid.value", values)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "match.entries.-1", itemChildBody)
				}
			}
			if true {
				itemBody, _ = sjson.Set(itemBody, "actions", []interface{}{})
				for _, childItem := range item.Actions {
					itemChildBody := ""
					if true {
						itemChildBody, _ = sjson.Set(itemChildBody, "set", []interface{}{})
						for _, childChildItem := range childItem.Set {
							itemChildChildBody := ""
							if !childChildItem.Preference.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preference.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "preference.value", childChildItem.Preference.ValueInt64())
								}
							}
							if !childChildItem.TlocListId.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocList.refId.value", childChildItem.TlocListId.ValueString())
								}
							}
							if !childChildItem.TlocAction.IsNull() {
								if true {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocAction.optionType", "global")
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "tlocAction.value", childChildItem.TlocAction.ValueString())
								}
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "set.-1", itemChildChildBody)
						}
					}
					itemBody, _ = sjson.SetRaw(itemBody, "actions.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, path+"sequences.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *TopologyCustomControl) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.TargetLevel = types.StringNull()

	if t := res.Get(path + "target.level.optionType"); t.Exists() {
		va := res.Get(path + "target.level.value")
		if t.String() == "global" {
			data.TargetLevel = types.StringValue(va.String())
		}
	}
	data.TargetInboundSiteIds = types.SetNull(types.StringType)

	if t := res.Get(path + "target.inboundHierarchyUuid.optionType"); t.Exists() {
		va := res.Get(path + "target.inboundHierarchyUuid.value")
		if t.String() == "global" {
			data.TargetInboundSiteIds = helpers.GetStringSet(va.Array())
		}
	}
	data.TargetOutboundSiteIds = types.SetNull(types.StringType)

	if t := res.Get(path + "target.outboundHierarchyUuid.optionType"); t.Exists() {
		va := res.Get(path + "target.outboundHierarchyUuid.value")
		if t.String() == "global" {
			data.TargetOutboundSiteIds = helpers.GetStringSet(va.Array())
		}
	}
	if value := res.Get(path + "sequences"); value.Exists() && len(value.Array()) > 0 {
		data.Sequences = make([]TopologyCustomControlSequences, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := TopologyCustomControlSequences{}
			item.Id = types.Int64Null()

			if t := v.Get("sequenceId.optionType"); t.Exists() {
				va := v.Get("sequenceId.value")
				if t.String() == "global" {
					item.Id = types.Int64Value(va.Int())
				}
			}
			item.Name = types.StringNull()

			if t := v.Get("sequenceName.optionType"); t.Exists() {
				va := v.Get("sequenceName.value")
				if t.String() == "global" {
					item.Name = types.StringValue(va.String())
				}
			}
			item.BaseAction = types.StringNull()

			if t := v.Get("baseAction.optionType"); t.Exists() {
				va := v.Get("baseAction.value")
				if t.String() == "global" {
					item.BaseAction = types.StringValue(va.String())
				}
			}
			item.Type = types.StringNull()

			if t := v.Get("sequenceType.optionType"); t.Exists() {
				va := v.Get("sequenceType.value")
				if t.String() == "global" {
					item.Type = types.StringValue(va.String())
				}
			}
			item.IpType = types.StringNull()

			if t := v.Get("sequenceIpType.optionType"); t.Exists() {
				va := v.Get("sequenceIpType.value")
				if t.String() == "global" {
					item.IpType = types.StringValue(va.String())
				}
			}
			if cValue := v.Get("match.entries"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.MatchEntries = make([]TopologyCustomControlSequencesMatchEntries, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyCustomControlSequencesMatchEntries{}
					cItem.SiteIds = types.SetNull(types.StringType)

					if t := cv.Get("hierarchyUuid.optionType"); t.Exists() {
						va := cv.Get("hierarchyUuid.value")
						if t.String() == "global" {
							cItem.SiteIds = helpers.GetStringSet(va.Array())
						}
					}
					item.MatchEntries = append(item.MatchEntries, cItem)
					return true
				})
			}
			if cValue := v.Get("actions"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.Actions = make([]TopologyCustomControlSequencesActions, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := TopologyCustomControlSequencesActions{}
					if ccValue := cv.Get("set"); ccValue.Exists() && len(ccValue.Array()) > 0 {
						cItem.Set = make([]TopologyCustomControlSequencesActionsSet, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := TopologyCustomControlSequencesActionsSet{}
							ccItem.Preference = types.Int64Null()

							if t := ccv.Get("preference.optionType"); t.Exists() {
								va := ccv.Get("preference.value")
								if t.String() == "global" {
									ccItem.Preference = types.Int64Value(va.Int())
								}
							}
							ccItem.TlocAction = types.StringNull()

							if t := ccv.Get("tlocAction.optionType"); t.Exists() {
								va := ccv.Get("tlocAction.value")
								if t.String() == "global" {
									ccItem.TlocAction = types.StringValue(va.String())
								}
							}
							cItem.Set = append(cItem.Set, ccItem)
							return true
						})
					}
					item.Actions = append(item.Actions, cItem)
					return true
				})
			}
			data.Sequences = append(data.Sequences, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is MANUALLY maintained (template markers intentionally removed) to read
// the keyless "actions"/"set" lists positionally and avoid a perpetual tloc_action diff.
// Do not re-add //template markers or this will be overwritten by the generator.
func (data *TopologyCustomControl) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.DefaultAction = types.StringNull()

	if t := res.Get(path + "defaultAction.optionType"); t.Exists() {
		va := res.Get(path + "defaultAction.value")
		if t.String() == "global" {
			data.DefaultAction = types.StringValue(va.String())
		}
	}
	data.TargetLevel = types.StringNull()

	if t := res.Get(path + "target.level.optionType"); t.Exists() {
		va := res.Get(path + "target.level.value")
		if t.String() == "global" {
			data.TargetLevel = types.StringValue(va.String())
		}
	}
	data.TargetInboundSiteIds = types.SetNull(types.StringType)

	if t := res.Get(path + "target.inboundHierarchyUuid.optionType"); t.Exists() {
		va := res.Get(path + "target.inboundHierarchyUuid.value")
		if t.String() == "global" {
			data.TargetInboundSiteIds = helpers.GetStringSet(va.Array())
		}
	}
	data.TargetOutboundSiteIds = types.SetNull(types.StringType)

	if t := res.Get(path + "target.outboundHierarchyUuid.optionType"); t.Exists() {
		va := res.Get(path + "target.outboundHierarchyUuid.value")
		if t.String() == "global" {
			data.TargetOutboundSiteIds = helpers.GetStringSet(va.Array())
		}
	}
	for i := range data.Sequences {
		keys := [...]string{"sequenceId"}
		keyValues := [...]string{strconv.FormatInt(data.Sequences[i].Id.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "sequences").ForEach(
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
		data.Sequences[i].Id = types.Int64Null()

		if t := r.Get("sequenceId.optionType"); t.Exists() {
			va := r.Get("sequenceId.value")
			if t.String() == "global" {
				data.Sequences[i].Id = types.Int64Value(va.Int())
			}
		}
		data.Sequences[i].Name = types.StringNull()

		if t := r.Get("sequenceName.optionType"); t.Exists() {
			va := r.Get("sequenceName.value")
			if t.String() == "global" {
				data.Sequences[i].Name = types.StringValue(va.String())
			}
		}
		data.Sequences[i].BaseAction = types.StringNull()

		if t := r.Get("baseAction.optionType"); t.Exists() {
			va := r.Get("baseAction.value")
			if t.String() == "global" {
				data.Sequences[i].BaseAction = types.StringValue(va.String())
			}
		}
		data.Sequences[i].Type = types.StringNull()

		if t := r.Get("sequenceType.optionType"); t.Exists() {
			va := r.Get("sequenceType.value")
			if t.String() == "global" {
				data.Sequences[i].Type = types.StringValue(va.String())
			}
		}
		data.Sequences[i].IpType = types.StringNull()

		if t := r.Get("sequenceIpType.optionType"); t.Exists() {
			va := r.Get("sequenceIpType.value")
			if t.String() == "global" {
				data.Sequences[i].IpType = types.StringValue(va.String())
			}
		}
		for ci := range data.Sequences[i].MatchEntries {
			keys := [...]string{"hierarchyUuid"}
			keyValues := [...]string{helpers.GetStringFromSet(data.Sequences[i].MatchEntries[ci].SiteIds).ValueString()}
			keyValuesVariables := [...]string{""}

			var cr gjson.Result
			r.Get("match.entries").ForEach(
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
			data.Sequences[i].MatchEntries[ci].SiteIds = types.SetNull(types.StringType)

			if t := cr.Get("hierarchyUuid.optionType"); t.Exists() {
				va := cr.Get("hierarchyUuid.value")
				if t.String() == "global" {
					data.Sequences[i].MatchEntries[ci].SiteIds = helpers.GetStringSet(va.Array())
				}
			}
		}
		// Manual: the "actions" and "set" lists have no scalar key to value-match
		// on (an action item only wraps the nested "set" list), so they are read
		// positionally. The API preserves list order, so index alignment is safe.
		for ci := range data.Sequences[i].Actions {
			var cr gjson.Result
			if arr := r.Get("actions").Array(); ci < len(arr) {
				cr = arr[ci]
			}
			for cci := range data.Sequences[i].Actions[ci].Set {
				var ccr gjson.Result
				if arr := cr.Get("set").Array(); cci < len(arr) {
					ccr = arr[cci]
				}
				data.Sequences[i].Actions[ci].Set[cci].Preference = types.Int64Null()

				if t := ccr.Get("preference.optionType"); t.Exists() {
					va := ccr.Get("preference.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Set[cci].Preference = types.Int64Value(va.Int())
					}
				}
				data.Sequences[i].Actions[ci].Set[cci].TlocAction = types.StringNull()

				if t := ccr.Get("tlocAction.optionType"); t.Exists() {
					va := ccr.Get("tlocAction.value")
					if t.String() == "global" {
						data.Sequences[i].Actions[ci].Set[cci].TlocAction = types.StringValue(va.String())
					}
				}
			}
		}
	}
}

// End of manual updateFromBody section.
