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

	database.DB.Model(models.Attendance{}).Where("user_id = ?", tempBody.UserID).Where("clock_out = 0").Update("clock_out", time.Now().UnixMilli())

	c.Status(200)

	return c.JSON(fiber.Map{
		"message": "Success Clock Out",
	})
}

func GetAttendances(c *fiber.Ctx) error {

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		// log.Println(&tempBody)
		log.Println(err)
		return err
	}
	var results []models.Attendance
	database.DB.Table("attendances").Order("id desc").Where("user_id = ?", tempBody.UserID).Find(&results)

	return c.JSON(fiber.Map{
		"message": "Success",
		"data":    results,
	})

}

func GetSummary(c *fiber.Ctx) error {
	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		// log.Println(&tempBody)
		log.Println(err)
		return err
	}

	//Overtime
	var resultsOvertime []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_out/1000)) = ?", 17).Where("DATE_PART('minute', to_timestamp(clock_out/1000)) > 0").Or("DATE_PART('hour', to_timestamp(clock_out/1000)) > 17").Where("user_id = ?", tempBody.UserID).Find(&resultsOvertime)

	var summaryOvertime int

	for _, result := range resultsOvertime {
		summaryOvertime += (time.Unix(0, result.ClockOut*int64(time.Millisecond)).Hour()-17)*60 + (time.Unix(0, result.ClockOut*int64(time.Millisecond)).Minute())
	}

	//Late
	var resultsLate []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_in/1000)) >= ?", 8).Where("DATE_PART('minute', to_timestamp(clock_in/1000)) > 0").Where("user_id = ?", tempBody.UserID).Find(&resultsLate)
	var summaryLate int

	for _, result := range resultsLate {
		summaryLate += ((time.Unix(0, result.ClockIn*int64(time.Millisecond)).Hour()-8)*60 + (time.Unix(0, result.ClockIn*int64(time.Millisecond)).Minute()))
	}

	//Early leave
	var resultsEarly []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_out/1000)) < ?", 17).Where("user_id = ?", tempBody.UserID).Find(&resultsEarly)

	var summaryEarly int

	for _, result := range resultsEarly {
		// log.Println(time.Unix(0, result.ClockOut*int64(time.Millisecond)))
		if time.Unix(0, result.ClockOut*int64(time.Millisecond)).Minute() > 0 {
			summaryEarly += ((16-time.Unix(0, result.ClockOut*int64(time.Millisecond)).Hour())*60 + (60 - time.Unix(0, result.ClockOut*int64(time.Millisecond)).Minute()))
		} else {

			summaryEarly += ((17 - time.Unix(0, result.ClockOut*int64(time.Millisecond)).Hour()) * 60)
		}

	}

	return c.JSON(fiber.Map{
		"overtime": summaryOvertime,
		"early":    summaryEarly,
		"late":     summaryLate,
	})

}

func GetOvertime(c *fiber.Ctx) error {
	// var attd []models.Attendance

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		// log.Println(&tempBody)
		log.Println(err)
		return err
	}

	var results []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_out/1000)) = ?", 17).Where("DATE_PART('minute', to_timestamp(clock_out/1000)) > 0").Or("DATE_PART('hour', to_timestamp(clock_out/1000)) > 17").Where("user_id = ?", tempBody.UserID).Find(&results)

	return c.JSON(fiber.Map{
		"data": results,
	})
}

func GetSummaryOvertime(c *fiber.Ctx) error {
	// var attd []models.Attendance

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		// log.Println(&tempBody)
		log.Println(err)
		return err
	}

	var results []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_out/1000)) = ?", 17).Where("DATE_PART('minute', to_timestamp(clock_out/1000)) > 0").Or("DATE_PART('hour', to_timestamp(clock_out/1000)) > 17").Where("user_id = ?", tempBody.UserID).Find(&results)

	var summary int

	for _, result := range results {
		summary += (time.Unix(0, result.ClockOut*int64(time.Millisecond)).Hour()-17)*60 + (time.Unix(0, result.ClockOut*int64(time.Millisecond)).Minute())
	}

	return c.JSON(fiber.Map{
		"data": summary,
	})
}

func GetLate(c *fiber.Ctx) error {

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		log.Println(err)
		return err
	}

	var results []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_in/1000)) >= ?", 8).Where("DATE_PART('minute', to_timestamp(clock_in/1000)) > 0").Where("user_id = ?", tempBody.UserID).Find(&results)

	return c.JSON(fiber.Map{
		"data": results,
	})

}

func GetSummaryLate(c *fiber.Ctx) error {

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		log.Println(err)
		return err
	}

	var results []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_in/1000)) >= ?", 8).Where("DATE_PART('minute', to_timestamp(clock_in/1000)) > 0").Where("user_id = ?", tempBody.UserID).Find(&results)
	var summary int

	for _, result := range results {
		summary += ((time.Unix(0, result.ClockIn*int64(time.Millisecond)).Hour()-8)*60 + (time.Unix(0, result.ClockIn*int64(time.Millisecond)).Minute()))
	}

	return c.JSON(fiber.Map{
		"data": summary,
	})

}

func GetEarlyleave(c *fiber.Ctx) error {

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		log.Println(err)
		return err
	}

	var results []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_out/1000)) < ?", 17).Where("user_id = ?", tempBody.UserID).Find(&results)

	return c.JSON(fiber.Map{
		"data": results,
	})

}

func GetSummaryEarlyleave(c *fiber.Ctx) error {

	tempBody := struct {
		UserID uint
	}{}
	if err := c.BodyParser(&tempBody); err != nil {
		log.Println(err)
		return err
	}

	var results []models.Attendance
	database.DB.Table("attendances").Where("DATE_PART('hour', to_timestamp(clock_out/1000)) < ?", 17).Where("user_id = ?", tempBody.UserID).Find(&results)

	var summary int

	for _, result := range results {
		// log.Println(time.Unix(0, result.ClockOut*int64(time.Millisecond)))
		if time.Unix(0, result.ClockOut*int64(time.Millisecond)).Minute() > 0 {
			summary += ((16-time.Unix(0, result.ClockOut*int64(time.Millisecond)).Hour())*60 + (60 - time.Unix(0, result.ClockOut*int64(time.Millisecond)).Minute()))
		} else {

			summary += ((17 - time.Unix(0, result.ClockOut*int64(time.Millisecond)).Hour()) * 60)
		}

	}
	return c.JSON(fiber.Map{
		"data": summary,
	})

}
