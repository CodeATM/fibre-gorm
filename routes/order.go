package routes

import (
	"time"

	database "github.com/CodeATM/fibre-gorm/Database"
	"github.com/CodeATM/fibre-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser((&order)); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)
	resoponseUser := CreateResponder(user)
	resonseProduct := CreateProductResponder(product)
	responseOrder := CreateResponseOrder(order, resoponseUser, resonseProduct)

	return c.Status(200).JSON(responseOrder)
}
