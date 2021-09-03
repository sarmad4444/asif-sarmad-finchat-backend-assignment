package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarmad4444/asif-sarmad-finchat-backend-assignment/core"
	"github.com/stripe/stripe-go/v72"
)

func init() {
	stripe.Key = "sk_test_51JSglkJ0UKK3DLecMf6izcBRqf5WjpI4oScoxGHzNZ1oSsgZ30T0QkwdiOpxjM8T2ugiJFhRWSX2amQOi87olrTR00qJctcfcd"
}

func main() {
	app := fiber.New()

	api := app.Group("/api")

	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("API Up and Running.")
	})

	core.MapRoutes(api)

	_ = app.Listen(":3000")
}