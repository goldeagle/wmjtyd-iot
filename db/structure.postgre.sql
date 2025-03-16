CREATE TABLE "wmjtyd_captcha" (
  "id" SERIAL PRIMARY KEY,
  "phone" varchar(255) NOT NULL,
  "code" varchar(255) NOT NULL,
  "expire_time" integer NOT NULL,
  "create_time" integer NOT NULL,
  "update_time" integer NOT NULL,
  "status" integer NOT NULL
);

CREATE TABLE "wmjtyd_token" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer NOT NULL,
  "token" varchar(255) NOT NULL,
  "expire_time" integer NOT NULL,
  "create_time" integer NOT NULL,
  "update_time" integer NOT NULL,
  "status" integer NOT NULL
);

CREATE TABLE "wmjtyd_user" (
  "id" SERIAL PRIMARY KEY,
  "group_id" integer NOT NULL,
  "username" varchar(255) NOT NULL,
  "nickname" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "mobile" varchar(255) NOT NULL,
  "avatar" varchar(255) NOT NULL,
  "gender" integer NOT NULL,
  "birthday" varchar(255) NOT NULL,
  "money" integer NOT NULL,
  "score" integer NOT NULL,
  "lastlogintime" integer NOT NULL,
  "lastloginip" varchar(255) NOT NULL,
  "loginfailure" integer NOT NULL,
  "joinip" varchar(255) NOT NULL,
  "jointime" integer NOT NULL,
  "motto" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "salt" varchar(255) NOT NULL,
  "status" varchar(255) NOT NULL,
  "updatetime" integer NOT NULL,
  "createtime" integer NOT NULL
);

CREATE TABLE "wmjtyd_device_info" (
  "id" SERIAL PRIMARY KEY,
  "uuid" char(36) NOT NULL,
  "serial" varchar(50),
  "name" varchar(50),
  "alias" varchar(100),
  "type" varchar(50) DEFAULT 'HMI',
  "state" smallint DEFAULT 1,
  "mqtt_client_id" varchar(50),
  "mqtt_password" varchar(50),
  "created_time" integer,
  "updated_time" integer,
  "model_id" integer,
  "data_source_id" integer,
  "location_id" integer,
  "configure_id" integer,
  "user_id" integer,
  "group_id" integer,
  "online_time" integer,
  "offline_time" integer
);

CREATE TABLE "wmjtyd_device_log" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "log_type" varchar(50),
  "log_level" varchar(20),
  "content" text,
  "timestamp" bigint
);

CREATE TABLE "wmjtyd_device_status" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "status" integer NOT NULL,
  "ts" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_warning" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "level" integer NOT NULL,
  "code" varchar(32) NOT NULL,
  "message" text NOT NULL,
  "ts" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_position" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "name" varchar(64) NOT NULL,
  "longitude" double precision NOT NULL,
  "latitude" double precision NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_net" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "mcu_id" varchar(64) NOT NULL,
  "position_id" integer,
  "ts" integer NOT NULL,
  "state" integer,
  "ip" varchar(32),
  "mask" varchar(32),
  "gw" varchar(64),
  "mac" varchar(64),
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_model" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "manufacturer" varchar(100) NOT NULL,
  "description" varchar(250) NOT NULL,
  "protocol" varchar(50) NOT NULL,
  "firmware_id" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_firmware" (
  "id" SERIAL PRIMARY KEY,
  "caption" varchar(100) NOT NULL,
  "model_ids" text NOT NULL,
  "version" varchar(50) NOT NULL,
  "url" varchar(250) NOT NULL,
  "create_time" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_config" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "config_key" varchar(100) NOT NULL,
  "config_val" text NOT NULL,
  "version" varchar(50) NOT NULL,
  "status" varchar(20) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_device_cmd" (
  "id" SERIAL PRIMARY KEY,
  "device_id" integer NOT NULL,
  "cmd_type" varchar(50) NOT NULL,
  "cmd_data" text NOT NULL,
  "status" varchar(20) NOT NULL,
  "exec_time" bigint NOT NULL,
  "result" text,
  "error_msg" text,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_remote_heart_beat" (
  "id" SERIAL PRIMARY KEY,
  "device_id" varchar(50) NOT NULL,
  "timestamp" timestamp NOT NULL,
  "status" smallint DEFAULT 1,
  "created_time" integer NOT NULL,
  "updated_time" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_remote_jobs" (
  "id" SERIAL PRIMARY KEY,
  "job_id" varchar(50) NOT NULL,
  "device_id" varchar(50) NOT NULL,
  "job_type" varchar(50),
  "status" smallint DEFAULT 1,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "created_time" integer NOT NULL,
  "updated_time" integer NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_file" (
  "file_id" SERIAL PRIMARY KEY,
  "file_name" varchar(255) NOT NULL,
  "file_path" varchar(255) NOT NULL,
  "file_size" bigint NOT NULL,
  "file_type" varchar(50) NOT NULL,
  "upload_time" bigint NOT NULL,
  "uploader_id" integer NOT NULL,
  "status" integer NOT NULL,
  "description" varchar(255),
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "wmjtyd_log" (
  "id" SERIAL PRIMARY KEY,
  "log_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "user_name" varchar(255) NOT NULL,
  "module" varchar(255) NOT NULL,
  "action" varchar(255) NOT NULL,
  "description" varchar(255) NOT NULL,
  "ip" varchar(255) NOT NULL,
  "create_time" bigint NOT NULL,
  "status" integer NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp
);

CREATE TABLE "wmjtyd_menu" (
  "id" SERIAL PRIMARY KEY,
  "menu_id" integer NOT NULL,
  "pid" integer NOT NULL,
  "controller_name" varchar(255) NOT NULL,
  "title" varchar(255) NOT NULL,
  "pk_id" varchar(255) NOT NULL,
  "table_name" varchar(255) NOT NULL,
  "is_create" integer NOT NULL,
  "status" integer NOT NULL,
  "sortid" integer NOT NULL,
  "table_status" integer NOT NULL,
  "is_url" integer NOT NULL,
  "url" varchar(255) NOT NULL,
  "menu_icon" varchar(255) NOT NULL,
  "tab_menu" varchar(255) NOT NULL,
  "app_id" integer NOT NULL,
  "is_submit" integer NOT NULL,
  "upload_config_id" integer NOT NULL,
  "connect" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" timestamp
);