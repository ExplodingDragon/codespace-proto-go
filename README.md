# Gitea Codespace Protocol for Go

This repository contains the Protocol Buffer definitions and generated Go
bindings shared by Gitea and the Gitea Codespace manager. Keeping the source
protocol and generated packages in one Go module makes each protocol change
reviewable and consumable as a single revision.

## Repository layout

- [`proto/codespace/v1`](proto/codespace/v1) contains the source Protocol Buffer
  definitions.
- [`codespace/v1`](codespace/v1) contains the generated protobuf messages and
  protocol checks.
- [`codespace/v1/codespacev1connect`](codespace/v1/codespacev1connect) contains
  the generated Connect RPC client and server bindings.

Application behavior does not belong in this module. Gitea implements the
control plane, while the Codespace manager and gateway consume these types to
implement runtime operations.

## Requirements

- Go 1.26.4 or later.
- `buf`, `protoc-gen-go`, and `protoc-gen-connect-go` for linting or regenerating
  bindings.

Install the pinned generator toolchain with:

```bash
make install
```

The tools are installed in the standard Go binary directory selected by
`go env GOBIN` or `go env GOPATH`.

## Generate bindings

After changing a file under `proto/`, format and regenerate the checked-in Go
bindings:

```bash
make format
make generate
```

Generation writes directly to `codespace/v1` because the protobuf `go_package`
is `gitea.dev/codespace-proto-go/codespace/v1`. There is no separate generated
repository or intermediate `gen` directory.

## Validation

Check protocol style and formatting:

```bash
make lint
```

Run the Go tests that verify the generated service names and protocol field
contracts:

```bash
make test
```

Run the complete protocol workflow before submitting a change:

```bash
make build
```

`make build` runs linting, regeneration, and tests in that order. Generated
changes must be committed together with their source `.proto` changes so both
consumers use the same contract.

## License

This project is licensed under the MIT License. See [`LICENSE`](LICENSE) for the
full text.
