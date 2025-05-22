package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/hibare/headscale-client-go/logger"
	"github.com/hibare/headscale-client-go/utils"
	hsClient "github.com/hibare/headscale-client-go/v1/client"
	"github.com/hibare/headscale-client-go/v1/nodes"
)

var hsClientNewClient = hsClient.NewClient

func listNodes(client hsClient.ClientInterface) (string, error) {
	ns, err := client.Nodes().List(context.Background(), nodes.NodeListFilter{})
	if err != nil {
		return "", err
	}
	result := ""
	for _, node := range ns.Nodes {
		result += fmt.Sprintf("Node: %+v\n", node)
	}
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

	fmt.Println("Listing Nodes")
	output, err := listNodes(client)
	if err != nil {
		panic(err)
	}
	fmt.Print(output)
}
