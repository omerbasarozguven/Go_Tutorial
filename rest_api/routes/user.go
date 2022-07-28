package routes

import (
	"errors"
	"rest_api/databse"
	"rest_api/models"

	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {
	// serves as a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) UserSerializer {
	return UserSerializer{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.User
	err := ctx.BodyParser(&user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	databse.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return ctx.Status(fiber.StatusOK).JSON(responseUser)
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users := []models.User{}
	databse.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return ctx.Status(fiber.StatusOK).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	databse.Database.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	userErr := findUser(id, &user)
	if userErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(userErr.Error())
	}

	responseUser := CreateResponseUser(user)
	return ctx.Status(fiber.StatusOK).JSON(responseUser)

}

func UpdateUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	userErr := findUser(id, &user)
	if userErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(userErr.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser
	updateDataErr := ctx.BodyParser(&updateData)
	if updateDataErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(userErr.Error())
	}
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	databse.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)
	return ctx.Status(fiber.StatusOK).JSON(responseUser)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	userErr := findUser(id, &user)
	if userErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(userErr.Error())
	}

	deleteUserErr := databse.Database.Db.Delete(&user).Error
	if deleteUserErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(deleteUserErr.Error())
	}

	return ctx.Status(fiber.StatusOK).SendString("succesfully deleted user")

}
