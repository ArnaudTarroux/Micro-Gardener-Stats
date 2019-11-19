package handler

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mg/microgardener/persistence/repository"
)

type WeatherApiHandler struct{}

func (h WeatherApiHandler) Handle(c *gin.Context) {
	fmt.Println("weather handler !!")

	eventRepository := new(repository.SqlEventRepository)
	event := eventRepository.GetLastEventByType("weather")

	var payload map[string]float32
	_ = json.Unmarshal([]byte(event.GetPayload()), &payload)

	c.JSON(200, gin.H{
		"temperature": payload["temperature"],
		"humidity":    payload["humidity"],
		"created_at":  event.GetCreatedAt(),
	})
}
