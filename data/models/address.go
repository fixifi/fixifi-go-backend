package models

import "time"

type Coordinate struct {
	Latitude  float64 `json:"latitude" gorm:"not null" validate:"required"`
	Longitude float64 `json:"longitude" gorm:"not null" validate:"required"`
}

type Address struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Coordinate Coordinate `json:"coordinate" gorm:"embedded" validate:"required"`
	Street     string     `json:"street" gorm:"not null" validate:"required"`
	City       string     `json:"city" gorm:"not null" validate:"required"`
	State      string     `json:"state" gorm:"not null" validate:"required"`
	ZipCode    string     `json:"zipCode" gorm:"not null" validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
