package model

import "encoding/json"

type Car struct {
	Name         string
	Price        float64
	ImageURL     string
	Size         string
	Fuel         float64
	Transmission string
	Engine       string
	Displacement float64 // 排量
	MaxSpeed     float64
	Acceleration float64
}

func FromJsonObj(o interface{}) (Car, error) {
	var car Car
	s, err := json.Marshal(o)
	if err != nil {
		return car, err
	}
	err = json.Unmarshal(s, &car)
	return car, err
}
