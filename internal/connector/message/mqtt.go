package message

import (
	"fmt"
	"log"
	"time"

	"wmjtyd-iot/internal/connector/tsdb"
	deviceMqtt "wmjtyd-iot/internal/module/device/logic/mqtt"
	"wmjtyd-iot/internal/module/device/model"

	mqttLib "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type MQTTClient struct {
	client     mqttLib.Client
	publisher  *deviceMqtt.Publisher
	subscriber *deviceMqtt.Subscriber
}

type MQTTConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	ClientID  string
	KeepAlive int
	QoS       int
	Namespace string
	GroupID   string
	NodeID    string
}

func StartMQTTClient() *MQTTClient {
	config := MQTTConfig{
		Host:      viper.GetString("mqtt.host"),
		Port:      viper.GetInt("mqtt.port"),
		Username:  viper.GetString("mqtt.username"),
		Password:  viper.GetString("mqtt.password"),
		ClientID:  viper.GetString("mqtt.client_id"),
		KeepAlive: viper.GetInt("mqtt.keepalive"),
		QoS:       viper.GetInt("mqtt.qos"),
		Namespace: viper.GetString("mqtt.namespace"),
		GroupID:   viper.GetString("mqtt.group_id"),
		NodeID:    viper.GetString("mqtt.node_id"),
	}

	opts := mqttLib.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Host, config.Port))
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetKeepAlive(time.Duration(config.KeepAlive) * time.Second)

	client := mqttLib.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT Host: %v", token.Error())
	}

	fmt.Println("Successfully connected to MQTT Host")

	// Get TDEngine connection
	tdengine, err := tsdb.GetTDEngine()
	if err != nil {
		log.Fatalf("Failed to get TDEngine connection: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	mqttClient := &MQTTClient{
		client:     client,
		publisher:  deviceMqtt.NewPublisher(client),
		subscriber: deviceMqtt.NewSubscriber(client, tdengine, logger),
	}

	// Start connection monitor in a separate goroutine
	go func() {
		for {
			if !client.IsConnected() {
				log.Println("MQTT client disconnected")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return mqttClient
}

func (m *MQTTClient) PublishDeviceBirth(deviceID string) error {
	return m.publisher.PublishBirth(deviceID)
}

func (m *MQTTClient) PublishDeviceDeath(deviceID string) error {
	return m.publisher.PublishDeath(deviceID)
}

func (m *MQTTClient) PublishDeviceData(deviceID string, data map[string]interface{}) error {
	return m.publisher.PublishData(deviceID, data)
}

func (m *MQTTClient) PublishDeviceDelete(deviceID string) error {
	return m.publisher.PublishDelete(deviceID)
}

func (m *MQTTClient) SubscribeDeviceCommands(deviceID string, handler func(cmd string, data *model.MQTTMessage)) error {
	return m.subscriber.SubscribeCommands(deviceID, handler)
}

func (m *MQTTClient) SubscribeDeviceData(deviceID string, handler func(data *model.MQTTMessage)) error {
	return m.subscriber.SubscribeData(deviceID, handler)
}
