package handlers

import (
	"fmt"

	"github.com/fixifi/fixifi-go-backend/data/models"
	"github.com/fixifi/fixifi-go-backend/util/response"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (mainHandler *MainHandler) CreateCategories() fiber.Handler {
	return func(c fiber.Ctx) error {
		var newCategory models.Category
		if err := c.Bind().Body(&newCategory); err != nil {
			return response.SendReponse(c, fiber.StatusBadRequest, response.GetErrorResponse(fmt.Errorf("invalid request format %v", err)))

		}

		// Check if category exists
		var existingCategory models.Category
		result := mainHandler.DB.Where("main_category = ?", newCategory.MainCategory).First(&existingCategory)
		if result.Error != gorm.ErrRecordNotFound {
			return response.SendReponse(c, fiber.StatusConflict, response.GetErrorResponse(fmt.Errorf("category already exists")))
		}

		// Create new category
		if err := mainHandler.DB.Create(&newCategory).Error; err != nil {
			return response.SendReponse(c, fiber.StatusInternalServerError, response.GetErrorResponse(fmt.Errorf("failed to create category :%v", err)))
		}

		return response.SendReponse(c, fiber.StatusCreated, response.ResponseWithData{
			Status:  response.StatusOK,
			Message: "Category created successfully",
			Data:    newCategory,
		})
	}
}
