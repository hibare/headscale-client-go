# Pre-Auth Keys

Pre-authentication keys (pre-auth keys) let nodes join your Headscale network without needing manual approval from the admin. You control each key's lifetime, whether it can be reused, whether the node is ephemeral, and what ACL tags it assigns.

## Accessing the Resource

```go
keys := client.PreAuthKeys()
```

## Operations

### List All Keys

Returns every pre-auth key and its properties.

```go
resp, err := client.PreAuthKeys().List(ctx)
for _, key := range resp.PreAuthKeys {
    fmt.Printf("%s — reusable: %v, expires: %v\n", key.ID, key.Reusable, key.Expiration)
}
```

### Create a Key

Generate a new pre-auth key. You must specify which user it belongs to.

```go
key, err := client.PreAuthKeys().Create(ctx, preauthkeys.CreatePreAuthKeyRequest{
    User:       "myuser",
    Reusable:   true,         // allow multiple nodes to use the same key
    Ephemeral:  false,        // if true, nodes are removed on disconnect
    Expiration: time.Now().Add(24 * time.Hour),
    ACLTags:    []string{"tag:web", "tag:dev"},
})
fmt.Println(key.PreAuthKey.Key)
```

The `Key` field contains the token that nodes use when joining. The other fields (ID, tags, expiration, etc.) are also available in the response.

### Expire a Key

Revoke a key immediately by its ID.

```go
client.PreAuthKeys().Expire(ctx, "key-id-123")
```

### Delete a Key

Remove a key by its ID.

```go
client.PreAuthKeys().Delete(ctx, "key-id-123")
```

## Types

**PreAuthKey** — the main entity:

| Field        | Type         | Description                                       |
| ------------ | ------------ | ------------------------------------------------- |
| `ID`         | `string`     | Unique identifier                                 |
| `Key`        | `string`     | The token value (only returned on create)         |
| `User`       | `users.User` | The user this key belongs to                      |
| `Reusable`   | `bool`       | Whether the key can be used more than once        |
| `Ephemeral`  | `bool`       | Whether nodes created with this key are temporary |
| `Used`       | `bool`       | Whether the key has been used                     |
| `Expiration` | `time.Time`  | When the key expires                              |
| `CreatedAt`  | `time.Time`  | When the key was created                          |
| `ACLTags`    | `[]string`   | Tags automatically assigned to joining nodes      |

**Request types:**

- `CreatePreAuthKeyRequest` — accepts `User` (required), `Reusable`, `Ephemeral`, `Expiration`, `ACLTags`.
- `ExpirePreAuthKeyRequest` — accepts `ID` (used internally).

**Response types:**

- `PreAuthKeysResponse` — wraps `[]PreAuthKey` (returned by List).
- `PreAuthKeyResponse` — wraps a single `PreAuthKey` (returned by Create).
