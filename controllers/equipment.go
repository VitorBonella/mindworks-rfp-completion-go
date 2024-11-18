package controllers

import (
	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"github.com/gofiber/fiber/v2"
)

func ListEquipment(c *fiber.Ctx) error{

	user, err := GetUser(c)
	if err != nil || user == nil || user.Id == 0{
		return err
	}

	equipment, err := database.ListEquipment(user.Id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	return c.JSON(equipment)
}

func CreateEquipment(c *fiber.Ctx) error{

	user, err := GetUser(c)
	if err != nil || user == nil || user.Id == 0{
		return err
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	equipment,err := models.NewEquipment(data["name"],data["download_link"],user.Id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = database.CreateEquipment(equipment)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	return c.JSON(equipment)
}

func UpdateEquipment(c *fiber.Ctx) error{

	return c.JSON(fiber.Map{
		"message":"Not working yet",
	})
}

func DeleteEquipment(c *fiber.Ctx) error{
	var data map[string]int

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	
	id := data["id"]

	equipment := models.Equipment{
		Id: uint(id),
	}

	err := database.DeleteEquipment(&equipment)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "DB ERROR",
		})
	}

	return c.JSON(equipment)

}