package api

import (
	"github.com/fixifi/fixifi-go-backend/handlers"
	"github.com/gofiber/fiber/v3"
)

//routes


func SetupRoute(app *fiber.App) {
	app.Get("/", handlers.WellComeHandler)
	// auth:=app.Group("/api/auth/")
	// auth.Get("/",)
}
