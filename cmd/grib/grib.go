package main

import (
	"grib/internal/app/grib/router"
	"grib/internal/pkg/migration"
	"grib/pkg/shared/database"
	sharedLogger "grib/pkg/shared/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := sharedLogger.NewLogger()
	logger.Info(os.Environ())
	dbConfig := database.DBConfig{
		Host:    os.Getenv("DB_HOST"),
		Name:    os.Getenv("DB_NAME"),
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		Port:    os.Getenv("DB_PORT"),
		Charset: "utf8mb4",
	}
	logger.Info("Initializing database...")
	db, err := database.NewDB(dbConfig, logger)
	if err != nil {
		logger.Fatalln("Failed to connect database")
		panic(err)
	}
	logger.Info("Initialize database successfully")

	defer database.CloseDB(logger, db)

	logger.Info("Migrating database...")
	err = migration.Migrate(db)
	if err != nil {
		logger.Fatalln("Failed to migrate database")
	}
	logger.Info("Migrate database successfully")

	engine := gin.New()
	router := router.InitializeRouter(engine, db, logger)
	router.SetupHandler()

	engine.Run(":" + os.Getenv("PORT"))
}
