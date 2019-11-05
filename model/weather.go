package model

import "encoding/json"

type Weather struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

func NewWeather(rawData []byte) Weather {
	weather := new(Weather)
	json.Unmarshal(rawData, &weather)

	return *weather
}
