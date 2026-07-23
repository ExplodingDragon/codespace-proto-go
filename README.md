# codespace-proto-go

Gitea Codespace 的共享协议定义和生成的 Go binding。

## Scope

该模块包含：

- Gitea 与 Codespace Manager 通信使用的 `.proto` 源文件。
- 从协议生成的 Go 代码。
- RPC 两端共用的枚举和请求、响应类型。

Gitea 服务逻辑位于 `gitea` 模块，Manager 与 Gateway 逻辑位于 `codespace` 模块；本模块只提供双方从同一协议生成的类型。

## 生成流程

本仓同时保存 `.proto` 源文件和生成后的 Go 代码。这样设计是为了让 Gitea 与 Codespace Manager 引用同一个 Go module，不需要在协议仓和生成仓之间再做一次同步，协议字段、Connect 服务名和生成类型能在一次提交中闭环。

常用命令：

- `make install` 安装 `buf`、`protoc-gen-go` 和 `protoc-gen-connect-go`。
- `make lint` 检查 proto 语法、命名和格式。
- `make format` 格式化 proto 文件。
- `make generate` 根据 `proto/` 下的协议重新生成 Go 代码。
- `make test` 执行协议包的 Go 测试，确认服务名和 `protocol_version` 字段编号等协议约定没有偏移。
- `make build` 依次执行 `lint`、`generate` 和 `test`，用于提交前确认协议仓闭环。

生成代码直接写入 `codespace/v1`，这是因为 `go_package` 指向 `gitea.dev/codespace-proto-go/codespace/v1`。本仓不再额外维护 `gen/` 目录或二次推送到其他 Go 仓库，原因是当前仓库本身就是 Gitea 与 Manager 消费的 Go module，增加中间目录会让开发者需要判断哪个目录才是最终依赖来源。
