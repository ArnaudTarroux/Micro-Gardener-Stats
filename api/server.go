package api

import (
	"fmt"
	"net/http"

	handler "github.com/mg/microgardener/api/handler"
)

func Init() {
	fmt.Println("API server started")
	fmt.Println("API ready to handle connexion")

	http.Handle("/api/weather", new(handler.WeatherApiHandler))

	http.ListenAndServe(":8000", nil)
}
