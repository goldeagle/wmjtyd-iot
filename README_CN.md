# wmjtyd-iot 后端 CLI

wmjtyd-iot (物联网控制与通信平台) 是一个用于物联网的后端系统。它提供了用于数据库迁移、服务器管理和版本控制的CLI工具。

## 可用命令

### `wmjtyd-iot migrate`
通过创建新数据库、导入SQL文件并设置用户权限来迁移数据库。

**用法:**
```bash
wmjtyd-iot migrate --user <用户名> --file <sql文件路径>
```

**参数:**
- `--user`: 具有创建数据库权限的数据库用户 (必填)
- `--file`: 要加载的数据库文件 (必填)
- `--host`: 数据库服务器主机 (默认: 127.0.0.1)
- `--port`: 数据库服务器端口 (默认: 3306)

### `wmjtyd-iot serve`
运行web api服务器和mqtt客户端。

**用法:**
```bash
wmjtyd-iot serve [参数]
```

**参数:**
- `--config`: 配置文件 (默认: config.yaml)
- `--host`: 服务器主机 (默认: 127.0.0.1)
- `--port`: 服务器端口 (默认: 3000)
- `--prefork`: 启用prefork (默认: false)
- `--http`: 启用HTTP服务器 (默认: true)
- `--grpc`: 启用gRPC服务器 (默认: false)
- `--websocket`: 启用WebSocket服务器 (默认: false)

### `wmjtyd-iot version`
打印wmjtyd-iot版本号和Go环境版本信息。

**用法:**
```bash
wmjtyd-iot version
```

## 快速开始

1. 克隆仓库
2. 运行 `go build` 构建二进制文件
3. 使用上述命令