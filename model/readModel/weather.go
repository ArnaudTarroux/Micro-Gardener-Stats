package readmodel

import "time"

type WeatherReadModel struct {
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	UpdatedAt   time.Time `json:"updated_at"`
}
