package models

import "gorm.io/gorm"

type Trip struct {
	gorm.Model
	TruckID   uint   `json:"truck_id"`
	DriverID  uint   `json:"driver_id"`
	StartedAt string `json:"started_at"` //ISO string de fecha
	EndedAt   string `json:"ended_at"`
}
