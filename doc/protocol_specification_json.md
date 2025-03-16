# 服务端与嵌入式设备JSON通信协议规范 v1.1

## 1. 协议概述

本协议规范定义了服务端与嵌入式设备之间的通信标准，采用JSON格式以确保通信的灵活性和可读性。本规范基于原二进制协议v1.1版本转换而来，保留了所有功能特性。

## 2. 通用协议定义

### 2.1 协议格式
```json
{
  "header": {
    "device_id": "设备唯一标识",
    "sensor_type": "传感器类型标识符",
    "command": "命令标识符",
    "timestamp": "时间戳(可选)"
  },
  "data": {
    // 根据命令和传感器类型不同，包含不同的数据字段
  }
}
```

### 2.2 传感器类型定义（sensor_type字段）

| 传感器类型标识符 | 说明 |
|------------|------|
| six_in_one | 温度、湿度、CO2、PM2.5、TVOC、CH2O |
| strain | 应变、温度 |
| displacement | 位移、温度 |
| noise | 噪声、温度 |
| voice_device | 用于语音播报功能 |
| custom | 用户自定义传感器类型 |

### 2.3 通用命令定义（command字段）

| 命令标识符 | 说明 |
|--------|------|
| heartbeat | 设备定期发送心跳信息 |
| data_report | 设备数据上报 |
| param_set | 设置设备参数 |
| param_query | 查询设备参数 |
| control | 控制设备动作 |
| status_query | 查询设备状态 |
| config_response | 响应配置命令 |
| error_report | 报告错误状态 |
| voice_data | 传输语音数据 |
| voice_control | 控制语音播放（开始/暂停/停止等） |

### 2.4 数据字段格式
根据命令不同，`data`字段包含不同的内容：

1. 心跳包（heartbeat）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "six_in_one",
    "command": "heartbeat"
  },
  "data": {
    "device_type": "设备类型代码",
    "uptime": 123456,
    "status": 1
  }
}
```

2. 数据上报（data_report）：
每种传感器类型有其特定的数据格式。

2.1 六合一传感器数据上报：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "six_in_one",
    "command": "data_report"
  },
  "data": {
    "temperature": 25.0,
    "humidity": 60.0,
    "co2": 1000,
    "pm25": 30,
    "tvoc": 500,
    "ch2o": 1.00
  }
}
```

2.2 应变传感器数据上报：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "strain",
    "command": "data_report"
  },
  "data": {
    "strain": 1000,
    "temperature": 25.0
  }
}
```

2.3 位移传感器数据上报：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "displacement",
    "command": "data_report"
  },
  "data": {
    "displacement": 10.00,
    "temperature": 25.0
  }
}
```

2.4 噪声传感器数据上报：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "noise",
    "command": "data_report"
  },
  "data": {
    "noise": 65.0,
    "temperature": 25.0
  }
}
```

3. 参数设置（param_set）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "six_in_one",
    "command": "param_set"
  },
  "data": {
    "param_id": "参数ID",
    "value": "参数值" // 可以是任何JSON支持的数据类型
  }
}
```

4. 参数查询（param_query）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "six_in_one",
    "command": "param_query"
  },
  "data": {
    "param_id": "参数ID"
  }
}
```

5. 控制命令（control）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "six_in_one",
    "command": "control"
  },
  "data": {
    "command_id": "命令ID",
    "params": {} // 命令参数，可以包含任何所需字段
  }
}
```

6. 语音数据传输（voice_data）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_data"
  },
  "data": {
    "voice_id": 1,
    "fragment_index": 0,
    "total_fragments": 5,
    "fragment_size": 1024,
    "format": "opus",
    "sample_rate": 16000,
    "channels": 1,
    "status": "start"
  }
}
```

注意：
- 实际音频数据通过UDP协议单独传输，不包含在此JSON消息中。
- 对于流式生成的语音，将`total_fragments`设为`0`表示数量未知。
- `status`字段可以是以下值之一：
  - `"start"`: 开始一个新的语音传输
  - `"end"`: 结束当前语音传输
  - 不提供此字段则表示一个完整的非流式语音传输

6.1 UDP语音数据格式:

通过UDP传输的音频数据采用以下二进制格式：

```
[数据包头(16字节)][加密的音频数据(N字节)]
```

 6.1.1 数据包头(PacketHeader)字段格式
| 偏移量 | 长度(字节) | 说明 |
|-------|----------|------|
| 0     | 1        | 数据包类型标识，固定值0x01 |
| 1-2   | 2        | 保留字段 |
| 2-4   | 2        | 音频数据大小(网络字节序) |
| 4-12  | 8        | 加密初始化向量(IV)，确保加密安全性 |
| 12-16 | 4        | 序列号(网络字节序)，每个包递增，用于检测丢包和顺序重组 |

6.1.2 加密方式
音频数据使用AES-CTR(计数器)模式加密，加密密钥在MQTT通道的握手过程中建立。加密过程：

1. 使用数据包头中的IV与密钥一起生成密钥流
2. 将密钥流与原始音频数据进行XOR操作，产生加密数据
3. 接收方使用相同的密钥和IV进行解密

6.1.3 序列号管理
- 每个新的语音会话从序列号1开始
- 发送方每发送一个包，序列号递增1
- 接收方通过检查序列号的连续性来检测丢包
- 当检测到序列号不连续时，表示可能有数据包丢失

7. 语音数据结束通知（用于流式生成）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_data"
  },
  "data": {
    "voice_id": 1,
    "status": "end",
    "total_fragments": 42  // 实际传输的总片段数
  }
}
```

8. 语音播放控制（voice_control）：
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_control"
  },
  "data": {
    "control": "play|pause|stop|volume",
    "voice_id": 1,
    "params": {
      // 根据control类型不同，可能包含不同参数
      "volume": 80 // 如果control为"volume"
    }
  }
}
```

## 3. 示例

### 3.1 数据上报示例

#### 3.1.1 六合一传感器数据上报
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "six_in_one",
    "command": "data_report",
    "timestamp": 1678985421
  },
  "data": {
    "temperature": 25.0,
    "humidity": 60.0,
    "co2": 1000,
    "pm25": 30,
    "tvoc": 500,
    "ch2o": 1.00
  }
}
```

#### 3.1.2 应变传感器数据上报
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "strain",
    "command": "data_report",
    "timestamp": 1678985425
  },
  "data": {
    "strain": 1000,
    "temperature": 25.0
  }
}
```

### 3.2 语音播报示例

#### 3.2.1 语音数据传输
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_data",
    "timestamp": 1678985430
  },
  "data": {
    "voice_id": 1678985430,
    "fragment_index": 0,
    "total_fragments": 0,  // 0表示流式生成，数量未知
    "fragment_size": 1024,
    "format": "opus",
    "sample_rate": 16000,
    "channels": 1,
    "status": "start"
  }
}
```

#### 3.2.2 语音数据结束通知
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_data",
    "timestamp": 1678985445
  },
  "data": {
    "voice_id": 1678985430,
    "status": "end",
    "total_fragments": 42  // 实际传输的总片段数
  }
}
```

#### 3.2.3 语音播放控制
```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_control",
    "timestamp": 1678985445
  },
  "data": {
    "control": "play",
    "voice_id": 1
  }
}
```

```json
{
  "header": {
    "device_id": "01",
    "sensor_type": "voice_device",
    "command": "voice_control",
    "timestamp": 1678985455
  },
  "data": {
    "control": "volume",
    "voice_id": 1,
    "params": {
      "volume": 80
    }
  }
}
```

## 4. 通信传输

### 4.1 通信方式
本协议可通过多种通信方式实现，包括但不限于：
- MQTT
- WebSocket
- HTTP REST API
- TCP/IP直接通信

### 4.2 安全性考虑
建议在生产环境中通过以下方式保障通信安全：
1. 使用TLS/SSL对通信加密
2. 实现访问令牌或API密钥认证机制
3. 对敏感数据进行额外加密

### 4.3 错误处理
当发生通信错误时，建议返回标准错误响应：
```json
{
  "header": {
    "device_id": "XX",
    "sensor_type": "XX",
    "command": "error_report"
  },
  "data": {
    "code": "错误代码",
    "message": "错误描述"
  }
}
```

## 5. 版本历史

| 版本 | 日期 | 修改说明 |
|------|------|----------|
| 1.0 | 2025-03-16 | 初始版本，基于二进制协议v1.1转换为JSON格式 |
| 1.1 | 2025-03-16 | 增加UDP语音数据传输格式 |
