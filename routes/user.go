package routes

import (
	database "github.com/CodeATM/fibre-gorm/Database"
	"github.com/CodeATM/fibre-gorm/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	// this is a serializer
	ID        uint   `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"Last_name"`
}

func CreateResponder(userModel models.User) User {
	return User{
		ID: userModel.ID, Firstname: userModel.Firstname, Lastname: userModel.Lastname,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponder((user))

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	// Fetch all users from the database
	database.Database.Db.Find(&users)

	// Prepare a slice to hold the response users
	responseUsers := []User{}

	// Convert each user to the response format and append it to responseUsers
	for _, user := range users {
		responseUser := CreateResponder(user)
		responseUsers = append(responseUsers, responseUser)
	}

	// Return the JSON response with a 200 status code
	return c.Status(200).JSON(responseUsers)
}
