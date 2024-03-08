package main

import (
	"encoding/json"
	"log"
	"testing"
	"time"
	 Sensor "pratica/SensorData"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const (
	MQTTBroker   = "tcp://localhost:4567"
	ClientID     = "test-client"
	MQTPTopic    = "/pond-2"
)

func configureMQTTClient() MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker(MQTTBroker)
	opts.SetClientID(ClientID)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect MQTT broker: %v", token.Error())
	}
	return client
}

func validateFields(t *testing.T, msg map[string]int, expectedFields []string) {
	for _, field := range expectedFields {
		if _, ok := msg[field]; !ok {
			t.Errorf("Expected field: %s", field)
			return
		}
	}
}

func TestConnection(t *testing.T) {
	client := configureMQTTClient()
	defer client.Disconnect(250)

	t.Log("Connection with broker MQTT succeeded")
}

func TestDataValidation(t *testing.T) {
	msg := Sensor.SensorData()
	expectedFields := []string{"freezer1", "freezer2", "Geladeira1", "Geladeira2"}
	validateFields(t, msg, expectedFields)
	t.Log("Data validation successful")
}

func TestPublisher(t *testing.T) {
	client := configureMQTTClient()
	defer client.Disconnect(250)

	received := make(chan bool)

	token := client.Subscribe(MQTPTopic, 1, func(client MQTT.Client, msg MQTT.Message) {

		var data map[string]int
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			t.Errorf("Error validating message: %v", err)
			return
		}

		expectedFields := []string{"freezer1", "freezer2", "Geladeira1", "Geladeira2"}
		validateFields(t, data, expectedFields)

		received <- true
	})
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to subscribe MQTT topic: %v", token.Error())
	}

	msg := Sensor.SensorData()
	jsonData, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Error converting to JSON: %v", err)
	}

	token = client.Publish(MQTPTopic, 0, false, string(jsonData))
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to publish message: %v", token.Error())
	}

	select {
	case <-received:
		t.Log("Message received")
	case <-time.After(5 * time.Second):
		t.Fatalf("Timeout")
	}
	
}

func Tests(t *testing.T) {
	t.Run("TestConnection", TestConnection)
    t.Run("TestDataValidation", TestDataValidation)
    t.Run("TestPublisher", TestPublisher)
}