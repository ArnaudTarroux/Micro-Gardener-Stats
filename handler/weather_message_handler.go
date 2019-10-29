package handler

import (
	"log"
)

type WeatherMessageHandler struct {
	name string
}

func (h WeatherMessageHandler) supports(message Message) bool {
	return true
}

func (h WeatherMessageHandler) execute(message Message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)
}
