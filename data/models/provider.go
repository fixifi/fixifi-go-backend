package models

import (
	"time"

	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	Name         string    `json:"name" gorm:"not null" validate:"required"`
	LicenseNo    string    `json:"licenseNo"`
	FoundingDate time.Time `json:"foundingDate" gorm:"not null" validate:"required"`
}

type Equipment struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null" validate:"required"`
	Type     string `json:"type" gorm:"not null" validate:"required"`
	Quantity int    `json:"quantity" gorm:"default:1"`
}

type Provider struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null" validate:"required"`
	MobileNo   string `json:"mobileNo" gorm:"not null;unique" validate:"required,e164"`
	CategoryID uint
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	AddressID  uint
	Address    Address `json:"address" gorm:"constraint:OnDelete:CASCADE;not null"`
	BusinessID uint
	Business   Business    `json:"business" gorm:"foreignKey:BusinessID"`
	Equipments []Equipment `json:"equipments" gorm:"many2many:provider_equipments;"`
}
