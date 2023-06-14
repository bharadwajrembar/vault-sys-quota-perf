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
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/secret/data/cloudhub-kv/01fdf72b-3302-4e51-8fa0-37f89d3e77ee", BaseURL),
			Body:   []byte(""),
		},
		{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/secret/data/app/common-cg-service-multi-cluster", BaseURL),
			Body:   []byte(""),
		},
		{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/secret/data/app/common-core-service", BaseURL),
			Body:   []byte(""),
		},
		{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/secret/data/app/common-auth-flow-service", BaseURL),
			Body:   []byte(""),
		},
		{
			Method: "POST",
			URL:    fmt.Sprintf("%s/v1/secret/data/vault-perf-test", BaseURL),
			Body:   []byte(KvData),
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
