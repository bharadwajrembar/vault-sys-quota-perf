# Setup Consul

```
$ git clone git@github.com:hashicorp/consul.git
$ cd consul
```

## Checkout to tag

```
$ git checkout <tag>
```

*Example:*

```
$ git checkout v1.9.1
```

## Fetch dependencies

```
$ go mod tidy
```

## Build the binary

```
$ go build
```

## Move binary

```
$ sudo mv consul /usr/local/bin/
```

## Verify

```
$ consul -v
Consul v1.9.1
Protocol 2 spoken by default, understands 2 to 3 (agent will automatically use protocol >2 when speaking to compatible agents)
```

# Start Consul

```
$ cd vault-quota
$ consul agent -server -bootstrap -ui -data-dir=/tmp/consul
$ consul agent -config-file=consul.hcl
```

## Open Consul IU
Navigate to http://localhost:8500