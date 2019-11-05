package handler

import (
	"fmt"
	"log"

	model "github.com/mg/microgardener/model"
)

type WeatherMessageHandler struct {
	name string
}

func (h WeatherMessageHandler) supports(message Message) bool {
	return message.MessageType == "weather"
}

func (h WeatherMessageHandler) execute(message Message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)

	weather := model.NewWeather(message.Payload)

	fmt.Println(weather.Temperature)
}
