package controllers

import (
	"net/http"
	"strings"

	"github.com/YohanJirage/Go_Lang/config"
	"github.com/YohanJirage/Go_Lang/models"
	"github.com/gofiber/fiber/v2"
)

func AddProduct(c *fiber.Ctx) error {
	var product models.Products

	if !strings.Contains(c.Path(), "/admin") {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized access"})
	}

	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if product.Name == "" || product.Description == "" || product.Price == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing required fields"})
	}

	err = config.Db.Create(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to add product"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "product added successfully", "product": product})
}

func ShowAllProducts(c *fiber.Ctx) error {

	if !strings.Contains(c.Path(), "/user") {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized access"})
	}

	var products []models.Products

	err := config.Db.Find(&products).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch products"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"product": products})
}
