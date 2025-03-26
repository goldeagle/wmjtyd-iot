-- TDengine数据库表结构定义

-- 设备基础信息表
CREATE TABLE IF NOT EXISTS devices (
    device_id VARCHAR(64) PRIMARY KEY,
    device_type VARCHAR(32),
    location VARCHAR(128),
    group_id VARCHAR(64),
    created_time TIMESTAMP
);

-- 六合一传感器数据表
CREATE TABLE IF NOT EXISTS six_in_one_data (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    temperature FLOAT,
    humidity FLOAT,
    co2 INT,
    pm25 INT,
    tvoc INT,
    ch2o FLOAT,
    uptime BIGINT,
    status INT
);

-- 应变传感器数据表
CREATE TABLE IF NOT EXISTS strain_data (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    strain INT,
    temperature FLOAT,
    uptime BIGINT,
    status INT
);

-- 位移传感器数据表
CREATE TABLE IF NOT EXISTS displacement_data (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    displacement FLOAT,
    temperature FLOAT
);

-- 噪声传感器数据表
CREATE TABLE IF NOT EXISTS noise_data (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    noise FLOAT,
    temperature FLOAT
);

-- 设备心跳数据表
CREATE TABLE IF NOT EXISTS heartbeat_data (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    device_type VARCHAR(32),
    uptime BIGINT,
    status INT
);

-- 设备命令记录表
CREATE TABLE IF NOT EXISTS command_log (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    command_id VARCHAR(64),
    command_type VARCHAR(32),
    params JSON,
    status VARCHAR(16)
);

-- 设备错误记录表
CREATE TABLE IF NOT EXISTS error_log (
    ts TIMESTAMP PRIMARY KEY,
    device_id VARCHAR(64),
    error_code VARCHAR(32),
    error_message VARCHAR(256)
);

-- 创建超级表(可选)
CREATE STABLE IF NOT EXISTS sensor_data_stable (
    ts TIMESTAMP,
    device_id VARCHAR(64),
    value FLOAT,
    status INT
) TAGS (
    sensor_type VARCHAR(32),
    location VARCHAR(128)
);
