package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"time"
)

var (
	kvPerfMetrics  vegeta.Metrics
	vaultKVTargets = []Endpoints{
		{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/secret/data/test", BaseURL),
			Body:   []byte(""),
		},
		{
			Method: "POST",
			URL:    fmt.Sprintf("%s/v1/secret/data/vault-perf-test", BaseURL),
			Body:   []byte(`{"data":{"test_1": "test"}}`),
		},
	}
)

func kvPerf(rate vegeta.Rate, duration time.Duration) *vegeta.Metrics {

	for _, vaultTarget := range vaultKVTargets {
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: vaultTarget.Method,
			URL:    vaultTarget.URL,
			Header: vaultHeader,
			Body:   vaultTarget.Body,
		})
		attacker := vegeta.NewAttacker()

		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			fmt.Println(res.URL, res.Code, res.Error)
			kvPerfMetrics.Add(res)
		}
		kvPerfMetrics.Close()
	}

	return &kvPerfMetrics
}
