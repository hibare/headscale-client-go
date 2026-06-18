# Getting Started

This guide walks through installing the library, creating a client, customizing it, and handling errors.

## Installation

```sh
go get github.com/hibare/headscale-client-go
```

## Creating a Client

You need two things to connect: your Headscale server URL and an API key.

```go
import hsClient "github.com/hibare/headscale-client-go/v1/client"

client, err := hsClient.NewClient("https://headscale.example.com", "your-api-key", hsClient.ClientOptions{})
```

`NewClient` returns a `ClientInterface` value. The zero `ClientOptions{}` gives you sensible defaults:
a 1-minute HTTP timeout, user-agent `headscale-client-go`, and JSON-structured info-level logging.

## Customizing the Client

Pass options to configure how the client behaves:

```go
type ClientOptions struct {
    HTTPClient *http.Client   // custom HTTP client (timeout, transport, etc.)
    UserAgent  *string        // custom User-Agent header
    Logger     logger.Logger  // custom logger implementation
    LogLevel   *logger.LogLevel // log verbosity (ignored if Logger is set)
}
```

**Custom timeout:**

```go
opt := hsClient.ClientOptions{
    HTTPClient: &http.Client{Timeout: 30 * time.Second},
}
```

**Custom user agent:**

```go
opt := hsClient.ClientOptions{
    UserAgent: utils.ToPtr("my-application/1.0.0"),
}
```

**Change log level:**

Levels: `LevelDebug`, `LevelInfo` (default), `LevelWarn`, `LevelError`.

```go
opt := hsClient.ClientOptions{
    LogLevel: utils.ToPtr(logger.LevelDebug),
}
```

**Custom logger:**

Implement the `Logger` interface (Info, Error, Warn, Debug methods) and pass it in:

```go
opt := hsClient.ClientOptions{
    Logger: myCustomLogger,
}
```

## Using Resources

Once you have a client, resource methods give you access to different parts of the Headscale API:

| Method                 | What it manages                                        |
| ---------------------- | ------------------------------------------------------ |
| `client.APIKeys()`     | API keys (create, list, expire, delete)                |
| `client.Nodes()`       | Network nodes (list, get, register, etc.)              |
| `client.Users()`       | User accounts (list, create, rename, delete)           |
| `client.Policy()`      | ACL policy document (get, update)                      |
| `client.PreAuthKeys()` | Pre-authentication keys (create, list, expire, delete) |

Each resource is documented in its own page (see the links in the README).

## Error Handling

When the Headscale API returns a non-2xx status, the client returns a typed error you can inspect:

```go
type APIError struct {
    StatusCode int    // e.g. 400, 401, 403, 500
    Body       string // raw response body from the API
}
```

Use `errors.As` to check for it:

```go
var apiErr *requests.APIError
if errors.As(err, &apiErr) {
    fmt.Printf("API error: %d — %s\n", apiErr.StatusCode, apiErr.Body)
}
```

Non-API errors (network timeouts, DNS failures) are returned as standard Go errors and won't match `*APIError`.
