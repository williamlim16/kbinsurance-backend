package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/williamlim16/kbinsurance-backend/database"
	"github.com/williamlim16/kbinsurance-backend/routes"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3006",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	routes.Setup(app)
	app.Listen(":" + port)
}
