# Gitea Codespace Protocol for Go

This module is the public Protocol Buffer contract between Gitea and a Codespace Manager. It contains the `codespace/v1` sources, generated Go messages, and Connect clients and servers. Authorization, persistence, Kubernetes resources, and internal Agent/Gateway/Cache protocols remain in their owning repositories.

## Generate

Use the Go version declared in [`go.mod`](go.mod). Install the pinned generators, format the sources, and update the checked-in bindings with:

```sh
make install
make format
make generate
```

Generated code and its `.proto` source are reviewed and released together. After publishing a revision, update the remote module dependency in both Gitea and Codespace.

## Validate

```sh
make lint
make test
```

`make build` runs formatting checks, generation, and tests as the complete protocol workflow.

## License

[MIT](LICENSE).
