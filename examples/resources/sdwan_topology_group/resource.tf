resource "sdwan_topology_group" "example" {
  name = "TG_1"
  description = "My topology group 1"
  solution = "sdwan"
  profiles = [
    {
      id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
    }
  ]
}
