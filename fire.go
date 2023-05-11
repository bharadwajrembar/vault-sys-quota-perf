package main

import (
	"fmt"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const (
	BaseURL              = "http://localhost:8200"
	X_VAULT_TOKEN_HEADER = "X-VAULT-TOKEN"
)

type Endpoints struct {
	Method  string
	URL     string
	Headers http.Header
}

var (
	vaultToken   = "<vault-token>"
	vault_header = http.Header{
		X_VAULT_TOKEN_HEADER: {vaultToken},
	}
	vaultTargets = []Endpoints{
		{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/secret/data/app/jaw-service", BaseURL),
		},
	}
)

func main() {
	rate := vegeta.Rate{Freq: 20, Per: 10 * time.Second}
	duration := 10 * time.Second
	for _, vaultTarget := range vaultTargets {
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: vaultTarget.Method,
			URL:    vaultTarget.URL,
			Header: vault_header,
		})
		attacker := vegeta.NewAttacker()

		var metrics vegeta.Metrics
		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			fmt.Println(res.URL, res.Code, res.Error, res.Latency)
			metrics.Add(res)
		}
		metrics.Close()

		fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	}

}
