package handlers

import (
	"github.com/fixifi/fixifi-go-backend/config"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type MainHandler struct {
	DB       *gorm.DB
	Validate *validator.Validate
	App      *fiber.App
	Cfg      *config.Config
}
