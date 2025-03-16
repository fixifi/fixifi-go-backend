package database

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/fixifi/fixifi-go-backend/config"
	"github.com/fixifi/fixifi-go-backend/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	Db *gorm.DB
}

func ConnectToDatabase(cfg *config.Config) *Postgres {
	dsnWithoutDB := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
		cfg.Host, cfg.Port, cfg.DbUsername, cfg.DbPassword, cfg.DbSslMode,
	)

	// Connect to PostgreSQL (without database)
	tempDB, err := gorm.Open(postgres.Open(dsnWithoutDB), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL (without DB): %v", err)
	}

	// Check if the database exists
	var exists bool
	checkDBSQL := fmt.Sprintf("SELECT EXISTS (SELECT FROM pg_database WHERE datname = '%s');", cfg.DbName)
	if err := tempDB.Raw(checkDBSQL).Scan(&exists).Error; err != nil {
		log.Fatalf("Error checking database existence: %v", err)
	}

	// Create database if it does not exist
	if !exists {
		createDBSQL := fmt.Sprintf("CREATE DATABASE %s;", cfg.DbName)
		if err := tempDB.Exec(createDBSQL).Error; err != nil {
			log.Fatalf("Failed to create database %s: %v", cfg.DbName, err)
		}
		slog.Info("Database created successfully !")
	} else {
		slog.Info("Database already exists, proceeding with connection.")
	}

	// DSN with database name (for actual connection)
	dsnWithDB := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DbUsername, cfg.DbPassword, cfg.DbName, cfg.DbSslMode,
	)

	// Connect to the actual database
	db, err := gorm.Open(postgres.Open(dsnWithDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database %s: %v", cfg.DbName, err)
	}

	slog.Info("Database connected successfully!")
	postgres := Postgres{Db: db}

	var tables []interface{}
	tables = append(tables, &models.Consumer{})
	tables = append(tables, &models.Category{})
	tables = append(tables, &models.Address{})
	tables = append(tables, &models.Category{})
	tables = append(tables, &models.Order{})
	tables = append(tables, &models.OrderImage{})
	tables = append(tables, &models.Review{})
	tables = append(tables, &models.Provider{})
	tables = append(tables, &models.Equipment{})
	tables = append(tables, &models.Provider{})

	err = postgres.CreateAllTables(tables)
	if err != nil {
		log.Fatalf("Failed to migrate tables: %v", err)
	}
	return &postgres
}
