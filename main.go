package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"kirmac-site-backend/common/app"
	"kirmac-site-backend/common/postgresql"
	"kirmac-site-backend/controller"
	"kirmac-site-backend/persistence"
	"kirmac-site-backend/services"
)

func main() {
	ctx := context.Background()
	c := fiber.New()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	propertyRepository := persistence.NewPropertyRepository(dbPool)

	propertyService := services.NewPropertyService(propertyRepository)

	propertyController := controller.NewPropertyController(propertyService)

	propertyController.RegisterRoutes(c)

	c.Listen(":8080")
}
