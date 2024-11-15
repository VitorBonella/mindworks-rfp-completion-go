package main

import (
	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
        AllowOrigins:     "https://localhost",
        AllowCredentials: true,
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
    }))

	_ , _ = database.NewDBConnection()

	routes.Setup(app)
	

	app.Listen(":7756")
}
