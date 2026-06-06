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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &NetworkHierarchyCflowdProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkHierarchyCflowdProfileParcelDataSource{}
)

func NewNetworkHierarchyCflowdProfileParcelDataSource() datasource.DataSource {
	return &NetworkHierarchyCflowdProfileParcelDataSource{}
}

type NetworkHierarchyCflowdProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *NetworkHierarchyCflowdProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network_hierarchy_cflowd_feature"
}

func (d *NetworkHierarchyCflowdProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Network Hierarchy Cflowd Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Computed:            true,
			},
			"network_hierarchy_id": schema.StringAttribute{
				MarkdownDescription: "The ID of the network hierarchy node (the Global node) the cflowd applies to",
				Required:            true,
			},
			"active_flow_timeout": schema.Int64Attribute{
				MarkdownDescription: "Active Flow Timeout(Seconds)",
				Computed:            true,
			},
			"inactive_flow_timeout": schema.Int64Attribute{
				MarkdownDescription: "Inactive Flow Timeout(Seconds)",
				Computed:            true,
			},
			"flow_refresh": schema.Int64Attribute{
				MarkdownDescription: "Flow Refresh Time(Seconds)",
				Computed:            true,
			},
			"sampling_interval": schema.Int64Attribute{
				MarkdownDescription: "Sampling Interval(Seconds)",
				Computed:            true,
			},
			"collect_tloc_loopback": schema.BoolAttribute{
				MarkdownDescription: "Collect SDWAN TLOC loopback interface name instead of physical",
				Computed:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: "FNF Protocol",
				Computed:            true,
			},
			"tos": schema.BoolAttribute{
				MarkdownDescription: "Collect TOS",
				Computed:            true,
			},
			"remarked_dscp": schema.BoolAttribute{
				MarkdownDescription: "Collect Re-marked DSCP",
				Computed:            true,
			},
			"collectors": schema.ListNestedAttribute{
				MarkdownDescription: "Collectors list",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"vpn_id": schema.Int64Attribute{
							MarkdownDescription: "VPN ID",
							Computed:            true,
						},
						"ip_address": schema.StringAttribute{
							MarkdownDescription: "Collector IPv4/IPv6 address",
							Computed:            true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: "Collector UDP port number,RFC recommended is 4739",
							Computed:            true,
						},
						"export_spreading": schema.BoolAttribute{
							MarkdownDescription: "Export spreading enable",
							Computed:            true,
						},
						"bfd_metrics_exporting": schema.BoolAttribute{
							MarkdownDescription: "BFD Metric Exporting enable knob",
							Computed:            true,
						},
						"exporting_interval": schema.Int64Attribute{
							MarkdownDescription: "BFD Export Interval(Seconds)",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworkHierarchyCflowdProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *NetworkHierarchyCflowdProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetworkHierarchyCflowd

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
