<div align="center">
  <img src="./assets/logo.png" alt="headscale-client-go Logo" width="200" height="200">

# Headscale Client Go

_A Go client for the [Headscale](https://headscale.net) HTTP API._

[![Go Reference](https://pkg.go.dev/badge/github.com/hibare/headscale-client-go.svg)](https://pkg.go.dev/github.com/hibare/headscale-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/hibare/headscale-client-go)](https://goreportcard.com/report/github.com/hibare/headscale-client-go)

</div>

## Features

- **API Keys**: List, create, expire, and delete API keys (by prefix or ID).
- **Nodes**: List, get, register, delete, expire, rename, tag, and backfill IPs.
- **Users**: List, create, delete, and rename users.
- **Policies**: Get and update policy documents.
- **Pre-Auth Keys**: List, create, expire, and delete pre-auth keys.
- Customizable HTTP client, user agent, and logger support.
- Idiomatic Go API with context support.

---

## Requirements

- **Go**: 1.26+
- **Headscale**: v0.28.0+

---

## Installation

```sh
go get github.com/hibare/headscale-client-go
```

---

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
    // Create client
    client, err := hsClient.NewClient(
        "http://headscale:8080",
        "your-api-key",
        hsClient.ClientOptions{},
    )
    if err != nil {
        panic(err)
    }

    // List nodes
    nodeList, err := client.Nodes().List(context.Background(), nodes.NodeListFilter{})
    if err != nil {
        panic(err)
    }

    for _, node := range nodeList.Nodes {
        fmt.Printf("Node: %s (%s)\n", node.Name, node.ID)
    }
}
```

---

## API Overview

| Resource               | Description          |
| ---------------------- | -------------------- |
| `client.APIKeys()`     | Manage API keys      |
| `client.Nodes()`       | Manage nodes         |
| `client.Users()`       | Manage users         |
| `client.Policy()`      | Manage policy        |
| `client.PreAuthKeys()` | Manage pre-auth keys |

For full API documentation, see [pkg.go.dev](https://pkg.go.dev/github.com/hibare/headscale-client-go).

---

## Examples

See the [`examples/`](examples/) directory for more usage examples.

---

## Development

```sh
# Run tests
make test

# Run E2E tests (requires Docker)
make e2e-test

# Lint
golangci-lint run

# Format
go fmt ./...
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

---

## License

MIT License. See [LICENSE](LICENSE) for details.

---

## Links

- [Headscale](https://headscale.net)
- [pkg.go.dev Documentation](https://pkg.go.dev/github.com/hibare/headscale-client-go)
- [GitHub Issues](https://github.com/hibare/headscale-client-go/issues)
