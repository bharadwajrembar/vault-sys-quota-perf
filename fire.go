package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const (
	BaseURL              = "http://localhost:8204"
	X_VAULT_TOKEN_HEADER = "X-VAULT-TOKEN"
)

type Endpoints struct {
	Method  string
	URL     string
	Headers http.Header
	Body    []byte
}

var (
	vaultToken  = os.Getenv("VAULT_TOKEN")
	vaultHeader = http.Header{
		X_VAULT_TOKEN_HEADER: {vaultToken},
	}
)

func main() {
	rate := vegeta.Rate{Freq: 20, Per: 1 * time.Second}
	duration := 10 * time.Second

	kvMetrics := kvPerf(rate, duration)
	if len(kvMetrics.Errors) != 0 {
		fmt.Printf("Errors when attacking: %s\n", kvMetrics.Errors)
	}

	fmt.Printf("Success for KV calls: %f\n", kvMetrics.Success*100)
	fmt.Printf("90th percentile for KV calls: %s\n", kvMetrics.Latencies.P90)
	fmt.Printf("95th percentile for KV calls: %s\n", kvMetrics.Latencies.P95)
	fmt.Printf("99th percentile for KV calls: %s\n", kvMetrics.Latencies.P99)

}
