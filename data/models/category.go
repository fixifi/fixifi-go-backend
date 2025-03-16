package models

import "time"

type Category struct {
	ID           uint     `json:"id" gorm:"primaryKey"`
	Icon         string   `json:"icon"`
	MainCategory string   `json:"mainCategory" gorm:"not null;uniqueIndex" validate:"required"`
	SubCategory  []string `json:"subCategory" gorm:"serializer:json" validate:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
