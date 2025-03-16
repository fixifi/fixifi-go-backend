package handlers

import (
	"fmt"

	"github.com/fixifi/fixifi-go-backend/data/models"
	"github.com/fixifi/fixifi-go-backend/util/response"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func (mainHandler *MainHandler) UpdateCategory() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get category ID from params
		categoryID := c.Params("id")
		if categoryID == "" {
			return response.SendReponse(c, fiber.StatusBadRequest,
				response.GetErrorResponse(fmt.Errorf("category ID is required")))
		}

		// Parse update request
		var updateCategory models.Category
		if err := c.Bind().Body(&updateCategory); err != nil {
			return response.SendReponse(c, fiber.StatusBadRequest,
				response.GetErrorResponse(fmt.Errorf("invalid request body: %v", err)))
		}

		// Validate update data
		if err := mainHandler.Validate.Struct(updateCategory); err != nil {
			return response.SendReponse(c, fiber.StatusBadRequest,
				response.GetErrorResponse(fmt.Errorf("validation failed: %v", err)))
		}

		// Check if category exists
		var existingCategory models.Category
		result := mainHandler.DB.First(&existingCategory, categoryID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return response.SendReponse(c, fiber.StatusNotFound,
					response.GetErrorResponse(fmt.Errorf("category not found")))
			}
			return response.SendReponse(c, fiber.StatusInternalServerError,
				response.GetErrorResponse(fmt.Errorf("database error: %v", result.Error)))
		}

		// Check if updated MainCategory already exists for other categories
		var duplicateCheck models.Category
		result = mainHandler.DB.Where("main_category = ? AND id != ?",
			updateCategory.MainCategory, categoryID).First(&duplicateCheck)
		if result.Error == nil {
			return response.SendReponse(c, fiber.StatusConflict,
				response.GetErrorResponse(fmt.Errorf("category with name %s already exists",
					updateCategory.MainCategory)))
		}

		// Update category
		updates := models.Category{
			Icon:         updateCategory.Icon,
			MainCategory: updateCategory.MainCategory,
			SubCategory:  updateCategory.SubCategory,
		}
		if err := mainHandler.DB.Model(&existingCategory).Updates(updates).Error; err != nil {
			return response.SendReponse(c, fiber.StatusInternalServerError,
				response.GetErrorResponse(fmt.Errorf("failed to update category: %v", err)))
		}

		return response.SendReponse(c, fiber.StatusOK, response.ResponseWithData{
			Status:  response.StatusOK,
			Message: "Category updated successfully",
			Data:    existingCategory,
		})
	}
}
