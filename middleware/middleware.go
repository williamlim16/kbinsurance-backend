package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/williamlim16/kbinsurance-backend/util"
)

func IsAuthenticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unathenticated",
		})
	}
	return c.Next()
}
