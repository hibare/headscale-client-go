# Contributing

## Prerequisites

- Go 1.26+
- Docker (E2E tests only)
- golangci-lint (`make init`)

## Workflow

1. Fork and clone the repo
2. `make init`
3. Create a branch: `git checkout -b feat/your-feature`
4. Make changes
5. **Pre-submit:** `go fmt ./... && golangci-lint run && make test`
6. Commit: `git commit -m "feat: description"`
7. Push and open a PR

## Guidelines

- All exported symbols need Go doc comments
- New code should include tests
- Update docs if behaviour changes

### PR Checklist

- [ ] `go build ./...` passes
- [ ] `make test` passes
- [ ] `golangci-lint run` passes

## Questions?

Open an [issue](https://github.com/hibare/headscale-client-go/issues).
