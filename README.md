# Gitea Codespace Protocol for Go

This module contains the Protocol Buffer sources and generated Go bindings
shared by Gitea Codespace. Protocol changes and generated packages are reviewed
and released together so every consumer can depend on one module revision.

## Packages

| Package | Boundary |
| --- | --- |
| `codespace/v1` | Gitea and Manager lifecycle control |
| `component/v1` | Manager, Gateway, and Cache control |
| `agent/v1` | Manager, Runtime Agent, and Gateway streams |

The matching sources live below `proto/`. Generated Connect clients and servers
are in each package's `*connect` directory. This module defines wire contracts;
authorization, persistence, and runtime behavior remain in Gitea or the
Codespace implementation.

## Toolchain

Use the Go version declared in [`go.mod`](go.mod). Install the pinned Protocol
Buffer tools with:

```sh
make install
```

The command uses the standard Go binary directory selected by `go env GOBIN`
or `go env GOPATH`.

## Generate and validate

After changing a file below `proto/`, format and regenerate the checked-in Go
bindings:

```sh
make format
make generate
```

Validate source style, generated contracts, and the module tests with:

```sh
make lint
make test
```

`make build` runs linting, generation, and tests as the complete protocol
workflow. Commit generated changes with their `.proto` sources, then update the
published dependency in Gitea and Codespace.

## License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE).
