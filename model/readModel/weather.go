package readModel

import "time"

type WeatherReadModel struct {
	Temperature float32
	Humidity    float32
	updatedAt   time.Time
}
