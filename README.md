# wmjtyd-iot Backend CLI

wmjtyd-iot (IoT Control and Communication Platform) is a backend system for IOT system.

## Available Commands

### `wmjtyd-iot migrate`
Migrate database by creating new database, importing SQL file and setting up user permissions.

**Usage:**
```bash
wmjtyd-iot migrate --user <username> --file <sql_file_path>
```

**Flags:**
- `--user`: Database user with create database right (required)
- `--file`: Database file to load (required)
- `--host`: Database server host (default: 127.0.0.1)
- `--port`: Database server port (default: 3306)

### `wmjtyd-iot serve`
Run web api server & mqtt client.

**Usage:**
```bash
wmjtyd-iot serve [flags]
```

**Flags:**
- `--config`: Config file (default: config.yaml)
- `--host`: Server host (default: 127.0.0.1)
- `--port`: Server port (default: 3000)
- `--prefork`: Enable prefork (default: false)
- `--http`: Enable HTTP server (default: true)
- `--grpc`: Enable gRPC server (default: false)
- `--websocket`: Enable WebSocket server (default: false)

### `wmjtyd-iot version`
Print the version number of wmjtyd-iot and Go environment.

**Usage:**
```bash
wmjtyd-iot version
```

## Getting Started

1. Clone the repository
2. Run `go build` to build the binary
3. Use the commands as described above