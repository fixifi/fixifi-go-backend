package models

import "gorm.io/gorm"

type Consumer struct {
	gorm.Model
	MobileNo        string    `json:"mobileNo" gorm:"not null;unique"`                 // Required & Unique
	Email           string    `json:"email" gorm:"not null;email"`                  // Required & Unique
	Name            string    `json:"name" gorm:"not null"`                            // Required
	BrowseService   *Category `json:"browseService" gorm:"foreignKey:BrowseServiceID"` // Optional relationship
	BrowseServiceID *uint     `json:"browseServiceId"`                                 // Foreign Key
	AddressID       uint      `json:"addressId"`                                       // Foreign Key
	Address         Address   `json:"address" gorm:"constraint:OnDelete:CASCADE;"`     // Required
}
