# wmjtyd-iot Backend CLI

[中文版](README_CN.md)

**Important Note**: This project is only for technical learning and exchange in certain scenarios, and has not reached a production-ready state. Please use with caution.

wmjtyd-iot (IoT Control and Communication Platform) is a comprehensive backend system for managing IoT devices, data, and communications. It provides:

- Device management and monitoring
- Real-time data collection and processing
- Message brokering via MQTT
- RESTful API for system integration
- Time-series data storage and analysis
- User authentication and authorization

## Table of Contents
- [Available Commands](#available-commands)
- [System Requirements](#system-requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

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

## System Requirements

### Required Services
- **Database**: MySQL 5.7+ or MariaDB 10.5+ or PostgreSQL 12+ (for relational data)
- **Time-Series Database**: TDengine 3.3+ (for time-series data storage)
- **Message Broker**: MQTT Broker (e.g. EMQX, Mosquitto)
- **Cache**: Redis 6.0+
- **Object Storage**: MinIO or S3-compatible storage

### Development Environment
- **Go**: 1.18+ (https://golang.org/dl/)
- **Make**: For build automation
- **Git**: For version control

### Build Commands
The following Makefile commands are available for development and build automation:

- `make daemon`: Build the application's backend daemon only
- `make build`: Build the application
- `make clean`: Clean build artifacts
- `make test`: Run tests
- `make fmt`: Format code
- `make run`: Build and run the application
- `make install-deps`: Install dependencies
- `make lint`: Run static analysis
- `make help`: Show available commands and their descriptions

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/goldeagle/wmjtyd-iot.git
   cd wmjtyd-iot-platform
   ```

2. Build the binary:
   ```bash
   go build
   ```

3. Install dependencies:
   ```bash
   make deps
   ```

4. Run the application:
   ```bash
   ./wmjtyd-iot serve
   ```

## Configuration

The system can be configured via `config/config.yaml`. Key configuration options include:

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

## API Documentation

The system provides RESTful APIs documented using Swagger. You can access the API documentation at:

- Swagger UI: `http://localhost:3000/swagger`
- OpenAPI Spec: `http://localhost:3000/swagger.json`

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please ensure your code follows our coding standards and includes appropriate tests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.