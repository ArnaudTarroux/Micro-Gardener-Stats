package query

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mg/microgardener/persistence/repository"
)

type WeatherQuery struct{}

func (h WeatherQuery) Handle(c *gin.Context) {
	fmt.Println("weather query !!")

	eventRepository := new(repository.SqlEventRepository)
	event := eventRepository.GetLastEventByType("weather")
	if nil == event {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	var payload map[string]float32
	_ = json.Unmarshal([]byte(event.GetPayload()), &payload)

	c.JSON(http.StatusOK, gin.H{
		"temperature": payload["temperature"],
		"humidity":    payload["humidity"],
		"created_at":  event.GetCreatedAt(),
	})
}
