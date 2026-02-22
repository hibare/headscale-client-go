# Contributing

Contributions are welcome! This document provides guidelines for contributing to headscale-client-go.

## Development Setup

### Prerequisites

- Go 1.26+
- Docker (for E2E tests)
- golangci-lint

### Setup

```sh
# Clone the repository
git clone https://github.com/hibare/headscale-client-go.git
cd headscale-client-go

# Install development tools
make init
```

## Code Style

- Run `go fmt ./...` before committing
- Follow Go conventions and idiomatic patterns
- Ensure all exported types have documentation comments
- Run `golangci-lint run` to check for issues

## Testing

### Unit Tests

```sh
make test
```

### E2E Tests

E2E tests require Docker and use testcontainers-go to spin up Headscale and Tailscale containers.

```sh
make e2e-test
```

## Pull Requests

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Make your changes
4. Run tests and lint
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin feature/my-feature`)
7. Open a Pull Request

### PR Guidelines

- Write clear commit messages
- Add tests for new functionality
- Update documentation if needed
- Ensure CI passes

## Pre-commit Hooks

```sh
make install-pre-commit
pre-commit run --all-files
```

## Questions?

Open an issue for bugs, features, or questions.
