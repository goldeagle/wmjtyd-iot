# wmjtyd-iot 后端 CLI

[English Version](README.md)

**重要提示**: 本项目仅用于某些场景下的技术学习和交流，并未达到产品化状态，请谨慎使用。

wmjtyd-iot (物联网控制与通信平台) 是一个全面的后端系统，用于管理物联网设备、数据和通信。它提供了以下功能：

- 设备管理和监控
- 实时数据收集和处理
- 通过MQTT进行消息代理
- 用于系统集成的RESTful API
- 时间序列数据存储和分析
- 用户认证和授权

## 目录
- [可用命令](#可用命令)
- [系统要求](#系统要求)
- [安装](#安装)
- [配置](#配置)
- [API文档](#api文档)
- [贡献指南](#贡献指南)
- [许可证](#许可证)

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

## 系统要求

### 所需服务
- **数据库**: MySQL 5.7+ 或 MaraiDB 10.5+ 或 PostgreSQL 12+ (用于关系型数据)
- **时间序列数据库**: TDengine 3.3+ (用于时间序列数据存储)
- **消息代理**: MQTT Broker (例如 EMQX, Mosquitto)
- **缓存**: Redis 6.0+
- **对象存储**: MinIO 或 S3兼容存储

### 开发环境
- **Go**: 1.18+ (https://golang.org/dl/)
- **Make**: 用于构建自动化
- **Git**: 用于版本控制

### 构建命令
以下Makefile命令可用于开发和构建自动化：

- `make daemon`: 仅构建应用程序的后台守护进程
- `make build`: 构建应用程序
- `make clean`: 清理构建产物
- `make test`: 运行测试
- `make fmt`: 格式化代码
- `make run`: 构建并运行应用程序
- `make install-deps`: 安装依赖
- `make lint`: 运行静态分析
- `make help`: 显示可用命令及其描述

## 安装

1. 克隆仓库：
   ```bash
   git clone https://github.com/goldeagle/wmjtyd-iot.git
   cd wmjtyd-iot-platform
   ```

2. 构建二进制文件：
   ```bash
   go build
   ```

3. 安装依赖：
   ```bash
   make deps
   ```

4. 运行应用程序：
   ```bash
   ./wmjtyd-iot serve
   ```

## 配置

系统可以通过`config/config.yaml`进行配置。关键配置选项包括：

```yaml
server:
  host: 0.0.0.0
  port: 3000
  timeout: 30s

database:
  mysql:
    host: localhost
    port: 3306
    user: root
    password: secret
    name: iot_platform

mqtt:
  broker: tcp://localhost:1883
  client_id: wmjtyd-iot
  qos: 1
```

## API文档

系统提供了使用Swagger记录的RESTful API。您可以通过以下方式访问API文档：

- Swagger UI: `http://localhost:3000/swagger`
- OpenAPI 规范: `http://localhost:3000/swagger.json`

## 贡献指南

我们欢迎贡献！请按照以下步骤进行：

1. Fork 仓库
2. 创建您的功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m '添加一些AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开Pull Request

请确保您的代码遵循我们的编码标准，并包含适当的测试。

## 许可证

本项目采用MIT许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。