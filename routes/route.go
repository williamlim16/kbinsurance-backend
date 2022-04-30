package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/williamlim16/kbinsurance-backend/controller"
	"github.com/williamlim16/kbinsurance-backend/middleware"
)

func Setup(app *fiber.App) {
	// app.Static("/", "./public")
	app.Static("/", "./public")

	app.Post("/api/login", controller.Login)
	app.Post("/api/register", controller.Register)

	app.Use(middleware.IsAuthenticate)
	app.Post("/api/attendance/checkin", controller.Checkin)                         //
	app.Put("/api/attendance/checkout", controller.Checkout)                        //
	app.Post("/api/attendances", controller.GetAttendances)                         //
	app.Post("/api/attendances/summary", controller.GetSummary)                     //
	app.Post("/api/attendance/overtime", controller.GetOvertime)                    //
	app.Post("/api/attendance/overtime/summary", controller.GetSummaryOvertime)     //
	app.Post("/api/attendance/late", controller.GetLate)                            //
	app.Post("/api/attendance/late/summary", controller.GetSummaryLate)             //
	app.Post("/api/attendance/earlyleave", controller.GetEarlyleave)                //
	app.Post("/api/attendance/earlyleave/summary", controller.GetSummaryEarlyleave) //
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
