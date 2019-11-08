package handler

import (
	"log"
	"time"

	"github.com/google/uuid"
	model "github.com/mg/microgardener/model"
	repositories "github.com/mg/microgardener/persistence/repository"
)

type weatherMessageHandler struct {
	name string
}

func (h weatherMessageHandler) supports(message message) bool {
	return message.MessageType == "weather"
}

func (h weatherMessageHandler) execute(message message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)

	// weather := model.NewWeatherFromJson(message.Payload)

	uuid, _ := uuid.NewUUID()
	payload := string(message.Payload)
	event := model.NewEvent(uuid.String(), message.ControllerID, message.MessageType, payload, time.Now())

	repository := new(repositories.SqlEventRepository)
	repository.Save(*event)
}
