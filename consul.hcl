node_name = "consul-server"

server = true

data_dir = "/tmp/consul"

bootstrap = true

log_level  = "INFO"

datacenter = "dc1"

ui_config {
  enabled = true
}

addresses {
  http = "0.0.0.0"
}
