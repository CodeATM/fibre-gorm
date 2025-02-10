package routes

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	database "github.com/CodeATM/fibre-gorm/Database"
	"github.com/CodeATM/fibre-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json: "product_name"`
	SerialNumber string `json: "serial_number"`
}

func CreateProductResponder(ProductModel models.Product) Product {
	return Product{
		ID: ProductModel.ID, Name: ProductModel.Name, SerialNumber: ProductModel.SerialNumber,
	}
}

func generateSerialNumber() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 100 and 999
	randomNumber := rand.Intn(900) + 100

	// Combine the prefix with the random number
	return fmt.Sprintf("adc-%d", randomNumber)
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	product.SerialNumber = generateSerialNumber()

	database.Database.Db.Create(&product)

	response := CreateProductResponder(product)
	return c.Status(200).JSON(response)
}

func GetProducts(c *fiber.Ctx) error {
	// Fetch all products from the database
	products := []models.Product{}
	database.Database.Db.Find(&products)

	// Prepare the response slice
	responseProduct := []Product{}

	// Transform each product into the response format
	for _, product := range products {
		response := CreateProductResponder(product)
		responseProduct = append(responseProduct, response)
	}

	// Return the response as JSON
	return c.Status(200).JSON(responseProduct)
}

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("Product does not exist")
	}

	return nil
}

func GetProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please ensure that :id is an integer",
		})
	}

	var product models.Product

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := CreateProductResponder(product)

	return c.Status(200).JSON(response)
}

func UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please ensure that :id is an integer",
		})
	}

	var product models.Product

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name string `json:"product_name"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = updateData.Name
	database.Database.Db.Save(&product)
	response := CreateProductResponder(product)
	return c.Status(200).JSON(response)
}

func DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please ensure that :id is an integer",
		})
	}
	var product models.Product

	if err := FindUser(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product); err != nil {
		c.Status(404).JSON(err.Error)
	}

	return c.Status(200).SendString("Successfully deleted a product")
}
