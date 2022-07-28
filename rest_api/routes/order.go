package routes

import (
	"errors"
	"rest_api/databse"
	"rest_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time         `json:"order_date"`
}

func CreateResponseOrder(orderModel models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{
		ID:        orderModel.ID,
		User:      user,
		Product:   product,
		CreatedAt: orderModel.CreatedAt,
	}
}

func CreateOrder(ctx *fiber.Ctx) error {
	var order models.Order
	err := ctx.BodyParser(&order)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var user models.User
	userErr := findUser(order.UserRefer, &user)
	if userErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(userErr.Error())
	}

	var product models.Product
	productErr := findProduct(order.ProductRefer, &product)
	if productErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(productErr.Error())
	}

	databse.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return ctx.Status(fiber.StatusOK).JSON(responseOrder)
}

func GetAllOrders(ctx *fiber.Ctx) error {
	orders := []models.Order{}
	databse.Database.Db.Find(&orders)
	responseOrders := []OrderSerializer{}
	for _, order := range orders {
		var user models.User
		var product models.Product
		databse.Database.Db.Find(&user, "id=?", order.UserRefer)
		databse.Database.Db.Find(&product, "id=?", order.ProductRefer)
		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
		responseOrders = append(responseOrders, responseOrder)
	}
	return ctx.Status(fiber.StatusOK).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	databse.Database.Db.Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var order models.Order
	orderErr := findOrder(id, &order)
	if orderErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(orderErr.Error())
	}

	var user models.User
	var product models.Product

	// databse.Database.Db.Find(&user, "id=?", order.UserRefer)
	// databse.Database.Db.Find(&product, "id=?", order.ProductRefer)
	// OR //
	databse.Database.Db.First(&user, order.UserRefer)
	databse.Database.Db.First(&product, order.ProductRefer)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return ctx.Status(fiber.StatusOK).JSON(responseOrder)
}
