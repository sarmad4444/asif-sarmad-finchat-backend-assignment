package customers

import "github.com/gofiber/fiber/v2"

func MapRoutes(router fiber.Router) {
	route := router.Group("/customer")
	{
		route.Post("/", CreateCustomer)
		route.Post("/", CreateToken)
	}
}