cluster_name = "vault"
log_level = "INFO"

ui = "true"

storage "consul" {
  address = "127.0.0.1:8500"
  path    = "vault"
}

listener "tcp" {
  address     = "127.0.0.1:8200"
  tls_disable = 1
}

telemetry {
    prometheus_retention_time = "5m"
    disable_hostname = true
}
