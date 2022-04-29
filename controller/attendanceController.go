package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/williamlim16/kbinsurance-backend/database"
	"github.com/williamlim16/kbinsurance-backend/models"
)

func Checkin(c *fiber.Ctx) error {
	tempBody := struct {
		UserID uint
	}{}

	if err := c.BodyParser(&tempBody); err != nil {
		// log.Println(&tempBody)
		log.Println(err)
		return err
	}

	var userData models.User

	database.DB.Where("id = ?", tempBody.UserID).First(&userData)

	if userData.ID == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	attendanceData := models.Attendance{
		UserID:  tempBody.UserID,
		ClockIn: time.Now().UnixMilli(),
	}

	err := database.DB.Create(&attendanceData)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "Success Clock In",
	})
}

func Checkout(c *fiber.Ctx) error {
	tempBody := struct {
		UserID uint
	}{}

	if err := c.BodyParser(&tempBody); err != nil {
		// log.Println(&tempBody)
		log.Println(err)
		return err
	}

	var userData models.User

	if err := database.DB.Where("id = ?", tempBody.UserID).First(&userData).Error; err != nil {

		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User does not exist",
		})

	}

	database.DB.Model(models.Attendance{}).Where("id = ?", tempBody.UserID).Update("clock_out", time.Now().UnixMilli())

	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "Success Clock Out",
	})
}
