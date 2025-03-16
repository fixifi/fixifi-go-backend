package api

import (
	"github.com/fixifi/fixifi-go-backend/handlers"
)

//routes

func SetupRoute(mainHandler *handlers.MainHandler) {
	mainHandler.App.Get("/", mainHandler.WellComeHandler())
	// auth:=app.Group("/api/auth/")
	// auth.Get("/",)

	//categories
	categories := mainHandler.App.Group("/api/categories")
	categories.Get("/get", mainHandler.FetchCategories())
	categories.Post("/create", mainHandler.CreateCategories())
	categories.Put("/update/:id", mainHandler.UpdateCategory())

	//account
	account := mainHandler.App.Group("/api/account")
	account.Post("/consumer/create", mainHandler.CreateConsumer())
	account.Post("/provider/create", mainHandler.CreateProvider())
	// account.Post("/provider/create", mainHandler.CreateProvider())
}
