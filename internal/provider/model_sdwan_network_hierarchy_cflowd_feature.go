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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type NetworkHierarchyCflowd struct {
	Id                  types.String                       `tfsdk:"id"`
	Version             types.Int64                        `tfsdk:"version"`
	Name                types.String                       `tfsdk:"name"`
	Description         types.String                       `tfsdk:"description"`
	NetworkHierarchyId  types.String                       `tfsdk:"network_hierarchy_id"`
	ActiveFlowTimeout   types.Int64                        `tfsdk:"active_flow_timeout"`
	InactiveFlowTimeout types.Int64                        `tfsdk:"inactive_flow_timeout"`
	FlowRefresh         types.Int64                        `tfsdk:"flow_refresh"`
	SamplingInterval    types.Int64                        `tfsdk:"sampling_interval"`
	CollectTlocLoopback types.Bool                         `tfsdk:"collect_tloc_loopback"`
	Protocol            types.String                       `tfsdk:"protocol"`
	Tos                 types.Bool                         `tfsdk:"tos"`
	RemarkedDscp        types.Bool                         `tfsdk:"remarked_dscp"`
	Collectors          []NetworkHierarchyCflowdCollectors `tfsdk:"collectors"`
}

type NetworkHierarchyCflowdCollectors struct {
	VpnId               types.Int64  `tfsdk:"vpn_id"`
	IpAddress           types.String `tfsdk:"ip_address"`
	Port                types.Int64  `tfsdk:"port"`
	ExportSpreading     types.Bool   `tfsdk:"export_spreading"`
	BfdMetricsExporting types.Bool   `tfsdk:"bfd_metrics_exporting"`
	ExportingInterval   types.Int64  `tfsdk:"exporting_interval"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data NetworkHierarchyCflowd) getModel() string {
	return "network_hierarchy_cflowd"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data NetworkHierarchyCflowd) getPath() string {
	return fmt.Sprintf("/v1/network-hierarchy/%v/network-settings/cflowd", url.QueryEscape(data.NetworkHierarchyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data NetworkHierarchyCflowd) toBody(ctx context.Context) string {
	body := ""
	path := "data."
	if !data.ActiveFlowTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"flowActiveTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"flowActiveTimeout.value", data.ActiveFlowTimeout.ValueInt64())
		}
	}
	if !data.InactiveFlowTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"flowInactiveTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"flowInactiveTimeout.value", data.InactiveFlowTimeout.ValueInt64())
		}
	}
	if !data.FlowRefresh.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"flowRefreshTime.optionType", "global")
			body, _ = sjson.Set(body, path+"flowRefreshTime.value", data.FlowRefresh.ValueInt64())
		}
	}
	if !data.SamplingInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"flowSamplingInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"flowSamplingInterval.value", data.SamplingInterval.ValueInt64())
		}
	}
	if !data.CollectTlocLoopback.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"collectTlocLoopback.optionType", "global")
			body, _ = sjson.Set(body, path+"collectTlocLoopback.value", data.CollectTlocLoopback.ValueBool())
		}
	}
	if !data.Protocol.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"protocol.optionType", "global")
			body, _ = sjson.Set(body, path+"protocol.value", data.Protocol.ValueString())
		}
	}
	if !data.Tos.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"customizedIpv4RecordFields.collectTos.optionType", "global")
			body, _ = sjson.Set(body, path+"customizedIpv4RecordFields.collectTos.value", data.Tos.ValueBool())
		}
	}
	if !data.RemarkedDscp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"customizedIpv4RecordFields.collectDscpOutput.optionType", "global")
			body, _ = sjson.Set(body, path+"customizedIpv4RecordFields.collectDscpOutput.value", data.RemarkedDscp.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"collectors", []interface{}{})
		for _, item := range data.Collectors {
			itemBody := ""
			if !item.VpnId.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vpnId.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vpnId.value", item.VpnId.ValueInt64())
				}
			}
			if !item.IpAddress.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "address.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "address.value", item.IpAddress.ValueString())
				}
			}
			if !item.Port.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "udpPort.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "udpPort.value", item.Port.ValueInt64())
				}
			}
			if !item.ExportSpreading.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "exportSpread.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "exportSpread.value", item.ExportSpreading.ValueBool())
				}
			}
			if !item.BfdMetricsExporting.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "bfdMetricsExport.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "bfdMetricsExport.value", item.BfdMetricsExporting.ValueBool())
				}
			}
			if !item.ExportingInterval.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "exportInterval.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "exportInterval.value", item.ExportingInterval.ValueInt64())
				}
			}
			body, _ = sjson.SetRaw(body, path+"collectors.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *NetworkHierarchyCflowd) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringNull()
	data.Description = types.StringNull()
	path := "payload.data."
	data.ActiveFlowTimeout = types.Int64Null()

	if t := res.Get(path + "flowActiveTimeout.optionType"); t.Exists() {
		va := res.Get(path + "flowActiveTimeout.value")
		if t.String() == "global" {
			data.ActiveFlowTimeout = types.Int64Value(va.Int())
		}
	}
	data.InactiveFlowTimeout = types.Int64Null()

	if t := res.Get(path + "flowInactiveTimeout.optionType"); t.Exists() {
		va := res.Get(path + "flowInactiveTimeout.value")
		if t.String() == "global" {
			data.InactiveFlowTimeout = types.Int64Value(va.Int())
		}
	}
	data.FlowRefresh = types.Int64Null()

	if t := res.Get(path + "flowRefreshTime.optionType"); t.Exists() {
		va := res.Get(path + "flowRefreshTime.value")
		if t.String() == "global" {
			data.FlowRefresh = types.Int64Value(va.Int())
		}
	}
	data.SamplingInterval = types.Int64Null()

	if t := res.Get(path + "flowSamplingInterval.optionType"); t.Exists() {
		va := res.Get(path + "flowSamplingInterval.value")
		if t.String() == "global" {
			data.SamplingInterval = types.Int64Value(va.Int())
		}
	}
	data.CollectTlocLoopback = types.BoolNull()

	if t := res.Get(path + "collectTlocLoopback.optionType"); t.Exists() {
		va := res.Get(path + "collectTlocLoopback.value")
		if t.String() == "global" {
			data.CollectTlocLoopback = types.BoolValue(va.Bool())
		}
	}
	data.Protocol = types.StringNull()

	if t := res.Get(path + "protocol.optionType"); t.Exists() {
		va := res.Get(path + "protocol.value")
		if t.String() == "global" {
			data.Protocol = types.StringValue(va.String())
		}
	}
	data.Tos = types.BoolNull()

	if t := res.Get(path + "customizedIpv4RecordFields.collectTos.optionType"); t.Exists() {
		va := res.Get(path + "customizedIpv4RecordFields.collectTos.value")
		if t.String() == "global" {
			data.Tos = types.BoolValue(va.Bool())
		}
	}
	data.RemarkedDscp = types.BoolNull()

	if t := res.Get(path + "customizedIpv4RecordFields.collectDscpOutput.optionType"); t.Exists() {
		va := res.Get(path + "customizedIpv4RecordFields.collectDscpOutput.value")
		if t.String() == "global" {
			data.RemarkedDscp = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "collectors"); value.Exists() && len(value.Array()) > 0 {
		data.Collectors = make([]NetworkHierarchyCflowdCollectors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := NetworkHierarchyCflowdCollectors{}
			item.VpnId = types.Int64Null()

			if t := v.Get("vpnId.optionType"); t.Exists() {
				va := v.Get("vpnId.value")
				if t.String() == "global" {
					item.VpnId = types.Int64Value(va.Int())
				}
			}
			item.IpAddress = types.StringNull()

			if t := v.Get("address.optionType"); t.Exists() {
				va := v.Get("address.value")
				if t.String() == "global" {
					item.IpAddress = types.StringValue(va.String())
				}
			}
			item.Port = types.Int64Null()

			if t := v.Get("udpPort.optionType"); t.Exists() {
				va := v.Get("udpPort.value")
				if t.String() == "global" {
					item.Port = types.Int64Value(va.Int())
				}
			}
			item.ExportSpreading = types.BoolNull()

			if t := v.Get("exportSpread.optionType"); t.Exists() {
				va := v.Get("exportSpread.value")
				if t.String() == "global" {
					item.ExportSpreading = types.BoolValue(va.Bool())
				}
			}
			item.BfdMetricsExporting = types.BoolNull()

			if t := v.Get("bfdMetricsExport.optionType"); t.Exists() {
				va := v.Get("bfdMetricsExport.value")
				if t.String() == "global" {
					item.BfdMetricsExporting = types.BoolValue(va.Bool())
				}
			}
			item.ExportingInterval = types.Int64Null()

			if t := v.Get("exportInterval.optionType"); t.Exists() {
				va := v.Get("exportInterval.value")
				if t.String() == "global" {
					item.ExportingInterval = types.Int64Value(va.Int())
				}
			}
			data.Collectors = append(data.Collectors, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *NetworkHierarchyCflowd) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringNull()
	data.Description = types.StringNull()
	path := "payload.data."
	data.ActiveFlowTimeout = types.Int64Null()

	if t := res.Get(path + "flowActiveTimeout.optionType"); t.Exists() {
		va := res.Get(path + "flowActiveTimeout.value")
		if t.String() == "global" {
			data.ActiveFlowTimeout = types.Int64Value(va.Int())
		}
	}
	data.InactiveFlowTimeout = types.Int64Null()

	if t := res.Get(path + "flowInactiveTimeout.optionType"); t.Exists() {
		va := res.Get(path + "flowInactiveTimeout.value")
		if t.String() == "global" {
			data.InactiveFlowTimeout = types.Int64Value(va.Int())
		}
	}
	data.FlowRefresh = types.Int64Null()

	if t := res.Get(path + "flowRefreshTime.optionType"); t.Exists() {
		va := res.Get(path + "flowRefreshTime.value")
		if t.String() == "global" {
			data.FlowRefresh = types.Int64Value(va.Int())
		}
	}
	data.SamplingInterval = types.Int64Null()

	if t := res.Get(path + "flowSamplingInterval.optionType"); t.Exists() {
		va := res.Get(path + "flowSamplingInterval.value")
		if t.String() == "global" {
			data.SamplingInterval = types.Int64Value(va.Int())
		}
	}
	data.CollectTlocLoopback = types.BoolNull()

	if t := res.Get(path + "collectTlocLoopback.optionType"); t.Exists() {
		va := res.Get(path + "collectTlocLoopback.value")
		if t.String() == "global" {
			data.CollectTlocLoopback = types.BoolValue(va.Bool())
		}
	}
	data.Protocol = types.StringNull()

	if t := res.Get(path + "protocol.optionType"); t.Exists() {
		va := res.Get(path + "protocol.value")
		if t.String() == "global" {
			data.Protocol = types.StringValue(va.String())
		}
	}
	data.Tos = types.BoolNull()

	if t := res.Get(path + "customizedIpv4RecordFields.collectTos.optionType"); t.Exists() {
		va := res.Get(path + "customizedIpv4RecordFields.collectTos.value")
		if t.String() == "global" {
			data.Tos = types.BoolValue(va.Bool())
		}
	}
	data.RemarkedDscp = types.BoolNull()

	if t := res.Get(path + "customizedIpv4RecordFields.collectDscpOutput.optionType"); t.Exists() {
		va := res.Get(path + "customizedIpv4RecordFields.collectDscpOutput.value")
		if t.String() == "global" {
			data.RemarkedDscp = types.BoolValue(va.Bool())
		}
	}
	for i := range data.Collectors {
		keys := [...]string{"vpnId"}
		keyValues := [...]string{strconv.FormatInt(data.Collectors[i].VpnId.ValueInt64(), 10)}
		keyValuesVariables := [...]string{""}

		var r gjson.Result
		res.Get(path + "collectors").ForEach(
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
		data.Collectors[i].VpnId = types.Int64Null()

		if t := r.Get("vpnId.optionType"); t.Exists() {
			va := r.Get("vpnId.value")
			if t.String() == "global" {
				data.Collectors[i].VpnId = types.Int64Value(va.Int())
			}
		}
		data.Collectors[i].IpAddress = types.StringNull()

		if t := r.Get("address.optionType"); t.Exists() {
			va := r.Get("address.value")
			if t.String() == "global" {
				data.Collectors[i].IpAddress = types.StringValue(va.String())
			}
		}
		data.Collectors[i].Port = types.Int64Null()

		if t := r.Get("udpPort.optionType"); t.Exists() {
			va := r.Get("udpPort.value")
			if t.String() == "global" {
				data.Collectors[i].Port = types.Int64Value(va.Int())
			}
		}
		data.Collectors[i].ExportSpreading = types.BoolNull()

		if t := r.Get("exportSpread.optionType"); t.Exists() {
			va := r.Get("exportSpread.value")
			if t.String() == "global" {
				data.Collectors[i].ExportSpreading = types.BoolValue(va.Bool())
			}
		}
		data.Collectors[i].BfdMetricsExporting = types.BoolNull()

		if t := r.Get("bfdMetricsExport.optionType"); t.Exists() {
			va := r.Get("bfdMetricsExport.value")
			if t.String() == "global" {
				data.Collectors[i].BfdMetricsExporting = types.BoolValue(va.Bool())
			}
		}
		data.Collectors[i].ExportingInterval = types.Int64Null()

		if t := r.Get("exportInterval.optionType"); t.Exists() {
			va := r.Get("exportInterval.value")
			if t.String() == "global" {
				data.Collectors[i].ExportingInterval = types.Int64Value(va.Int())
			}
		}
	}
}

// End of section. //template:end updateFromBody
