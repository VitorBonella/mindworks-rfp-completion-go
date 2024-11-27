package controllers

import (
	"strconv"

	"github.com/VitorBonella/mindworks-rfp-completion-go/database"
	"github.com/VitorBonella/mindworks-rfp-completion-go/models"
	"github.com/gofiber/fiber/v2"
)

func RFPResult(c *fiber.Ctx) error{

	queryValue := c.Query("id")

	id, err := strconv.Atoi(queryValue)
	if err != nil{
		return c.JSON(fiber.Map{
			"message": "Invalid Id",
		})
	}

	rfp, err := database.GetRFP(uint(id))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Can't Find RFP",
		})
	}

	if rfp.Status != models.RFPStatusFinished{
		return c.JSON(fiber.Map{
			"message": "RFP Unprocessed or Finished with Error",
		})
	}

	result_per_equipament := make(map[string]models.QuestionMap)
	
	for _,e := range rfp.Equipments{
		result, err := database.GetResults(rfp.Id,e.Id)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Error getting result",
			})
		}
		
		resultEquip := models.ConcatResults(result)
		if resultEquip != nil{
			result_per_equipament[e.Name] = *resultEquip
		}

	}
	return c.JSON(result_per_equipament)
}