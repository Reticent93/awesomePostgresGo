package models

import "gorm.io/gorm"

//Cars schema of the cars table
type Cars struct {
	gorm.Model
	CarMake  string `json:"car_make"`
	CarModel string `json:"car_model" gorm:"unique_index"`
	Mileage  int    `json:"mileage"`
	Clean    bool   `json:"clean"`
}
