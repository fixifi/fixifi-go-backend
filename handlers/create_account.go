package handlers

import (
	// "fmt"

	"fmt"
	"time"

	"github.com/fixifi/fixifi-go-backend/util/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type Consumer struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Phone     string    `json:"phone" validate:"required,min=10"`
	Email     string    `json:"email" validate:"required,email=valid"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (mainHandler *MainHandler) CreateConsumer() fiber.Handler {
	return func(c fiber.Ctx) error {
		var newConsumer Consumer
		err := c.Bind().Body(&newConsumer)
		if err != nil {
			return response.SendReponse(c, fiber.StatusBadRequest, response.GetErrorResponse(fmt.Errorf("invalid Request: %w", err)))
		}

		if err = mainHandler.Validate.Struct(newConsumer); err != nil {
			validationErros := err.(validator.ValidationErrors)
			return response.SendReponse(c, fiber.StatusBadRequest, response.ValidationError(validationErros))
		}
		
		if err := mainHandler.DB.Create(&newConsumer).Error; err != nil {
			return response.SendReponse(c, fiber.StatusInternalServerError, response.GetErrorResponse(fmt.Errorf("failed to create user :%v", err)))
		}

		return response.SendReponse(c, fiber.StatusCreated, response.Response{
			Status:  response.StatusOK,
            Message: "User created successfully",
		})
	}
}
