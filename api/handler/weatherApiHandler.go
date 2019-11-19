package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	repositories "github.com/mg/microgardener/persistence/repository"
)

type WeatherApiHandler struct{}

type weatherReadModel struct {
	temperature float32 `json:"temperature"`
	humidity    float32 `json:"humidity"`
	createdAt   time.Time
}

func (h WeatherApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("weather handler !!")
	w.Header().Add("Content-Type", "application/json")

	eventRepository := new(repositories.SqlEventRepository)
	event := eventRepository.GetLastEventByType("weather")

	readModel := weatherReadModel{}
	json.Unmarshal([]byte(event.GetPayload()), &readModel)
	readModel.createdAt = event.GetCreatedAt()

	fmt.Fprint(w, readModel)
}
