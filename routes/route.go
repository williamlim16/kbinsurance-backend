package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/williamlim16/kbinsurance-backend/controller"
)

func Setup(app *fiber.App) {
	// app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hellow wordldfsjafklds;")
	})

	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.Register)

	app.Post("/api/attendance/checkin", controller.Checkin)  //
	app.Put("/api/attendance/checkout", controller.Checkout) //
	// app.Use(middleware.IsAuthenticate)
	// app.Get("/api/trashcans", controller.GetTrashCan)
	// app.Get("/api/trashcan/:id/edit", controller.EditTrashCan)
	// app.Post("/api/trashcan", controller.CreateTrashCan)
	// app.Put("/api/trashcan/:id", controller.UpdateTrashCan)
	// app.Delete("/api/trashcan/:id", controller.DeleteTrashCan)
	// app.Get("/api/trash", controller.GetTrash)
	// app.Post("/api/trash", controller.CreateTrash)
	// app.Put("/api/trash/:id", controller.UpdateTrash)
	// app.Delete("/api/trash/:id", controller.DeleteTrash)
	// app.Post("/api/trash/testing", controller.Register2)
}
