package routes

import "github.com/gofiber/fiber/v2"

func AdminRoutes(app *fiber.App) {

	user := app.Group("/admin")
	user.Get("/GetCategories")

}
