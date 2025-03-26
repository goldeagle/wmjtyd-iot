package mqtt

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mqttLib "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"wmjtyd-iot/internal/module/device/model"
)

type Subscriber struct {
	client   mqttLib.Client
	tdengine *sql.DB
	logger   *zap.Logger
}

func NewSubscriber(client mqttLib.Client, tdengine *sql.DB, logger *zap.Logger) *Subscriber {
	return &Subscriber{
		client:   client,
		tdengine: tdengine,
		logger:   logger,
	}
}

func (s *Subscriber) getTopic(deviceID string, messageType string) string {
	baseTopic := viper.GetString("mqtt.topic")
	return fmt.Sprintf("%s/%s/%s", baseTopic, deviceID, messageType)
}

func (s *Subscriber) SubscribeAll(deviceID string) error {
	if err := s.SubscribeCommands(deviceID, s.handleCommand); err != nil {
		return err
	}
	if err := s.SubscribeData(deviceID, s.handleData); err != nil {
		return err
	}
	return nil
}

func (s *Subscriber) SubscribeCommands(deviceID string, handler func(cmd string, data *model.MQTTMessage)) error {
	topic := s.getTopic(deviceID, "cmd/+")
	qos := byte(viper.GetInt("mqtt.qos"))

	token := s.client.Subscribe(topic, qos, func(client mqttLib.Client, msg mqttLib.Message) {
		var message model.MQTTMessage
		if err := json.Unmarshal(msg.Payload(), &message); err != nil {
			s.logger.Error("Failed to unmarshal command payload", zap.Error(err))
			return
		}

		cmd := ""
		topicParts := strings.Split(msg.Topic(), "/")
		if len(topicParts) > 0 {
			cmd = topicParts[len(topicParts)-1]
		}

		handler(cmd, &message)
	})

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe to commands: %v", token.Error())
	}
	return nil
}

func (s *Subscriber) SubscribeData(deviceID string, handler func(data *model.MQTTMessage)) error {
	topic := s.getTopic(deviceID, "data")
	qos := byte(viper.GetInt("mqtt.qos"))

	token := s.client.Subscribe(topic, qos, func(client mqttLib.Client, msg mqttLib.Message) {
		var message model.MQTTMessage
		if err := json.Unmarshal(msg.Payload(), &message); err != nil {
			s.logger.Error("Failed to unmarshal data payload", zap.Error(err))
			return
		}

		handler(&message)
	})

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe to data: %v", token.Error())
	}
	return nil
}

func (s *Subscriber) handleCommand(cmd string, msg *model.MQTTMessage) {
	ctx := context.Background()

	switch cmd {
	case "param_set":
		if paramData, ok := msg.Data.(*model.ParamSetData); ok {
			s.saveParamSetData(ctx, msg.Header, paramData)
		}
	case "param_query":
		if paramData, ok := msg.Data.(*model.ParamQueryData); ok {
			s.saveParamQueryData(ctx, msg.Header, paramData)
		}
	case "control":
		if controlData, ok := msg.Data.(*model.ControlData); ok {
			s.saveControlData(ctx, msg.Header, controlData)
		}
	case "voice_control":
		if voiceControlData, ok := msg.Data.(*model.VoiceControlData); ok {
			s.saveVoiceControlData(ctx, msg.Header, voiceControlData)
		}
	default:
		s.logger.Warn("Unknown command type", zap.String("command", cmd))
	}
}

func (s *Subscriber) handleData(msg *model.MQTTMessage) {
	ctx := context.Background()

	switch msg.Header.SensorType {
	case "six_in_one":
		if data, ok := msg.Data.(*model.SixInOneData); ok {
			s.saveSixInOneData(ctx, msg.Header, data)
		}
	case "strain":
		if data, ok := msg.Data.(*model.StrainData); ok {
			s.saveStrainData(ctx, msg.Header, data)
		}
	case "displacement":
		if data, ok := msg.Data.(*model.DisplacementData); ok {
			s.saveDisplacementData(ctx, msg.Header, data)
		}
	case "noise":
		if data, ok := msg.Data.(*model.NoiseData); ok {
			s.saveNoiseData(ctx, msg.Header, data)
		}
	case "voice":
		if data, ok := msg.Data.(*model.VoiceData); ok {
			s.saveVoiceData(ctx, msg.Header, data)
		}
	case "heartbeat":
		if data, ok := msg.Data.(*model.HeartbeatData); ok {
			s.saveHeartbeatData(ctx, msg.Header, data)
		}
	default:
		s.logger.Warn("Unknown sensor type", zap.String("sensorType", msg.Header.SensorType))
	}
}

// 数据保存方法实现
func (s *Subscriber) saveSixInOneData(ctx context.Context, header model.MQTTMessageHeader, data *model.SixInOneData) {
	query := `INSERT INTO six_in_one_data VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.Temperature,
		data.Humidity,
		data.CO2,
		data.PM25,
		data.TVOC,
		data.CH2O,
		data.Uptime,
		data.Status)

	if err != nil {
		s.logger.Error("Failed to save six-in-one data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveStrainData(ctx context.Context, header model.MQTTMessageHeader, data *model.StrainData) {
	query := `INSERT INTO strain_data VALUES (?, ?, ?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.Strain,
		data.Temperature,
		data.Uptime,
		data.Status)

	if err != nil {
		s.logger.Error("Failed to save strain data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveDisplacementData(ctx context.Context, header model.MQTTMessageHeader, data *model.DisplacementData) {
	query := `INSERT INTO displacement_data VALUES (?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.Displacement,
		data.Temperature)

	if err != nil {
		s.logger.Error("Failed to save displacement data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveNoiseData(ctx context.Context, header model.MQTTMessageHeader, data *model.NoiseData) {
	query := `INSERT INTO noise_data VALUES (?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.Noise,
		data.Temperature)

	if err != nil {
		s.logger.Error("Failed to save noise data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveVoiceData(ctx context.Context, header model.MQTTMessageHeader, data *model.VoiceData) {
	query := `INSERT INTO command_log VALUES (?, ?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		fmt.Sprintf("voice_%d", data.VoiceID),
		"voice_command",
		fmt.Sprintf(`{"fragment_index":%d,"total_fragments":%d,"format":"%s"}`,
			data.FragmentIndex,
			data.TotalFragments,
			data.Format))

	if err != nil {
		s.logger.Error("Failed to save voice data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveHeartbeatData(ctx context.Context, header model.MQTTMessageHeader, data *model.HeartbeatData) {
	query := `INSERT INTO heartbeat_data VALUES (?, ?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.DeviceType,
		data.Uptime,
		data.Status)

	if err != nil {
		s.logger.Error("Failed to save heartbeat data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveParamSetData(ctx context.Context, header model.MQTTMessageHeader, data *model.ParamSetData) {
	value, _ := json.Marshal(data.Value)
	query := `INSERT INTO param_set_data VALUES (?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.ParamID,
		string(value))

	if err != nil {
		s.logger.Error("Failed to save param set data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveParamQueryData(ctx context.Context, header model.MQTTMessageHeader, data *model.ParamQueryData) {
	query := `INSERT INTO param_query_data VALUES (?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.ParamID)

	if err != nil {
		s.logger.Error("Failed to save param query data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveVoiceControlData(ctx context.Context, header model.MQTTMessageHeader, data *model.VoiceControlData) {
	params, _ := json.Marshal(data.Params)
	query := `INSERT INTO command_log VALUES (?, ?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		fmt.Sprintf("voice_%d", data.VoiceID),
		"voice_control",
		string(params))

	if err != nil {
		s.logger.Error("Failed to save voice control data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveControlData(ctx context.Context, header model.MQTTMessageHeader, data *model.ControlData) {
	params, _ := json.Marshal(data.Params)
	query := `INSERT INTO command_log VALUES (?, ?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.CommandID,
		"control",
		string(params))

	if err != nil {
		s.logger.Error("Failed to save control data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}

func (s *Subscriber) saveErrorReportData(ctx context.Context, header model.MQTTMessageHeader, data *model.ErrorReportData) {
	query := `INSERT INTO error_log VALUES (?, ?, ?, ?)`
	_, err := s.tdengine.ExecContext(ctx, query,
		time.Now(),
		header.DeviceID,
		data.Code,
		data.Message)

	if err != nil {
		s.logger.Error("Failed to save error report data",
			zap.Error(err),
			zap.String("deviceID", header.DeviceID))
	}
}
