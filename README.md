<div align="center">
  <img src="./assets/logo.png" alt="headscale-client-go Logo" width="200" height="200">

# Headscale Client Go

_A Go client for the [Headscale](https://headscale.net) HTTP API._

[![Go Reference](https://pkg.go.dev/badge/github.com/hibare/headscale-client-go.svg)](https://pkg.go.dev/github.com/hibare/headscale-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/hibare/headscale-client-go)](https://goreportcard.com/report/github.com/hibare/headscale-client-go)

</div>

A Go client library for the Headscale HTTP API — manage users, nodes, API keys, pre-auth keys, and ACL policies.

## Requirements

Go 1.26+, Headscale v0.28.0+

## Installation

```
go get github.com/hibare/headscale-client-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    hsClient "github.com/hibare/headscale-client-go/v1/client"
    "github.com/hibare/headscale-client-go/v1/nodes"
)

func main() {
    client, err := hsClient.NewClient("http://headscale:8080", "your-api-key", hsClient.ClientOptions{})
    if err != nil {
        panic(err)
    }
    nodeList, err := client.Nodes().List(context.Background(), nodes.NodeListFilter{})
    if err != nil {
        panic(err)
    }
    for _, node := range nodeList.Nodes {
        fmt.Printf("Node: %s (%s)\n", node.Name, node.ID)
    }
}
```

## Documentation

| Resource                                  | What it covers                                   |
| ----------------------------------------- | ------------------------------------------------ |
| [Setup & Customization](docs/overview.md) | Install, client setup, options, error handling   |
| [API Keys](docs/apikeys.md)               | Create, list, expire, delete API keys            |
| [Nodes](docs/nodes.md)                    | List, get, register, rename, tag, approve routes |
| [Users](docs/users.md)                    | List, create, rename, delete users               |
| [Policy](docs/policy.md)                  | Read and update ACL documents                    |
| [Pre-Auth Keys](docs/preauthkeys.md)      | Create, list, expire, delete pre-auth keys       |

## Development

```bash
make test         # unit tests
make e2e-test     # E2E tests (requires Docker)
golangci-lint run
go fmt ./...
```

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE).
