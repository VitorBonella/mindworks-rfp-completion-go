package routes

import (
	"github.com/VitorBonella/mindworks-rfp-completion-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//AUTH ROUTES
	app.Get("/api/user", controllers.User)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	app.Put("/api/apikey", controllers.SetUserApiKey)

	//Equipment
	app.Get("/api/equipments", controllers.ListEquipment)
	app.Post("/api/equipment", controllers.CreateEquipment)
	app.Put("/api/equipment", controllers.UpdateEquipment)
	app.Delete("/api/equipment", controllers.DeleteEquipment)

	//RFP
	app.Post("/api/rfp", controllers.CreateRFP)
	app.Get("/api/rfps", controllers.ListRFP)
	app.Put("/api/rfp/reprocess", controllers.ReprocessRFP)

	//RFP Result
	app.Get("/api/rfp/result",controllers.RFPResult)
	app.Get("/api/rfp/result/old",nil)
}