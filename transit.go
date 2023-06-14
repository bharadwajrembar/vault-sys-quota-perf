package main

import (
	"fmt"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"time"
)

var (
	transitPerfMetrics  vegeta.Metrics
	vaultTransitTargets = []Endpoints{
		{
			Method: "POST",
			URL:    fmt.Sprintf("%s/v1/transit/encrypt/test", BaseURL),
			Body:   []byte(transitEncryptData),
		},
		{
			Method: "POST",
			URL:    fmt.Sprintf("%s/v1/transit/decrypt/test", BaseURL),
			Body:   []byte(transitDecryptData),
		},
	}
)

func transitPerf(rate vegeta.Rate, duration time.Duration) *vegeta.Metrics {

	for _, vaultTarget := range vaultTransitTargets {
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: vaultTarget.Method,
			URL:    vaultTarget.URL,
			Header: vaultHeader,
			Body:   vaultTarget.Body,
		})
		attacker := vegeta.NewAttacker()

		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			fmt.Println(res.URL, res.Code, res.Error)
			transitPerfMetrics.Add(res)
		}
		transitPerfMetrics.Close()
	}

	return &transitPerfMetrics
}