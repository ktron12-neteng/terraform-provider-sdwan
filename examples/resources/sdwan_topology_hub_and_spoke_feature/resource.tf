resource "sdwan_topology_hub_and_spoke_feature" "example" {
	name = "Example"
	description = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  vpns = ["VPN10"]
  hubs = ["dab0b0a9-32bb-4c1b-a32b-3472c1b1c992"]
  spokes = [
    {
      name = "Spokes"
      site_ids = ["ad8729e6-d454-4008-ab1f-1fc08d82706e"]
      hubs = [
        {
          site_ids = ["dab0b0a9-32bb-4c1b-a32b-3472c1b1c992"]
          preference = 10
        }
      ]
    }
  ]
}
