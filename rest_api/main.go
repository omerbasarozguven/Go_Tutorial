package main

import (
	"log"
	"rest_api/databse"
	"rest_api/routes"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString("hello world")
}

func setupRoutes(app *fiber.App) {
	//welcome ep
	app.Get("/api", helloWorld)

	//user ep
	user_routes := app.Group("/api/user")
	user_routes.Post("/create", routes.CreateUser)
	user_routes.Get("/getAll", routes.GetAllUsers)
	user_routes.Get("/get/:id", routes.GetUser)
	user_routes.Post("update/:id", routes.UpdateUser)
	user_routes.Delete("/delete/:id", routes.DeleteUser)

	//product ep
	product_routes := app.Group("/api/product")
	product_routes.Post("/create", routes.Createproduct)
	product_routes.Get("/getAll", routes.GetAllProducts)
	product_routes.Get("/get/:id", routes.GetProduct)
	product_routes.Post("update/:id", routes.UpdateProduct)
	product_routes.Delete("/delete/:id", routes.DeleteProduct)

	// order ep
	order_routes := app.Group("api/order")
	order_routes.Post("/create", routes.CreateOrder)
	order_routes.Get("/getAll", routes.GetAllOrders)
	order_routes.Get("/get/:id", routes.GetOrder)

}

func main() {
	databse.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))

}
