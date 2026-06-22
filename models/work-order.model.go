package models

import "gorm.io/gorm"

type WorkOrder struct {
	gorm.Model
	OrderNumber string  `gorm:"uniqueIndex;not null;" json:"order_number"`
	Description string  `gorm:"type:text" json:"description"`
	Status      string  `gorm:"default:'PENDING'" json:"status"`
	TripID      *uint   `json:"trip_id"`
	Trip        *Trip   `gorm:"foreignKey:tripID;constraint:onUpdate:CASCADE,onDelete:SET NULL;" json:"trip,omitempty"`
	Places      []Place `gorm:"many2many:work_order_places"`
}
