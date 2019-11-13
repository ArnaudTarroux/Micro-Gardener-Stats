package handler

import (
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Handler
type Handler interface {
	execute(message Message)
	supports(message Message) bool
}

// Message containing topic and payload
type Message struct {
	Topic        string
	ControllerID string
	MessageType  string
	Payload      []byte
}

var handlers = []Handler{
	new(WeatherMessageHandler),
	new(FanMessageHandler),
}

func processMessage(message Message) {
	for _, handler := range handlers {
		if handler.supports(message) {
			go handler.execute(message)
		}
	}
}

// DefaultPublishHandler Consume all message
func DefaultPublishHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Message received %d : %s from %s\n", msg.MessageID(), msg.Payload(), msg.Topic())

	topic := msg.Topic()
	splittedTopic := strings.Split(topic, "/")
	if len(splittedTopic) < 3 {
		log.Printf("Cannot handle the message, topic invalid: %s \n", topic)
		return
	}

	message := Message{Topic: topic, Payload: msg.Payload(), ControllerID: splittedTopic[1], MessageType: splittedTopic[2]}
	processMessage(message)
}
