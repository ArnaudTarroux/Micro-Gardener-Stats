package action

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mg/microgardener/publisher"
)

type FanAction struct{}

type FanActionBody struct {
	Day   float32 `json:"day" binding:"required"`
	Night float32 `json:"night" binding:"required"`
}

func (a FanAction) Handle(c *gin.Context) {
	fmt.Println("fan action !!")

	var bodyJson FanActionBody
	if err := c.ShouldBindJSON(&bodyJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	payload := map[string]interface{}{
		"action": 10004,
		"value": map[string]float32{
			"day":   bodyJson.Day,
			"night": bodyJson.Night,
		},
	}

	jsonString, _ := json.Marshal(payload)
	publisher.PublishTo("/mg/control", jsonString)

	c.JSON(http.StatusNoContent, gin.H{})
}
