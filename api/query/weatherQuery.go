package query

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mg/microgardener/persistence/repository"
)

type WeatherQuery struct{}

func (h WeatherQuery) Handle(c *gin.Context) {
	fmt.Println("weather query !!")

	weatherRepository := new(repository.SqlWeatherRepository)
	err, weather := weatherRepository.GetByController("stats")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
	}

	c.JSON(http.StatusOK, weather)
}
