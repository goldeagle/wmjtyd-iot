package message

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

type MQTTConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	ClientID  string
	KeepAlive int
	QoS       int
	Namespace string
}

func StartMQTTClient() {
	config := MQTTConfig{
		Host:      viper.GetString("mqtt.host"),
		Port:      viper.GetInt("mqtt.port"),
		Username:  viper.GetString("mqtt.username"),
		Password:  viper.GetString("mqtt.password"),
		ClientID:  viper.GetString("mqtt.client_id"),
		KeepAlive: viper.GetInt("mqtt.keepalive"),
		QoS:       viper.GetInt("mqtt.qos"),
		Namespace: viper.GetString("mqtt.namespace"),
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Host, config.Port))
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetKeepAlive(time.Duration(config.KeepAlive) * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		topic := msg.Topic()
		if config.Namespace != "" {
			topic = config.Namespace + "/" + topic
		}
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), topic)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to MQTT Host: %v", token.Error())
	}

	fmt.Println("Successfully connected to MQTT Host")

	// Example publish function
	Publish := func(topic string, payload string) {
		token := client.Publish(topic, byte(config.QoS), false, payload)
		token.Wait()
		if token.Error() != nil {
			log.Printf("Failed to publish message: %v", token.Error())
		}
	}

	// Publish a test message
	Publish("test/topic", "Hello MQTT")

	// Start message handler in a separate goroutine
	go func() {
		for {
			if !client.IsConnected() {
				log.Println("MQTT client disconnected")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
