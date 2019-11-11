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

	uuid, _ := uuid.NewUUID()
	payload := string(message.Payload)
	event := model.NewEvent(uuid.String(), message.ControllerID, message.MessageType, payload, time.Now())

	eventRepository := new(repositories.SqlEventRepository)
	eventRepository.Save(*event)

	weather := model.NewWeatherFromJson(message.Payload)
	weatherRepository := new(repositories.SqlWeatherRepository)
	weatherRepository.Save(*weather, message.ControllerID)
}
