package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/williamlim16/kbinsurance-backend/database"
	"github.com/williamlim16/kbinsurance-backend/models"
	"github.com/williamlim16/kbinsurance-backend/util"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+]=\-]+@[a-z0-9. %+\-]`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	tempBody := struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		Phone     string
	}{}
	var userData models.User

	if err := c.BodyParser(&tempBody); err != nil {
		log.Println(&tempBody)
		log.Println(err)
		return err
	}
	if len(tempBody.Password) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 character",
		})
	}

	database.DB.Where("email=?", strings.TrimSpace(tempBody.Email)).First(&userData)

	if userData.ID != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	user := models.User{
		FirstName: tempBody.FirstName,
		LastName:  tempBody.LastName,
		Phone:     tempBody.Phone,
		Email:     tempBody.Email,
	}

	user.SetPassword(tempBody.Password)
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
	})
}

func Login(c *fiber.Ctx) error {

	tempBody := struct {
		Email    string
		Password string
	}{}

	if err := c.BodyParser(&tempBody); err != nil {
		fmt.Println("Unable to parse body")
	}

	var user models.User
	database.DB.Where("email=?", tempBody.Email).First(&user)

	if user.ID == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email doesn't exist",
		})
	}

	if err := user.ComparePassword(tempBody.Password); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Wrong password",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.ID)))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Login success",
		"user":    user,
		"token":   token,
	})

}

type Claims struct {
	jwt.StandardClaims
}
