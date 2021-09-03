package payments

import "github.com/gofiber/fiber/v2"

func MapRoutes(router fiber.Router) {
	route := router.Group("/payments")
	{
		route.Post("/", CreatePayment)
		route.Get("/:customer_id", CustomerPayments)
	}
}