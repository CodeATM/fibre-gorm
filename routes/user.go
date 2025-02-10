package routes

import (
	"errors"
	"strconv"

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

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ? ", id)

	if user.ID == 0 {
		return errors.New("User does not exist")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please ensure that :id is an integer",
		})
	}
	var user models.User

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponder((user))

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please ensure that :id is an integer",
		})
	}
	var user models.User

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateUser struct {
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
	}

	var updateData updateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.Firstname = updateData.Firstname
	user.Lastname = updateData.Lastname

	database.Database.Db.Save(&user)

	responseUser := CreateResponder(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please ensure that :id is an integer",
		})
	}
	var user models.User

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user); err != nil {
		c.Status(404).JSON(err.Error)
	}

	return c.Status(200).SendString("Successfully deleted a user")
}
