# codespace-proto-go

Gitea Codespace 的共享协议定义和生成的 Go binding。

## Scope

该模块包含：

- Gitea 与 Codespace Manager 通信使用的 `.proto` 源文件。
- 从协议生成的 Go 代码。
- RPC 两端共用的枚举和请求、响应类型。

Gitea 服务逻辑位于 `gitea` 模块，Manager 与 Gateway 逻辑位于 `codespace` 模块；本模块只提供双方从同一协议生成的类型。
