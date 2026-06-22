package models

import "gorm.io/gorm"

type TripStop struct {
	gorm.Model
	TripID     uint
	PlaceID    uint
	DetectedAt string
	LeftAt     string
}
