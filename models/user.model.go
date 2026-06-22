package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string
	Password  string
	Name      string
	Role      string   `gorm:"default:'COMPANY_ADMIN';NOT NULL" json:"role"`
	CompanyID *uint    `json:"company_id"`
	Company   *Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:CASCADE,onDelete:RESTRICT" json:"company,omitempty"`
}
