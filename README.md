# codespace-proto-go

Shared protocol definitions and generated Go bindings for the Gitea codespace system.

## Scope

This module is intended to hold:

- `.proto` source files for Gitea and codespace manager communication.
- Generated Go code from those proto files.
- Thin shared enums and request / response types that must stay protocol-aligned.

Business logic should stay in the `gitea` and `gitea-codespace` modules.
