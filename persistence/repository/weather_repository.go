package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	model "github.com/mg/microgardener/model"
	persistence "github.com/mg/microgardener/persistence"
)

type WeatherRepository interface {
	Save(model.Weather)
}

type SqlWeatherRepository struct {
}

func (repository SqlWeatherRepository) Save(weather model.Weather, controller string) {
	db := persistence.Init()
	defer db.Close()

	queryResult, err := db.Query("SELECT * FROM public.weather WHERE controller = $1", controller)
	defer queryResult.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(queryResult)
	queryResult.Next()

	uuid, _ := uuid.NewUUID()
	sqlStmt := `
		INSERT INTO public.weather (id, controller, humidity, temperature, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	result, err := db.Exec(
		sqlStmt,
		uuid.String(),
		controller,
		weather.Humidity,
		weather.Temperature,
		time.Now(),
	)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New weather created %d", rowsAffected)
}

func (repository SqlWeatherRepository) Update(weather model.Weather, controller string) {

}
