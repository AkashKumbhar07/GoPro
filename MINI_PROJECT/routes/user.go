package routes

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App) {

	admin := app.Group("/admin")
	admin.Get("/GetCategories")

}
