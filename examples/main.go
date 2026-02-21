// Package main provides an example usage of the headscale-client-go library.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/hibare/headscale-client-go/logger"
	"github.com/hibare/headscale-client-go/utils"
	hsClient "github.com/hibare/headscale-client-go/v1/client"
	"github.com/hibare/headscale-client-go/v1/nodes"
)

var hsClientNewClient = hsClient.NewClient
var stdout = os.Stdout

func listNodes(client hsClient.ClientInterface) (string, error) {
	ns, err := client.Nodes().List(context.Background(), nodes.NodeListFilter{})
	if err != nil {
		return "", err
	}
	result := ""
	var resultSb23 strings.Builder
	for _, node := range ns.Nodes {
		fmt.Fprintf(&resultSb23, "Node: %+v\n", node)
	}
	result += resultSb23.String()
	return result, nil
}

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	clientURL := os.Getenv("HS_SERVER_URL")
	clientToken := os.Getenv("HS_SERVER_TOKEN")

	client, err := hsClientNewClient(clientURL, clientToken, hsClient.ClientOptions{
		LogLevel: utils.ToPtr(logger.LevelDebug), // Change to LevelInfo for less verbose output
	})
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintln(stdout, "Listing Nodes")
	output, err := listNodes(client)
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprint(stdout, output)
}
