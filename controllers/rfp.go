package controllers

import (
	"time"

	"github.com/VitorBonella/mindworks-rfp-completion-go/controllers/dto"
	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"github.com/gofiber/fiber/v2"
)


func CreateRFP(c *fiber.Ctx) error{

	user, err := GetUser(c)
	if err != nil || user == nil || user.Id == 0{
		return err
	}

	var data dto.RFP

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if len(data.Requirements) == 0{
		return c.JSON(fiber.Map{
			"message": "Empty Requirements",
		})
	}
	if len(data.Equipments) == 0{
		return c.JSON(fiber.Map{
			"message": "Empty Equipments",
		})
	}

	rfp, err := models.NewRFP(data.Name,data.Requirements,data.Equipments,user.Id)
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

	user, err := GetUser(c)
	if err != nil || user == nil || user.Id == 0{
		return err
	}

	rfps, err := database.ListRFP(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	return c.JSON(rfps)
}

func ReprocessRFP(c *fiber.Ctx) error {

	var data map[string]uint

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	rfp, err := database.GetRFP(data["id"])
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Can't Find RFP",
		})
	}

	err = database.DeleteResults(rfp.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	now := time.Now()
	rfp.CreationDate = &now
	rfp.EndDate = nil

	err = database.SetRFPStatus(rfp, models.RFPStatusCreated)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}