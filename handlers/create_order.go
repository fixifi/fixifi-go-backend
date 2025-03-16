package handlers

import (
	"fmt"

	"github.com/fixifi/fixifi-go-backend/data/models"
	"github.com/fixifi/fixifi-go-backend/util/response"
	"github.com/gofiber/fiber/v3"
)

func (mainHandler *MainHandler) CreateOrader() fiber.Handler {

	return func(c fiber.Ctx) error {
		var newOrder models.Order
		newOrder.Status = models.Created

		if err := c.Bind().Body(&newOrder); err != nil {
			return response.SendReponse(c, fiber.StatusBadRequest, response.GetErrorResponse(fmt.Errorf("invalid Request: %w", err)))
		}

		if err := mainHandler.DB.Create(&newOrder).Error; err != nil {
			return response.SendReponse(c, fiber.StatusInternalServerError, response.GetErrorResponse(fmt.Errorf("failed to create order :%v", err)))
		}


		return c.JSON(fiber.Map{
			"message": "CreateOrder",
		})
	}
}		
