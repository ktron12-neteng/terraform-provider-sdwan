resource "sdwan_topology_custom_control_feature" "example" {
	name = "Example"
	description = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  default_action = "reject"
  target_level = "SITE"
  target_inbound_site_ids = ["e00b80d7-826f-40e1-a7fe-7a542d42e059"]
  target_outbound_site_ids = ["e00b80d7-826f-40e1-a7fe-7a542d42e059"]
  sequences = [
    {
      id = 10
      name = "Backup_TLOC"
      base_action = "accept"
      type = "route"
      ip_type = "ipv4"
      match_entries = [
        {
          site_ids = ["e00b80d7-826f-40e1-a7fe-7a542d42e059"]
        }
      ]
      actions = [
        {
          set = [
            {
              preference = 5000
              tloc_list_id = "33923a1b-0e38-49dd-b833-1ff47661311d"
              tloc_action = "backup"
            }
          ]
        }
      ]
    }
  ]
}
