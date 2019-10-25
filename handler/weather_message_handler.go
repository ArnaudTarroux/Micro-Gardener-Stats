package handler

import "fmt"

type WeatherMessageHandler struct {
	name string
}

func (h WeatherMessageHandler) supports(message Message) bool {
	return true
}

func (h WeatherMessageHandler) execute(message Message) {
	fmt.Println(h.name)
	fmt.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)
}

func (h WeatherMessageHandler) new() *WeatherMessageHandler {
	handler := new(WeatherMessageHandler)
	handler.name = "WeatherMessageHandler"

	return handler
}
