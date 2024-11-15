package controllers

import (
	"github.com/VitorBonella/mindworks-rfp-completion-go/controllers/dto"
	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"github.com/gofiber/fiber/v2"
)


func CreateRFP(c *fiber.Ctx) error{

	//_, err := GetUser(c)
	//if err != nil{
	//	return err
	//}

	var data dto.RFP

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	rfp, err := models.NewRFP(data.Name,data.Requirements,data.Equipments,1)
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = database.CreateRFP(rfp)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	return c.JSON(rfp)

}

func ListRFP(c *fiber.Ctx) error {

	//_, err := GetUser(c)
	//if err != nil{
	//	return err
	//}

	rfps, err := database.ListRFP(1)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	return c.JSON(rfps)
}