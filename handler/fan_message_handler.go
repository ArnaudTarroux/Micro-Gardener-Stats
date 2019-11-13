package handler

import (
	"fmt"
	"log"

	model "github.com/mg/microgardener/model"
)

type FanMessageHandler struct {
	name string
}

func (h FanMessageHandler) supports(message Message) bool {
	return message.MessageType == "fan"
}

func (h FanMessageHandler) execute(message Message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)

	fan := model.NewFan(message.Payload)

	fmt.Println(fan.Duty)
}
