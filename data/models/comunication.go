package models

import "time"

type Communication struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Message    string  `json:"message" gorm:"not null" validate:"required"`
	Cost       float32 `json:"cost" gorm:"not null" validate:"required,gt=0"`
	IsAccepted bool    `json:"isAccepted" gorm:"not null" validate:"required"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
