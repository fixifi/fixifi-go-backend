package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fixifi/fixifi-go-backend/cmd/api"
	"github.com/fixifi/fixifi-go-backend/cmd/initconst"

	"github.com/fixifi/fixifi-go-backend/config"
	database "github.com/fixifi/fixifi-go-backend/db/postgres"
	"github.com/fixifi/fixifi-go-backend/handlers"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Load configuration
	cfg := config.GetConfig()
	postgresDb := database.ConnectToDatabase(cfg)
	fmt.Println(postgresDb)
	app := fiber.New()

	mainHandler := handlers.MainHandler{
		DB:       postgresDb.Db,
		Validate: validator.New(),
		App:      app,
		Cfg:      cfg,
	}

	// Initialize predefined categories
	// initconst.MustInit(&mainHandler)
	initconst.MustInit(&mainHandler)

	api.SetupRoute(&mainHandler)

	// -------->Graceful shutdown
	ok := make(chan os.Signal, 1)
	signal.Notify(ok, os.Interrupt, syscall.SIGINT, syscall.SIGABRT)
	addr := fmt.Sprintf(":%d", 8055)
	go func() {
		slog.Info("Server starting at", "address", addr)

		if err := app.Listen(addr); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	<-ok
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	slog.Info("Server gracefully stopped.")
}
