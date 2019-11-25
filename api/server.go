package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mg/microgardener/api/action"
	"github.com/mg/microgardener/api/query"
)

func Init() {
	fmt.Println("API server started")
	fmt.Println("API ready to handle connexion")

	r := gin.Default()
	r.GET("/api/weather", new(query.WeatherQuery).Handle)
	r.POST("/api/fan", new(action.FanAction).Handle)
	r.Run("0.0.0.0:8000")
}
