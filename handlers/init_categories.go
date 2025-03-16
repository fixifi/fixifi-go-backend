// package initconst

package handlers

import (
	"fmt"
	"log"

	"github.com/fixifi/fixifi-go-backend/data/models"
	"gorm.io/gorm"
	// "github.com/fixifi/fixifi-go-backend/handlers"
)

var PredefinedCategories = []models.Category{
	{
		Icon:         "🧱",
		MainCategory: "Masons",
		SubCategory: []string{
			"Red bricklaying",
			"Concrete block laying",
			"Firebrick work",
			"Plastering",
			"Tile grouting",
		},
	},
	{
		Icon:         "🏛",
		MainCategory: "Stone Masons",
		SubCategory: []string{
			"Marble installation",
			"Granite installation",
			"Sandstone work",
			"Kota stone fitting",
		},
	},
	{
		Icon:         "🪵",
		MainCategory: "Tile Setters",
		SubCategory: []string{
			"Wall tiling",
			"Floor tiling",
			"Decorative tiling",
			"Granite & marble polishing",
		},
	},
	{
		Icon:         "🛠",
		MainCategory: "Carpenters",
		SubCategory: []string{
			"Wooden frame making",
			"Metal frame making",
			"Trim carpentry",
			"Furniture making",
			"Door & window fitting",
			"False ceiling installation",
		},
	},
	{
		Icon:         "🔌",
		MainCategory: "Electricians",
		SubCategory: []string{
			"Wiring installation",
			"Lighting installation",
			"Machinery wiring",
			"Electrical maintenance",
			"Solar panel installation",
		},
	},
	{
		Icon:         "🔧",
		MainCategory: "Plumbers",
		SubCategory: []string{
			"Water pipe installation",
			"Gas pipeline fitting",
			"Steam pipeline work",
			"Drainage & sewage system setup",
		},
	},
	{
		Icon:         "🔥",
		MainCategory: "Welders & Fabricators",
		SubCategory: []string{
			"Beam welding",
			"Steel frame welding",
			"Gas pipeline welding",
			"Window & gate fabrication",
		},
	},
	{
		Icon:         "🎨",
		MainCategory: "Painters",
		SubCategory: []string{
			"Wall painting",
			"Decorative painting",
			"Weatherproof painting",
			"Spray painting",
			"Wood polishing",
		},
	},
	{
		Icon:         "💧",
		MainCategory: "Waterproofing Experts",
		SubCategory: []string{
			"Roof waterproofing",
			"Basement waterproofing",
			"Bathroom waterproofing",
		},
	},
	{
		Icon:         "🏠",
		MainCategory: "Roofers",
		SubCategory: []string{
			"Asphalt shingle installation",
			"Metal shingle installation",
			"TPO membrane installation",
			"Thatched roofing",
		},
	},
	{
		Icon:         "❄",
		MainCategory: "HVAC Technicians",
		SubCategory: []string{
			"Residential HVAC installation",
			"Industrial HVAC installation",
			"Air duct cleaning",
			"Chiller & cooling tower maintenance",
		},
	},
	{
		Icon:         "🏗",
		MainCategory: "Steelworkers",
		SubCategory: []string{
			"Rebar installation",
			"Structural frame work",
			"Foundation steel placement",
			"Mesh reinforcement",
		},
	},
	{
		Icon:         "🏗",
		MainCategory: "Scaffolding Experts",
		SubCategory: []string{
			"Temporary scaffolding setup",
			"Permanent scaffolding setup",
		},
	},
	{
		Icon:         "🪟",
		MainCategory: "Glass Workers",
		SubCategory: []string{
			"Window glass installation",
			"Toughened glass installation",
			"Glass facade installation",
		},
	},
	{
		Icon:         "⛏",
		MainCategory: "Excavation Workers",
		SubCategory: []string{
			"Manual excavation",
			"JCB operation",
			"Borewell drilling",
		},
	},
}

func (mainHandler *MainHandler) InitializeCategories() error {
	for _, category := range PredefinedCategories {
		// Check if category exists
		var existingCategory models.Category
		result := mainHandler.DB.Where("main_category = ?", category.MainCategory).First(&existingCategory)

		// Only create if category doesn't exist
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Creating new category: %s", category.MainCategory)
			if err := mainHandler.DB.Create(&category).Error; err != nil {
				return fmt.Errorf("failed to create category %s: %v", category.MainCategory, err)
			}
		} else if result.Error != nil {
			// Handle other database errors
			return fmt.Errorf("error checking category %s: %v", category.MainCategory, result.Error)
		} else {
			log.Printf("Skipping existing category: %s", category.MainCategory)
		}
	}
	return nil
}

func MustInit(mainHandler MainHandler) {
	for _, category := range PredefinedCategories {
		result := mainHandler.DB.Where("main_category = ?", category.MainCategory).First(&models.Category{})
		if result.Error != nil {
			if err := mainHandler.DB.Create(&category).Error; err != nil {
				log.Printf("Failed to create category %s: %v", category.MainCategory, err)
				continue
			}
			log.Printf("Created category: %s", category.MainCategory)
		}
	}
}
