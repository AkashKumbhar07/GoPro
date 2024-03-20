package main

import (
	"log"

	"github.com/YohanJirage/Go_Lang/config"
	"github.com/YohanJirage/Go_Lang/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {

	routes.AdminRoutes(app)
	routes.UserRoutes(app)

}

func main() {

	err := godotenv.Load(".env")
	errorCheck(err)

	config.Db, err = config.DB_Connect()
	if err != nil {
		log.Fatal("Failed To Coonect")
	}

	r := Repository{
		DB: config.Db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

	errorCheck(err)

}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
