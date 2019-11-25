package repository

import (
	"github.com/mg/microgardener/model"
	"github.com/mg/microgardener/model/readModel"
)

type WeatherRepository interface {
	func GetByController(controller string) Weather
	func UpdateForController(controller string, weather Weather)
}

type SqlWeatherRepository struct{}

func (repository SqlWeatherRepository) GetByController(controller string) WeatherReadModel {

}

func (repository SqlWeatherRepository) UpdateForController(controller string, weather Weather) {
	
}