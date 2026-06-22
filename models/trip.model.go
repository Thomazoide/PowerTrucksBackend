package models

import "gorm.io/gorm"

type Trip struct {
	gorm.Model
	TruckID   uint       `json:"truck_id"`
	DriverID  uint       `json:"driver_id"`
	StartedAt string     `json:"started_at"` //ISO string de fecha
	EndedAt   string     `json:"ended_at"`
	Driver    Driver     `gorm:"foreignKey:DriverID" json:"driver"`
	Truck     Truck      `gorm:"foreignKey:TruckID" json:"truck"`
	TripStops []TripStop `gorm:"foreignKey:TripID" json:"trip_stops,omitempty"`
	WorkOrder *WorkOrder `gorm:"backref" json:"work_order,omitempty"`
}
