resource "sdwan_network_hierarchy_security_logging_feature" "example" {
	name = "Example"
	description = "My Example"
  network_hierarchy_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"
  high_speed_logging = [
    {
      name = "server1"
      vrf = "VPN10"
      server_ip = "10.1.200.253"
      port = 2055
    }
  ]
  utd_syslog = [
    {
      vpn = "VPN10"
      server_ip = "10.1.200.253"
    }
  ]
}
