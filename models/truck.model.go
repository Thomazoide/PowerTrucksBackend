package models

import "gorm.io/gorm"

type Truck struct {
	gorm.Model
	Plate     string  `json:"plate"`
	Latitude  float64 `gorm:"default:NULL;" json:"latitude"`
	Longitude float64 `gorm:"default:NULL;" json:"longitude"`
	Speed     float32 `gorm:"default:NULL;" json:"speed"`
	Trips     []Trip  `gorm:"foreignKey:TruckID"`
}
