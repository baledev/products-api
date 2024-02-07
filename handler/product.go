package handler

import (
	"baledev.com/products-api/database"
	"baledev.com/products-api/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllProducts query all products
// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {string} map[string]{}
// @Failure 503 {string} map[string]{}
// @Router /api/product [get]
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []models.Product
	db.Find(&products)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

// Get query product by ID given
// @Summary Get query product by ID given
// @Description Get query product by ID given
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {string} map[string]{}
// @Router /api/product/{id} [get]
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var product models.Product
	db.Find(&product, id)
	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

// Create new product
// @Summary Create new product
// @Description Create new product
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param product body models.ProductData true "Product"
// @Success 200 {string} map[string]{}
// @Router /api/product [post]
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create product",
			"data":    err,
		})
	}
	db.Create(&product)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created product",
		"data":    product,
	})
}

// DeleteProduct delete product
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product models.Product
	db.First(&product, id)
	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	db.Delete(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
