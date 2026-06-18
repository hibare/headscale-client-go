# Policy

The policy resource lets you read and update the ACL (access control list) document that controls traffic between nodes on your Headscale network. The document uses the same JSON/HuJSON format as Tailscale ACLs.

## Accessing the Resource

```go
policy := client.Policy()
```

## Operations

### Get Current Policy

Returns the current ACL document and when it was last updated.

```go
policy, err := client.Policy().Get(ctx)
fmt.Printf("Policy:\n%s\n", policy.Policy)
```

### Update Policy

Replace the entire ACL document with a new one.

```go
newPolicy := `{
    "acls": [
        {"action": "accept", "src": ["*"], "dst": ["*:*"]}
    ]
}`

resp, err := client.Policy().Update(ctx, newPolicy)
fmt.Printf("Updated at: %s\n", resp.UpdatedAt)
```

The policy string accepts any valid Headscale ACL document.

## Types

**Policy** — the current ACL state:

| Field       | Type     | Description                       |
| ----------- | -------- | --------------------------------- |
| `Policy`    | `string` | Raw ACL document (JSON or HuJSON) |
| `UpdatedAt` | `string` | ISO 8601 timestamp of last update |

**Request types:**

- `UpdatePolicyRequest` — contains `Policy string` (the full ACL document).

**Response types:**

- `UpdatePolicyResponse` — contains `Policy string` and `UpdatedAt string` (returned by Update).
