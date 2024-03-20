package controllers

import (
	"net/http"

	"github.com/YohanJirage/Go_Lang/config"
	"github.com/YohanJirage/Go_Lang/models"
	"github.com/gofiber/fiber/v2"
)

func AddCategory(context *fiber.Ctx) error {
	var category models.Categories
	err := context.BodyParser(&category)
	if err != nil {
		context.Status(http.StatusNotAcceptable).JSON(&fiber.Map{"message": "Failed to parse"})
	}

	err = config.Db.Create(&category).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed To Create new Category"})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Category Created", "data": category})
	return nil

}

func EditCategory(context *fiber.Ctx) error {
	var category models.Categories
	err := context.BodyParser(&category)
	if err != nil {
		context.Status(http.StatusNotAcceptable).JSON(&fiber.Map{"message": "Failed to parse"})
	}

	err = config.Db.Where("id = ? ", category.ID).Updates(&category).Error

	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed To Category Update"})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Category Updated", "data": category})
	return nil

}

func DeleteCategory(context *fiber.Ctx) error {
	var category models.Categories
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "PLEASE PROVIDE VALID ID, ID CANNOT BE EMPTY"})
		return nil
	}
	err := config.Db.Delete(category, id).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "PLEASE PROVIDE VALID ID "})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book Is Deletes"})

	return nil
}
func GetCategories(context *fiber.Ctx) error {
	categories := &[]models.Categories{}

	err := config.Db.Find(&categories).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Failed To Fetch All Categories"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Categories are successfully Fetched", "data": categories})

	return nil
}
