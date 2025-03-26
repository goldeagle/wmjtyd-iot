package mqtt

import (
	"encoding/json"
	"fmt"
	"time"

	mqttLib "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

type Publisher struct {
	client mqttLib.Client
}

func NewPublisher(client mqttLib.Client) *Publisher {
	return &Publisher{client: client}
}

func (p *Publisher) getTopic(deviceID string, messageType string) string {
	baseTopic := viper.GetString("mqtt.topic")
	return fmt.Sprintf("%s/%s/%s", baseTopic, deviceID, messageType)
}

func (p *Publisher) PublishBirth(deviceID string) error {
	topic := p.getTopic(deviceID, "birth")
	payload := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"online":    true,
	}
	return p.publish(topic, payload)
}

func (p *Publisher) PublishDeath(deviceID string) error {
	topic := p.getTopic(deviceID, "death")
	payload := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"online":    false,
	}
	return p.publish(topic, payload)
}

func (p *Publisher) PublishData(deviceID string, data map[string]interface{}) error {
	topic := p.getTopic(deviceID, "data")
	return p.publish(topic, data)
}

func (p *Publisher) PublishDelete(deviceID string) error {
	topic := p.getTopic(deviceID, "delete")
	payload := map[string]interface{}{
		"timestamp": time.Now().Unix(),
	}
	return p.publish(topic, payload)
}

func (p *Publisher) publish(topic string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	qos := byte(viper.GetInt("mqtt.qos"))
	retained := false
	token := p.client.Publish(topic, qos, retained, jsonData)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to publish message: %v", token.Error())
	}
	return nil
}
