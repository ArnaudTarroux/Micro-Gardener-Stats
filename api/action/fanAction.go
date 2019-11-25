package action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FanAction struct{}

type FanActionBody struct {
	Day   float32 `json:"day" binding:"required"`
	Night float32 `json:"night" binding:"required"`
}

func (a FanAction) Handle(c *gin.Context) {
	fmt.Println("fan action !!")

	var json FanActionBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, json)
}
