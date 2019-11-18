package handler

import (
	"fmt"
	"net/http"

	repositories "github.com/mg/microgardener/persistence/repository"
)

type WeatherApiHandler struct{}

func (h WeatherApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("weather handler !!")
	w.Header().Add("Content-Type", "application/json")

	eventRepository := new(repositories.SqlEventRepository)
	event := eventRepository.GetLastEventByType("weather")

	fmt.Fprint(w, event)
}
