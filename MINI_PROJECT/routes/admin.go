package routes

import (
	"github.com/YohanJirage/Go_Lang/controllers"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {

	admin := app.Group("/admin")

	admin.Post("/CreateCategory", controllers.AddCategory)
	admin.Put("/UpdateCategory", controllers.EditCategory)
	admin.Delete("/DeleteCategory", controllers.DeleteCategory)
	admin.Get("/GetCategories", controllers.GetCategories)
	admin.Post("/addProduct", controllers.AddProduct)

}
