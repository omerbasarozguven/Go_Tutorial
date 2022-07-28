package routes

import (
	"errors"
	"rest_api/databse"
	"rest_api/models"

	"github.com/gofiber/fiber/v2"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) ProductSerializer {
	return ProductSerializer{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func Createproduct(ctx *fiber.Ctx) error {
	var product models.Product
	err := ctx.BodyParser(&product)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	databse.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return ctx.Status(fiber.StatusOK).JSON(responseProduct)
}

func GetAllProducts(ctx *fiber.Ctx) error {
	products := []models.Product{}
	databse.Database.Db.Find(&products)
	responseProducts := []ProductSerializer{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return ctx.Status(fiber.StatusOK).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	databse.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil
}

func GetProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	var product models.Product
	productErr := findProduct(id, &product)
	if productErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(productErr.Error())

	}
	responseProduct := CreateResponseProduct(product)
	return ctx.Status(fiber.StatusOK).JSON(responseProduct)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	var product models.Product
	productErr := findProduct(id, &product)
	if productErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(productErr.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct
	updateDataErr := ctx.BodyParser(&updateData)
	if updateDataErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(updateDataErr.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	databse.Database.Db.Save(&product)
	responseProduct := CreateResponseProduct(product)
	return ctx.Status(fiber.StatusOK).JSON(responseProduct)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	var product models.Product
	productErr := findProduct(id, &product)
	if productErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(productErr.Error())
	}

	deleteProductErr := databse.Database.Db.Delete(&product).Error
	if deleteProductErr != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(deleteProductErr.Error())
	}
	return ctx.Status(fiber.StatusOK).SendString("succesfully deleted product")
}
