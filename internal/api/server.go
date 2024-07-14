package api

import (
	"log"
	"net/http"

	"github.com/Emmanuel-MacAnThony/gocommerce/config"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/api/rest"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/api/rest/handlers"
	"github.com/Emmanuel-MacAnThony/gocommerce/internal/domain"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()
	app.Get("/health", HealthCheck)
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection failed %v \n", err)
	}
	log.Println("database connected 🚀🚀🚀")

	// run migration
	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}
	setupRoutes(rh)
	app.Listen(config.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "I am breathing",
	})

}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}
