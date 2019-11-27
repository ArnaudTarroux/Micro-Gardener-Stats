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
	eventRepository := new(repositories.SqlEventRepository)

	now := time.Now()
	lastEvent := eventRepository.GetLastEventByType(message.ControllerID, message.MessageType)
	lastRegisteredWeatherEvent = lastEvent.GetCreatedAt()
	// We persist the events in database only each 15 minutes
	// if !now.After(lastRegisteredWeatherEvent.Add(time.Minute * 15)) {
	// 	return
	// }

	uuid, _ := uuid.NewUUID()
	payload := string(message.Payload)
	event := model.NewEvent(uuid.String(), message.ControllerID, message.MessageType, payload, now)

	eventRepository.Save(event)

	weather := model.NewWeatherFromJson(message.Payload)
	weatherRepository := new(repositories.SqlWeatherRepository)
	weatherRepository.UpdateForController(message.ControllerID, *weather)
}
