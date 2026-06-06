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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type NetworkHierarchySecurityLogging struct {
	Id                 types.String                                      `tfsdk:"id"`
	Version            types.Int64                                       `tfsdk:"version"`
	Name               types.String                                      `tfsdk:"name"`
	Description        types.String                                      `tfsdk:"description"`
	NetworkHierarchyId types.String                                      `tfsdk:"network_hierarchy_id"`
	HighSpeedLogging   []NetworkHierarchySecurityLoggingHighSpeedLogging `tfsdk:"high_speed_logging"`
	UtdSyslog          []NetworkHierarchySecurityLoggingUtdSyslog        `tfsdk:"utd_syslog"`
}

type NetworkHierarchySecurityLoggingHighSpeedLogging struct {
	Name     types.String `tfsdk:"name"`
	Vrf      types.String `tfsdk:"vrf"`
	ServerIp types.String `tfsdk:"server_ip"`
	Port     types.Int64  `tfsdk:"port"`
}

type NetworkHierarchySecurityLoggingUtdSyslog struct {
	Vpn      types.String `tfsdk:"vpn"`
	ServerIp types.String `tfsdk:"server_ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data NetworkHierarchySecurityLogging) getModel() string {
	return "network_hierarchy_security_logging"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NetworkHierarchySecurityLogging) getPath() string {
	return fmt.Sprintf("/v1/network-hierarchy/%v/network-settings/security-logging", url.QueryEscape(data.NetworkHierarchyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is MANUALLY maintained (markers removed): highSpeedLogging.name is a plain (non-optionType) field.
func (data NetworkHierarchySecurityLogging) toBody(ctx context.Context) string {
	body := ""
	path := "data."
	if true {
		body, _ = sjson.Set(body, path+"highSpeedLogging", []interface{}{})
		for _, item := range data.HighSpeedLogging {
			itemBody := ""
			if !item.Name.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
				}
			}
			if !item.Vrf.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vrf.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vrf.value", item.Vrf.ValueString())
				}
			}
			if !item.ServerIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serverIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serverIp.value", item.ServerIp.ValueString())
				}
			}
			if !item.Port.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "port.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "port.value", item.Port.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"highSpeedLogging.-1", itemBody)
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"utdSyslog", []interface{}{})
		for _, item := range data.UtdSyslog {
			itemBody := ""
			if !item.Vpn.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpn.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpn.value", item.Vpn.ValueString())
				}
			}
			if !item.ServerIp.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "serverIp.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "serverIp.value", item.ServerIp.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"utdSyslog.-1", itemBody)
		}
	}
	return body
}

// End of manual section.

// Section below is MANUALLY maintained (markers removed): highSpeedLogging.name is a plain (non-optionType) field.
func (data *NetworkHierarchySecurityLogging) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringNull()
	data.Description = types.StringNull()
	path := "payload.data."
	if value := res.Get(path + "highSpeedLogging"); value.Exists() && len(value.Array()) > 0 {
		data.HighSpeedLogging = make([]NetworkHierarchySecurityLoggingHighSpeedLogging, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchySecurityLoggingHighSpeedLogging{}
			item.Name = types.StringNull()

			if value := v.Get("name"); value.Exists() && value.String() != "" {
				item.Name = types.StringValue(value.String())
			}
			item.Vrf = types.StringNull()

			if t := v.Get("vrf.optionType"); t.Exists() {
				va := v.Get("vrf.value")
				if t.String() == "global" {
					item.Vrf = types.StringValue(va.String())
				}
			}
			item.ServerIp = types.StringNull()

			if t := v.Get("serverIp.optionType"); t.Exists() {
				va := v.Get("serverIp.value")
				if t.String() == "global" {
					item.ServerIp = types.StringValue(va.String())
				}
			}
			item.Port = types.Int64Null()

			if t := v.Get("port.optionType"); t.Exists() {
				va := v.Get("port.value")
				if t.String() == "global" {
					item.Port = types.Int64Value(va.Int())
				}
			}
			data.HighSpeedLogging = append(data.HighSpeedLogging, item)
			return true
		})
	}
	if value := res.Get(path + "utdSyslog"); value.Exists() && len(value.Array()) > 0 {
		data.UtdSyslog = make([]NetworkHierarchySecurityLoggingUtdSyslog, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchySecurityLoggingUtdSyslog{}
			item.Vpn = types.StringNull()

			if t := v.Get("vpn.optionType"); t.Exists() {
				va := v.Get("vpn.value")
				if t.String() == "global" {
					item.Vpn = types.StringValue(va.String())
				}
			}
			item.ServerIp = types.StringNull()

			if t := v.Get("serverIp.optionType"); t.Exists() {
				va := v.Get("serverIp.value")
				if t.String() == "global" {
					item.ServerIp = types.StringValue(va.String())
				}
			}
			data.UtdSyslog = append(data.UtdSyslog, item)
			return true
		})
	}
}

// End of manual section.

// Section below is MANUALLY maintained (markers removed): highSpeedLogging.name is a plain (non-optionType) field.
func (data *NetworkHierarchySecurityLogging) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringNull()
	data.Description = types.StringNull()
	path := "payload.data."
	for i := range data.HighSpeedLogging {
		var r gjson.Result
		res.Get(path + "highSpeedLogging").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == data.HighSpeedLogging[i].Name.ValueString() {
					r = v
					return false
				}
				return true
			},
		)
		data.HighSpeedLogging[i].Name = types.StringNull()

		if value := r.Get("name"); value.Exists() && value.String() != "" {
			data.HighSpeedLogging[i].Name = types.StringValue(value.String())
		}
		data.HighSpeedLogging[i].Vrf = types.StringNull()

		if t := r.Get("vrf.optionType"); t.Exists() {
			va := r.Get("vrf.value")
			if t.String() == "global" {
				data.HighSpeedLogging[i].Vrf = types.StringValue(va.String())
			}
		}
		data.HighSpeedLogging[i].ServerIp = types.StringNull()

		if t := r.Get("serverIp.optionType"); t.Exists() {
			va := r.Get("serverIp.value")
			if t.String() == "global" {
				data.HighSpeedLogging[i].ServerIp = types.StringValue(va.String())
			}
		}
		data.HighSpeedLogging[i].Port = types.Int64Null()

		if t := r.Get("port.optionType"); t.Exists() {
			va := r.Get("port.value")
			if t.String() == "global" {
				data.HighSpeedLogging[i].Port = types.Int64Value(va.Int())
			}
		}
	}
	for i := range data.UtdSyslog {
		keys := [...]string{"vpn"}
		keyValues := [...]string{data.UtdSyslog[i].Vpn.ValueString()}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "utdSyslog").ForEach(
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
		data.UtdSyslog[i].Vpn = types.StringNull()

		if t := r.Get("vpn.optionType"); t.Exists() {
			va := r.Get("vpn.value")
			if t.String() == "global" {
				data.UtdSyslog[i].Vpn = types.StringValue(va.String())
			}
		}
		data.UtdSyslog[i].ServerIp = types.StringNull()

		if t := r.Get("serverIp.optionType"); t.Exists() {
			va := r.Get("serverIp.value")
			if t.String() == "global" {
				data.UtdSyslog[i].ServerIp = types.StringValue(va.String())
			}
		}
	}
}

// End of manual section.
