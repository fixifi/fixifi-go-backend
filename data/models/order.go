package models

import (
	"github.com/fixifi/fixifi-go-backend/types"
	"gorm.io/gorm"
)

const (
	Created             types.OrderStatus = "created"
	Pending             types.OrderStatus = "pending"
	Canceled            types.OrderStatus = "canceled"
	Accepted            types.OrderStatus = "accepted"
	Fixing              types.OrderStatus = "fixing"
	Completed           types.OrderStatus = "completed"
	ReachedAtLocation   types.OrderStatus = "reachedAtLocation"
	CancelledByConsumer types.OrderStatus = "cancelledByConsumer"
	CancelledByProvider types.OrderStatus = "cancelledByProvider"
)

type OrderImage struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	BeforeFix []string `json:"beforeFix" gorm:"type:text[]"`
	AfterFix  []string `json:"afterFix" gorm:"type:text[]"`
}

type Review struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Images     []string `json:"images" gorm:"type:text[]"`
	ReviewText string   `json:"reviewText" gorm:"not null" validate:"required"`
	Rating     float32  `json:"rating" gorm:"not null" validate:"required,min=0,max=5"`
}

type Order struct {
	gorm.Model
	ConsumerID uint              `json:"consumerId" gorm:"not null;constraint:OnDelete:CASCADE"`
	ProviderID uint              `json:"providerId" gorm:"not null;constraint:OnDelete:CASCADE"`
	CategoryID uint              `json:"categoryId"`
	Category   Category          `json:"category" gorm:"foreignKey:CategoryID"`
	Status     types.OrderStatus `json:"status" gorm:"not null"`
	AddressID  uint
	Address    Address `json:"address" gorm:"constraint:OnDelete:CASCADE;not null"`
	ImagesID   uint
	Images     OrderImage `json:"images" gorm:"foreignKey:ImagesID"`
	ReviewID   uint
	Review     Review `json:"review" gorm:"foreignKey:ReviewID"`
}
