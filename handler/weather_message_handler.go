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

var lastRegisteredWeatherEvent time.Time

func (h WeatherMessageHandler) execute(message message) {
	log.Printf("Executing... ControllerID: %s | MessageType: %s\n", message.ControllerID, message.MessageType)

	now := time.Now()
	// We persist the events in database only each 15 minutes
	if !now.After(lastRegisteredWeatherEvent.Add(time.Minute * 15)) {
		return
	}

	lastRegisteredWeatherEvent = now
	uuid, _ := uuid.NewUUID()
	payload := string(message.Payload)
	event := model.NewEvent(uuid.String(), message.ControllerID, message.MessageType, payload, now)

	eventRepository := new(repositories.SqlEventRepository)
	eventRepository.Save(event)
}
