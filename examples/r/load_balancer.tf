resource "sakuracloud_load_balancer" "foobar" {
  name         = "foobar"
  switch_id    = sakuracloud_switch.foobar.id
  vrid         = 1
  ip_addresses = ["192.168.11.101"]
  netmask      = 24
  gateway      = "192.168.11.1"

  description = "description"
  tags        = ["tag1", "tag2"]

  vip {
    vip          = "192.168.11.201"
    port         = 80
    delay_loop   = 10
    sorry_server = "192.168.11.21"

    server {
      ip_address     = "192.168.11.51"
      check_protocol = "http"
      check_path     = "/health"
      check_status   = 200
    }

    server {
      ip_address     = "192.168.11.52"
      check_protocol = "http"
      check_path     = "/health"
      check_status   = 200
    }
  }
}

resource "sakuracloud_switch" "foobar" {
  name = "foobar"
}