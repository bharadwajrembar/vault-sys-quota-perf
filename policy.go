package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"time"
)

var (
	policyPerfMetrics  vegeta.Metrics
	vaultPolicyTargets = []Endpoints{
		{
			Method: "GET",
			URL:    fmt.Sprintf("%s/v1/sys/policy/read-only", BaseURL),
			Body:   []byte(""),
		},
	}
)

func policyPerf(rate vegeta.Rate, duration time.Duration) *vegeta.Metrics {

	for _, vaultTarget := range vaultPolicyTargets {
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: vaultTarget.Method,
			URL:    vaultTarget.URL,
			Header: vaultHeader,
			Body:   vaultTarget.Body,
		})
		attacker := vegeta.NewAttacker()

		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			fmt.Println(res.URL, res.Code, res.Error, res.Body)
			policyPerfMetrics.Add(res)
		}
		policyPerfMetrics.Close()
	}

	return &policyPerfMetrics
}
