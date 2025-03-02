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
	"github.com/fixifi/fixifi-go-backend/config"
	"github.com/fixifi/fixifi-go-backend/db/postgres"
	"github.com/gofiber/fiber/v3"
)

func main() {
	// Load configuration
	cfg := config.GetConfig()


	postgresDb := postgres.ConnectToDatabase(cfg)
	fmt.Println( postgresDb)


	app := fiber.New()

	api.SetupRoute(app)
	ok := make(chan os.Signal, 1)
	signal.Notify(ok, os.Interrupt, syscall.SIGINT, syscall.SIGABRT)

	go func() {
		if err := app.Listen(":8055"); err != nil {
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
