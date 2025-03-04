package api

import (
	"github.com/fixifi/fixifi-go-backend/handlers"
)

//routes

func SetupRoute(mainHandler *handlers.MainHandler) {
	mainHandler.App.Get("/", mainHandler.WellComeHandler())
	// auth:=app.Group("/api/auth/")
	// auth.Get("/",)
	account := mainHandler.App.Group("/api/account")
	account.Post("/consumer/create", mainHandler.CreateConsumer())
	// account.Post("/provider/create", mainHandler.CreateProvider())
}
