package models

import (
	"time"
)

type Truck struct {
	ID        uint
	CreatedAt *time.Time `gorm:"default:'CURRENT_TIMESTAMP'"`
	UpdatedAt *time.Time `gorm:"default:'CURRENT_TIMESTAMP'"`
	DeletedAt *time.Time `gorm:"default:NULL"`
	Plate     string     `json:"plate"`
	Latitude  *float64   `gorm:"default:NULL;" json:"latitude"`
	Longitude *float64   `gorm:"default:NULL;" json:"longitude"`
	Brand     string     `json:"brand"`
	Model     string     `json:"model"`
	Speed     *float32   `gorm:"default:NULL;" json:"speed"`
	Trips     []Trip     `gorm:"foreignKey:TruckID"`
	CompanyID *uint      `json:"company_id"`
	Company   *Company   `gorm:"foreignKey:CompanyID;constraint:onUpdate:CASCADE,onDelete:RESTRICT" json:"company,omitempty"`
}
