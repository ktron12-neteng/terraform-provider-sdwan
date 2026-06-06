resource "sdwan_network_hierarchy_cflowd_feature" "example" {
	name = "Example"
	description = "My Example"
  network_hierarchy_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"
  active_flow_timeout = 600
  inactive_flow_timeout = 60
  flow_refresh = 600
  sampling_interval = 1
  collect_tloc_loopback = false
  protocol = "ipv4"
  tos = false
  remarked_dscp = false
  collectors = [
    {
      vpn_id = 10
      ip_address = "10.10.101.2"
      port = 4739
      export_spreading = false
      bfd_metrics_exporting = false
      exporting_interval = 60
    }
  ]
}
