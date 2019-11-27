package repository

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mg/microgardener/model"
	"github.com/mg/microgardener/model/readmodel"
	persistence "github.com/mg/microgardener/persistence"
)

type WeatherRepository interface {
	GetByController(controller string) (error, readmodel.WeatherReadModel)
	UpdateForController(controller string, weather model.Weather)
}

type SqlWeatherRepository struct{}

func (repository SqlWeatherRepository) GetByController(controller string) (error, readmodel.WeatherReadModel) {
	db := persistence.Init()
	defer db.Close()

	sqlStmt := `
		SELECT temperature, humidity, updated_at FROM public.weather WHERE controller = $1
	`
	rows := db.QueryRow(sqlStmt, controller)

	var weather readmodel.WeatherReadModel
	err := rows.Scan(&weather.Temperature, &weather.Humidity, &weather.UpdatedAt)
	if err != nil {
		log.Print(err)
		return err, weather
	}

	return nil, weather
}

func (repository SqlWeatherRepository) UpdateForController(controller string, weather model.Weather) {
	weatherRepository := new(SqlWeatherRepository)
	err, _ := weatherRepository.GetByController(controller)
	if nil != err {
		db := persistence.Init()
		defer db.Close()

		sqlStmt := `
			INSERT INTO public.weather (id, controller, humidity, temperature, updated_at)
			VALUES ($1, $2, $3, $4, $5)
		`
		uuid, _ := uuid.NewUUID()
		_, err := db.Exec(
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

		log.Printf("New projection created for controller %s", controller)

		return
	}

	db := persistence.Init()
	defer db.Close()

	sqlStmt := `
		UPDATE public.weather
			SET humidity = $1, temperature = $2, updated_at = $3
		WHERE controller = $4
	`

	_, err = db.Exec(
		sqlStmt,
		weather.Humidity,
		weather.Temperature,
		time.Now(),
		controller,
	)
	if err != nil {
		panic(err)
	}

	log.Printf("Projection for controller %s was updated", controller)
}
