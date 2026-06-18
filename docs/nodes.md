# Nodes

Nodes are the machines — servers, containers, laptops — connected to your Headscale network. You can list them, inspect their details, register new ones, remove them, change their name, approve their routes, assign tags, and backfill IPs.

## Accessing the Resource

```go
nodes := client.Nodes()
```

## Operations

### List Nodes

Lists all nodes, optionally filtered by user.

```go
resp, err := client.Nodes().List(ctx, nodes.NodeListFilter{})
resp, err := client.Nodes().List(ctx, nodes.NodeListFilter{User: "myuser"})
```

Each node includes its name, IP addresses, online status, user, tags, and more.

### Get a Node

Fetch a single node by its ID.

```go
node, err := client.Nodes().Get(ctx, "node-id-123")
```

### Register a Node

Register a new node using a user and a pre-auth key.

```go
node, err := client.Nodes().Register(ctx, "myuser", "pre-auth-key-value")
```

### Delete a Node

Remove a node from the network by its ID.

```go
client.Nodes().Delete(ctx, "node-id-123")
```

### Expire a Node

Force a node to expire immediately.

```go
client.Nodes().Expire(ctx, "node-id-123")
```

### Rename a Node

Change a node's display name.

```go
node, err := client.Nodes().Rename(ctx, "node-id-123", "new-name")
```

### Approve Routes

Approve subnet routes advertised by a node. Routes are CIDR notation strings like `10.0.0.0/24`.

```go
node, err := client.Nodes().ApproveRoutes(ctx, "node-id-123", []string{"10.0.0.0/24", "192.168.1.0/24"})
```

### Add Tags

Assign ACL tags to a node. Tags must follow the `tag:` prefix convention.

```go
node, err := client.Nodes().AddTags(ctx, "node-id-123", []string{"tag:web", "tag:production"})
```

### Backfill IPs

Backfill IP address assignments for nodes. Pass `true` to confirm the operation.

```go
result, err := client.Nodes().BackfillIPs(ctx, true)
```

### Check if a Node is an Exit Node

The `Node` type has a helper method that returns `true` if the node's approved routes include `0.0.0.0/0` or `::/0`.

```go
n, _ := client.Nodes().Get(ctx, "node-id-123")
if n.Node.IsExitNode() {
    fmt.Println("this node routes all traffic")
}
```

## Types

**Node** — the main entity. Key fields:

| Field             | Type                      | Description                                |
| ----------------- | ------------------------- | ------------------------------------------ |
| `ID`              | `string`                  | Unique identifier                          |
| `Name`            | `string`                  | Node's hostname                            |
| `GivenName`       | `string`                  | Display name (can be renamed)              |
| `IPAddresses`     | `[]string`                | Assigned Tailscale IPs                     |
| `User`            | `users.User`              | The user who owns this node                |
| `Online`          | `bool`                    | Whether the node is currently connected    |
| `Tags`            | `[]string`                | ACL tags assigned to the node              |
| `ApprovedRoutes`  | `[]string`                | CIDR routes this node can advertise        |
| `AvailableRoutes` | `[]string`                | Routes the node is advertising             |
| `LastSeen`        | `time.Time`               | Last contact with the control server       |
| `Expiry`          | `time.Time`               | When the node's key expires                |
| `PreAuthKey`      | `*preauthkeys.PreAuthKey` | The pre-auth key used to register (if any) |

**Request types:**

- `NodeListFilter` — optional `User` field to filter by owner.
- `ApproveRoutesRequest` — contains `Routes []string`.
- `AddTagsRequest` — contains `Tags []string`.

**Response types:**

- `NodesResponse` — wraps `[]Node` (returned by List).
- `NodeResponse` — wraps a single `Node` (returned by Get, Register, Rename, ApproveRoutes, AddTags).
- `BackfillIPsResponse` — wraps `Changes []string` (returned by BackfillIPs).
