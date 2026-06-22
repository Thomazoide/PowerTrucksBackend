package models

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Latitude  float64 `gorm:"NOT NULL;" json:"latitude"`
	Longitude float64 `gorm:"NOT NULL;" json:"longitude"`
	Name      string  `gorm:"NOT NULL" json:"name"`
	BleMAC    string  `gorm:"NOT NULL;uniqueIndex" json:"ble_mac"`
}
