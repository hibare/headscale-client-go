# Users

Users represent people or service identities in Headscale. Each node belongs to a user. You can list, create, rename, and delete them.

## Accessing the Resource

```go
users := client.Users()
```

## Operations

### List Users

Lists all users, optionally filtered by ID, name, or email.

```go
resp, err := client.Users().List(ctx, users.UserListFilter{})
resp, err := client.Users().List(ctx, users.UserListFilter{Name: "john"})
resp, err := client.Users().List(ctx, users.UserListFilter{Email: "john@example.com"})
```

### Create a User

Create a new user with a name, display name, and email.

```go
user, err := client.Users().Create(ctx, users.CreateUserRequest{
    Name:        "johndoe",
    DisplayName: "John Doe",
    Email:       "john@example.com",
})
```

### Rename a User

Change a user's name by their ID.

```go
user, err := client.Users().Rename(ctx, "user-id-123", "newname")
```

### Delete a User

Remove a user by their ID.

```go
client.Users().Delete(ctx, "user-id-123")
```

## Types

**User** — the main entity:

| Field           | Type        | Description                 |
| --------------- | ----------- | --------------------------- |
| `ID`            | `string`    | Unique identifier           |
| `Name`          | `string`    | Username                    |
| `DisplayName`   | `string`    | Human-readable display name |
| `Email`         | `string`    | Email address               |
| `Provider`      | `string`    | Auth provider (e.g. `oidc`) |
| `ProviderID`    | `string`    | ID from the auth provider   |
| `ProfilePicURL` | `string`    | Avatar URL                  |
| `CreatedAt`     | `time.Time` | When the user was created   |

**Request types:**

- `UserListFilter` — optional `ID`, `Name`, `Email` filters.
- `CreateUserRequest` — accepts `Name` (required), `DisplayName`, `Email`, `PictureURL`.

**Response types:**

- `UsersResponse` — wraps `[]User` (returned by List).
- `UserResponse` — wraps a single `User` (returned by Create and Rename).
