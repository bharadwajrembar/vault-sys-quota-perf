package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

const (
	X_VAULT_TOKEN_HEADER = "X-VAULT-TOKEN"
	transitEncryptData   = `{"plaintext":"dmF1bHQgcGVyZiB0ZXN0Cg=="}`
	transitDecryptData   = `{"ciphertext":"vault:v1:/iAYV66W13DP7Zu6JwYpDbbo1rGh6iZz5aRJHDIz9zO8+mNemQXQBNtQ+g8="}`
	KvData               = `{"data":{"key1": "value1", "key2":"value2","key3":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"}}`
)

type Endpoints struct {
	Method  string
	URL     string
	Headers http.Header
	Body    []byte
}

var (
	BaseURL     = os.Getenv("VAULT_URL")
	vaultToken  = os.Getenv("VAULT_TOKEN")
	vaultHeader = http.Header{
		X_VAULT_TOKEN_HEADER: {vaultToken},
	}
)

func main() {
	rate := vegeta.Rate{Freq: 400, Per: 1 * time.Second}
	duration := 10 * time.Second

	// KV perf metrics
	kvMetrics := kvPerf(rate, duration)
	if len(kvMetrics.Errors) != 0 {
		fmt.Printf("Errors when attacking: %s\n", kvMetrics.Errors)
	}

	fmt.Printf("Success for KV calls: %f\n", kvMetrics.Success*100)
	fmt.Printf("90th percentile for KV calls: %s\n", kvMetrics.Latencies.P90)
	fmt.Printf("95th percentile for KV calls: %s\n", kvMetrics.Latencies.P95)
	fmt.Printf("99th percentile for KV calls: %s\n", kvMetrics.Latencies.P99)

	// Transit perf metrics
	transitMetrics := transitPerf(rate, duration)
	if len(transitMetrics.Errors) != 0 {
		fmt.Printf("Errors when attacking: %s\n", transitMetrics.Errors)
	}

	fmt.Printf("Success for transit calls: %f\n", transitMetrics.Success*100)
	fmt.Printf("90th percentile for transit calls: %s\n", transitMetrics.Latencies.P90)
	fmt.Printf("95th percentile for transit calls: %s\n", transitMetrics.Latencies.P95)
	fmt.Printf("99th percentile for transit calls: %s\n", transitMetrics.Latencies.P99)

	// Policy perf metrics
	policyMetrics := policyPerf(rate, duration)
	if len(policyMetrics.Errors) != 0 {
		fmt.Printf("Errors when attacking: %s\n", policyMetrics.Errors)
	}

	fmt.Printf("Success for policy calls: %f\n", policyMetrics.Success*100)
	fmt.Printf("90th percentile for policy calls: %s\n", policyMetrics.Latencies.P90)
	fmt.Printf("95th percentile for policy calls: %s\n", policyMetrics.Latencies.P95)
	fmt.Printf("99th percentile for policy calls: %s\n", policyMetrics.Latencies.P99)

}
