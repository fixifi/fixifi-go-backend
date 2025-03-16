package handlers

import (
	"fmt"

	"github.com/fixifi/fixifi-go-backend/data/models"
	"github.com/fixifi/fixifi-go-backend/util/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func (mainHandler *MainHandler) CreateProvider() fiber.Handler {
	return func(c fiber.Ctx) error {
		var newProvider models.Provider
		err := c.Bind().Body(&newProvider)
		if err != nil {
			return response.SendReponse(c, fiber.StatusBadRequest, response.GetErrorResponse(fmt.Errorf("invalid Request: %w", err)))
		}

		if err = mainHandler.Validate.Struct(newProvider); err != nil {
			validationErros := err.(validator.ValidationErrors)
			return response.SendReponse(c, fiber.StatusBadRequest, response.ValidationError(validationErros))
		}

		if err := mainHandler.DB.Create(&newProvider).Error; err != nil {
			return response.SendReponse(c, fiber.StatusInternalServerError, response.GetErrorResponse(fmt.Errorf("failed to create user :%v", err)))
		}

		return response.SendReponse(c, fiber.StatusCreated, response.Response{
			Status:  response.StatusOK,
			Message: "provider created successfully",
		})
	}
}
