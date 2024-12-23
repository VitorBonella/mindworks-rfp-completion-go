package main

import (
	"os"

	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/routes"
	"github.com/VitorBonella/mindworks-rfp-completion-go/worker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	
	if os.Getenv("ENV") != "production"{
		app.Use(cors.New(cors.Config{
			AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowOrigins:     "https://orange-waddle-67v949p953x6pg-5173.app.github.dev/",
			AllowCredentials: true,
			AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		}))
	}
	_ , _ = database.NewDBConnection()

	routes.Setup(app)

	pageList := []string{"/",
						"/rfps",
						"/new_rfp",
						"/equipment",
						"/login",
						"/equipment",
						"/rfp_detail/:id"}

	if os.Getenv("ENV") == "production"{
		for _, p := range pageList{
			app.Static(p,"./client/dist")
		}
	}

	go worker.RunQueue()

	if os.Getenv("ENV") == "production"{
		app.Listen(":8000")
	} else{
		app.Listen(":5174")
	}
}
