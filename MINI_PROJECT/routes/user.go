package routes

import (
	"github.com/YohanJirage/Go_Lang/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	user := app.Group("/admin")
	user.Get("/GetCategories", controllers.GetCategories)
	user.Get("/getAllProduct", controllers.ShowAllProducts)

}
