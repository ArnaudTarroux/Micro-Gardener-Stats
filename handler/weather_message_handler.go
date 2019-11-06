package handler

import (
	"fmt"
	"log"

	model "github.com/mg/microgardener/model"
)

type weatherMessageHandler struct {
	name string
}

func (h weatherMessageHandler) supports(message message) bool {
	return message.MessageType == "weather"
}

func (h weatherMessageHandler) execute(message message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)

	weather := model.NewWeatherFromJson(message.Payload)

	fmt.Println(weather.Temperature)
}
