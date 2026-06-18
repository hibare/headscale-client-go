# API Keys

API keys authenticate your client when talking to Headscale. Each key has a prefix you can use to identify it, an expiration date, and a last-seen timestamp. You can create, list, expire, and delete them.

## Accessing the Resource

```go
keys := client.APIKeys()
```

## Operations

### List All Keys

Returns every API key and its metadata.

```go
resp, err := client.APIKeys().List(ctx)
for _, key := range resp.APIKeys {
    fmt.Printf("%s (%s)\n", key.Prefix, key.ID)
}
```

### Create a Key

Generate a new API key with an expiration date.

```go
resp, err := client.APIKeys().Create(ctx, apikeys.CreateAPIKeyRequest{
    Expiration: time.Now().Add(90 * 24 * time.Hour),
})
```

The `resp.APIKey` field contains the full token string (e.g. `hskey-api-...`). **Save it immediately** — this is the only time it's returned.

### Expire a Key

Revoke a key by its prefix or by its numeric ID.

```go
client.APIKeys().Expire(ctx, "abcde")    // by prefix
client.APIKeys().ExpireByID(ctx, "1")    // by ID
```

### Delete a Key

Remove a key by its prefix.

```go
client.APIKeys().Delete(ctx, "abcde")
```

## Types

APIKey represents a key's metadata:

| Field        | Type        | Description                                        |
| ------------ | ----------- | -------------------------------------------------- |
| `ID`         | `string`    | Numeric identifier assigned by Headscale           |
| `Prefix`     | `string`    | First characters of the token (identifies the key) |
| `Expiration` | `time.Time` | When this key stops working                        |
| `CreatedAt`  | `time.Time` | When the key was created                           |
| `LastSeen`   | `time.Time` | Last time this key was used                        |

**Request types:**

- `CreateAPIKeyRequest` — accepts `Expiration` (the only required field).
- `ExpireAPIKeyRequest` — accepts either `Prefix` or `ID` (used internally).

**Response types:**

- `APIKeysResponse` — wraps `[]APIKey` (returned by List).
- `CreateAPIKeyResponse` — wraps `APIKey string` (the full token, returned once by Create).
