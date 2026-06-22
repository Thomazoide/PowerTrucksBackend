package models

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Rut       string `gorm:"unique;index;" json:"rut"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Email     string `gorm:"unique" json:"email"`
	Cellphone string `json:"cellphone"`
	Trips     []Trip `gorm:"foreignKey:DriverID"`
}
