package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mg/microgardener/api/handler"
)

func Init() {
	fmt.Println("API server started")
	fmt.Println("API ready to handle connexion")

	r := gin.Default()
	r.GET("/api/weather", new(handler.WeatherApiHandler).Handle)
	r.Run("0.0.0.0:8000")
}
