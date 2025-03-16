package handlers

import (
	"fmt"

	"github.com/fixifi/fixifi-go-backend/data/models"
	"github.com/fixifi/fixifi-go-backend/util/response"
	"github.com/gofiber/fiber/v3"
)

func (mainHandler *MainHandler) FetchCategories() fiber.Handler {
	return func(c fiber.Ctx) error {
		var categories []models.Category

		if err := mainHandler.DB.Find(&categories).Error; err != nil {
			return response.SendReponse(c, fiber.StatusInternalServerError,
				response.GetErrorResponse(fmt.Errorf("failed to fetch categories: %v", err)))
		}

		return response.SendReponse(c, fiber.StatusOK, response.ResponseWithData{
			Status:  response.StatusOK,
			Message: "Categories fetched successfully",
			Data:    categories,
		})
	}
}

