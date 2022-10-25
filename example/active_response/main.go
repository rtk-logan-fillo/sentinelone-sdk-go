package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone"
)

// contains the agent by disconnecting it from the network
func Contain(client *sentinelone.APIClient, name string) {
	fmt.Printf("containing agent '%s'\n", name)
	changeAgentNetworkStatus(client, name, "disconnected")
}

// uncontains the agent by reconnecting it to the network
func Uncontain(client *sentinelone.APIClient, name string) {
	fmt.Printf("uncontaining agent '%s'\n", name)
	changeAgentNetworkStatus(client, name, "connected")
}

// changes agent status to desired status
func changeAgentNetworkStatus(client *sentinelone.APIClient, name string, desiredStatus string) {
	agent := getAgent(client, name)
	id := agent.GetId()
	fmt.Printf("current network status '%s', desired network status '%s' \n", agent.GetNetworkStatus(), desiredStatus)
	var err error
	switch desiredStatus {
	case "connected":
		filter := sentinelone.NewAgentsSchemasAgentsActionSchemaFilter()
		filter.SetIds([]string{id})
		params := sentinelone.NewAgentsSchemasAgentsActionSchema(*filter)
		_, _, err = client.AgentActionsApi.ConnectAgents(context.Background()).Body(*params).Execute()
	case "disconnected":
		filter := sentinelone.NewAgentsSchemasAgentsDangerousActionSchemaFilter()
		filter.SetIds([]string{id})
		params := sentinelone.NewAgentsSchemasAgentsDangerousActionSchema(*filter)
		_, _, err = client.AgentActionsApi.DisconnectAgents(context.Background()).Body(*params).Execute()
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("presumably %s agent '%s', polling for confirmation\n", desiredStatus, name)
	poll(func() bool {
		agent := getAgent(client, name)
		actualStatus := agent.GetNetworkStatus()
		fmt.Printf("polling network status, expecting '%s', got '%s'\n", desiredStatus, actualStatus)
		return actualStatus == desiredStatus
	})
	fmt.Printf("succesfully %s agent '%s'\n", desiredStatus, name)
}

// gets the agent by computer name
func getAgent(client *sentinelone.APIClient, name string) sentinelone.AgentsSchemasAgentViewSchemaMany200DataInner {
	res, _, err := client.AgentsApi.GetAgents(context.Background()).ComputerName(name).Execute()
	if err != nil {
		panic(err)
	}
	if len(res.Data) <= 0 {
		panic(fmt.Sprintf("could not find agent %s", name))
	}
	agent := res.Data[0]
	return agent
}

type pollFn func() bool

// polls every one second, times out and returns error after 5 iterations
func poll(fn pollFn) {
	iter := 0
	iterMax := 10
	freq := 1 * time.Second
	for !fn() {
		time.Sleep(freq)
		iter++
		if iter == iterMax {
			panic(errors.New("poll timeout"))
		}
	}
}

// gets client config, uses env var for api token
func getConfig() *sentinelone.Configuration {
	token := os.Getenv("SENTINELONE_API_TOKEN")
	if token == "" {
		panic("you must run \"export SENTINELONE_API_TOKEN=<A valid API token>\"")
	}
	cfg := sentinelone.NewConfiguration()
	cfg.Host = "usea1-partners.sentinelone.net"
	cfg.UserAgent = "Arctic-Wolf/1.0"
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("APIToken %s", token))
	return cfg
}

// contain and uncontain the endpoint "Cherry"
func main() {
	client := sentinelone.NewAPIClient(getConfig())
	name := "Cherry"
	Contain(client, name)
	time.Sleep(5 * time.Second)
	fmt.Println("sleeping...")
	Uncontain(client, name)
}
