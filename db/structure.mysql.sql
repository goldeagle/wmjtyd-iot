CREATE DATABASE IF NOT EXISTS `wmjtyd-iot`;
CREATE TABLE IF NOT EXISTS `wmjtyd_captcha` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `phone` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `expire_time` int NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  `status` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_token` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` int NOT NULL,
  `token` varchar(255) NOT NULL,
  `expire_time` int NOT NULL,
  `create_time` int NOT NULL,
  `update_time` int NOT NULL,
  `status` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_user` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `group_id` int NOT NULL,
  `username` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `mobile` varchar(255) NOT NULL,
  `avatar` varchar(255) NOT NULL,
  `gender` int NOT NULL,
  `birthday` varchar(255) NOT NULL,
  `money` int NOT NULL,
  `score` int NOT NULL,
  `lastlogintime` int NOT NULL,
  `lastloginip` varchar(255) NOT NULL,
  `loginfailure` int NOT NULL,
  `joinip` varchar(255) NOT NULL,
  `jointime` int NOT NULL,
  `motto` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `salt` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `updatetime` int NOT NULL,
  `createtime` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_info` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `uuid` char(36) NOT NULL,
  `serial` varchar(50),
  `name` varchar(50),
  `alias` varchar(100),
  `type` enum('GW','HMI','PLC','CAMERA','SENSOR','MISC') DEFAULT 'HMI',
  `state` tinyint(4) DEFAULT 1,
  `mqtt_client_id` varchar(50),
  `mqtt_password` varchar(50),
  `created_time` int,
  `updated_time` int,
  `model_id` int,
  `location_id` int,
  `user_id` int,
  `group_id` int,
  `online_time` int,
  `offline_time` int
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_log` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int unsigned NOT NULL,
  `log_type` varchar(50),
  `log_level` varchar(20),
  `content` text,
  `timestamp` bigint
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_status` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int NOT NULL,
  `status` int NOT NULL COMMENT '状态:0-离线,1-在线,2-故障,3-维护',
  `ts` int NOT NULL COMMENT '状态时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_warning` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int NOT NULL,
  `level` int NOT NULL COMMENT '告警级别:1-一般,2-严重,3-紧急',
  `code` varchar(32) NOT NULL,
  `message` text NOT NULL,
  `ts` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_position` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int NOT NULL,
  `name` varchar(64) NOT NULL,
  `longitude` double NOT NULL,
  `latitude` double NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_net` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int NOT NULL,
  `mcu_id` varchar(64) NOT NULL,
  `position_id` int,
  `ts` int NOT NULL,
  `state` int COMMENT '当前设备状态：0-无应答（备用）；1-正常工作；2-有警报；3-有故障；4-无法工作',
  `ip` varchar(32),
  `mask` varchar(32),
  `gw` varchar(64),
  `mac` varchar(64),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_model` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(100) NOT NULL,
  `manufacturer` varchar(100) NOT NULL,
  `description` varchar(250) NOT NULL,
  `protocol` varchar(50) NOT NULL,
  `firmware_id` int unsigned NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_firmware` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `caption` varchar(100) NOT NULL,
  `model_ids` text NOT NULL,
  `version` varchar(50) NOT NULL,
  `url` varchar(250) NOT NULL,
  `create_time` bigint NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_config` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int unsigned NOT NULL,
  `config_key` varchar(100) NOT NULL,
  `config_val` text NOT NULL,
  `version` varchar(50) NOT NULL,
  `status` varchar(20) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_device_cmd` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` int unsigned NOT NULL,
  `cmd_type` varchar(50) NOT NULL,
  `cmd_data` text NOT NULL,
  `status` varchar(20) NOT NULL,
  `exec_time` bigint NOT NULL,
  `result` text,
  `error_msg` text,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_remote_heart_beat` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `device_id` varchar(50) NOT NULL,
  `timestamp` datetime NOT NULL,
  `status` tinyint(4) DEFAULT 1,
  `created_time` int NOT NULL,
  `updated_time` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_remote_jobs` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `job_id` varchar(50) NOT NULL,
  `device_id` varchar(50) NOT NULL,
  `job_type` varchar(50),
  `status` tinyint(4) DEFAULT 1,
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `created_time` int NOT NULL,
  `updated_time` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_file` (
  `file_id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `file_name` varchar(255) NOT NULL,
  `file_path` varchar(255) NOT NULL,
  `file_size` bigint NOT NULL,
  `file_type` varchar(50) NOT NULL,
  `upload_time` bigint NOT NULL,
  `uploader_id` int NOT NULL,
  `status` int NOT NULL,
  `description` varchar(255),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `log_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `user_name` varchar(255) NOT NULL,
  `module` varchar(255) NOT NULL,
  `action` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `ip` varchar(255) NOT NULL,
  `create_time` bigint(20) NOT NULL,
  `status` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `log_id` (`log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `wmjtyd_menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` int(11) NOT NULL,
  `pid` int(11) NOT NULL,
  `controller_name` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `pk_id` varchar(255) NOT NULL,
  `table_name` varchar(255) NOT NULL,
  `is_create` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `sortid` int(11) NOT NULL,
  `table_status` int(11) NOT NULL,
  `is_url` int(11) NOT NULL,
  `url` varchar(255) NOT NULL,
  `menu_icon` varchar(255) NOT NULL,
  `tab_menu` varchar(255) NOT NULL,
  `app_id` int(11) NOT NULL,
  `is_submit` int(11) NOT NULL,
  `upload_config_id` int(11) NOT NULL,
  `connect` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;