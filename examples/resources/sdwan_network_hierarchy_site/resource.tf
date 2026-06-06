resource "sdwan_network_hierarchy_site" "example" {
  name = "DC1"
  description = "My site"
  parent_id = "3c2b8678-6fe5-42b1-a01d-c1cdf7ca7999"
  site_id = 100
}
