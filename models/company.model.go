package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName string   `json:"company_name"`
	CompanyLogo *string  `gorm:"default:NULL" json:"company_logo,omitempty"`
	CompanyRut  string   `gorm:"varchar(25)" json:"company_rut"`
	Users       []User   `gorm:"foreignKey:CompanyID" json:"-"`
	Drivers     []Driver `gorm:"foreignKey:CompanyID" json:"-"`
	Trucks      []Truck  `gorm:"foreignKey:CompanyID" json:"-"`
}
