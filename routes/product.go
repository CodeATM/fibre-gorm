package routes

import (
	"fmt"
	"math/rand"
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
