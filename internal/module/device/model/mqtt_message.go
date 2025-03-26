package model

import (
	"time"
)

// MQTTMessageHeader 消息头
type MQTTMessageHeader struct {
	DeviceID   string    `json:"device_id"`
	DeviceType string    `json:"device_type"`
	SensorType string    `json:"sensor_type"`
	Command    string    `json:"command"`
	Location   string    `json:"location"`
	GroupID    string    `json:"group_id"`
	Timestamp  time.Time `json:"timestamp"`
}

// MQTTMessage MQTT消息结构
type MQTTMessage struct {
	Header MQTTMessageHeader `json:"header"`
	Data   interface{}       `json:"data"`
}

// HeartbeatData 心跳数据
type HeartbeatData struct {
	DeviceType string `json:"device_type"`
	Uptime     int64  `json:"uptime"`
	Status     int    `json:"status"`
}

// SixInOneData 六合一传感器数据
type SixInOneData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	CO2         int     `json:"co2"`
	PM25        int     `json:"pm25"`
	TVOC        int     `json:"tvoc"`
	CH2O        float32 `json:"ch2o"`
	Uptime      int64   `json:"uptime"`
	Status      int     `json:"status"`
}

// StrainData 应变传感器数据
type StrainData struct {
	Strain      int     `json:"strain"`
	Temperature float32 `json:"temperature"`
	Uptime      int64   `json:"uptime"`
	Status      int     `json:"status"`
}

// DisplacementData 位移传感器数据
type DisplacementData struct {
	Displacement float32 `json:"displacement"`
	Temperature  float32 `json:"temperature"`
}

// NoiseData 噪声传感器数据
type NoiseData struct {
	Noise       float32 `json:"noise"`
	Temperature float32 `json:"temperature"`
}

// ParamSetData 参数设置数据
type ParamSetData struct {
	ParamID string      `json:"param_id"`
	Value   interface{} `json:"value"`
}

// ParamQueryData 参数查询数据
type ParamQueryData struct {
	ParamID string `json:"param_id"`
}

// ControlData 控制命令数据
type ControlData struct {
	CommandID string                 `json:"command_id"`
	Params    map[string]interface{} `json:"params"`
}

// VoiceData 语音数据
type VoiceData struct {
	VoiceID        int    `json:"voice_id"`
	FragmentIndex  int    `json:"fragment_index"`
	TotalFragments int    `json:"total_fragments"`
	FragmentSize   int    `json:"fragment_size"`
	Format         string `json:"format"`
	SampleRate     int    `json:"sample_rate"`
	Channels       int    `json:"channels"`
	Status         string `json:"status"`
}

// VoiceControlData 语音控制数据
type VoiceControlData struct {
	Control string                 `json:"control"`
	VoiceID int                    `json:"voice_id"`
	Params  map[string]interface{} `json:"params"`
}

// ErrorReportData 错误报告数据
type ErrorReportData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
