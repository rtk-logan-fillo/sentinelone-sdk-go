package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone"
)

// gets client config, uses env var for api token
func getConfig() *sentinelone.Configuration {
	token := os.Getenv("SENTINELONE_API_TOKEN")
	if token == "" {
		panic("you must run \"export SENTINELONE_API_TOKEN=<A valid API token>\"")
	}
	cfg := sentinelone.NewConfiguration()
	cfg.Host = "usea1-partners.sentinelone.net"
	cfg.UserAgent = "ArcticWolf"
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("APIToken %s", token))
	return cfg
}

// get endpoint agent "Cherry" and print various data
func main() {
	client := sentinelone.NewAPIClient(getConfig())
	res, _, err := client.AgentsApi.GetAgents(context.Background()).ComputerName("Cherry").Execute()
	if err != nil {
		panic(err)
	}
	data := res.Data
	for _, endpoint := range data {
		msg := fmt.Sprintf(
			"got endpoint '%s (%s)' with network status '%s' and '%d' active threats (infected: %t, quarantined: %t)",
			endpoint.GetComputerName(),
			endpoint.GetExternalIp(),
			endpoint.GetNetworkStatus(),
			endpoint.GetActiveThreats(),
			endpoint.GetInfected(),
			endpoint.GetNetworkQuarantineEnabled())
		fmt.Println(msg)
	}
}
