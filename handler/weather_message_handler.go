package handler

import (
	"log"
	"time"

	"github.com/google/uuid"
	model "github.com/mg/microgardener/model"
	repositories "github.com/mg/microgardener/persistence/repository"
)

type WeatherMessageHandler struct {
	name string
}

func (h WeatherMessageHandler) supports(message message) bool {
	return message.MessageType == "weather"
}

func (h WeatherMessageHandler) execute(message message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)

	uuid, _ := uuid.NewUUID()
	payload := string(message.Payload)
	event := model.NewEvent(uuid.String(), message.ControllerID, message.MessageType, payload, time.Now())

	eventRepository := new(repositories.SqlEventRepository)
	eventRepository.Save(event)
}
