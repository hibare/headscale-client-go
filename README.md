# Headscale Client Go

[![Go Reference](https://pkg.go.dev/badge/github.com/tailscale/tailscale-client-go/v2.svg)](https://pkg.go.dev/github.com/hibare/headscale-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/hibare/headscale-client-go)](https://goreportcard.com/report/github.com/hibare/headscale-client-go)

A client implementation for the [Headscale](https://headscale.net) HTTP API.

## Example (Using API Key)

```go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	hsClient "github.com/hibare/headscale-client-go"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug) // Optional

	serverUrl := os.Getenv("HS_SERVER_URL")
	apiToken := os.Getenv("HS_API_TOKEN")

	client, err := hsClient.NewClient(serverUrl, apiToken, hsClient.HeadscaleClientOptions{})
	if err != nil {
		panic(err)
	}

	nodes, err := client.Nodes().List(context.Background())
	if err != nil {
		panic(err)
	}

	for _, node := range nodes.Nodes {
		fmt.Printf("Node: %s\n", node.Name)
	}

}

```
