package model

import "encoding/json"

type Fan struct {
	Duty float32 `json:"duty"`
}

func NewFan(rawData []byte) Fan {
	fan := new(Fan)
	json.Unmarshal(rawData, &fan)

	return *fan
}
