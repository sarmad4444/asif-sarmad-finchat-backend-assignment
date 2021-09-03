package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sarmad4444/asif-sarmad-finchat-backend-assignment/customers"
	payments "github.com/sarmad4444/asif-sarmad-finchat-backend-assignment/payment"
)

func MapRoutes (router fiber.Router) {
	customers.MapRoutes(router)
	payments.MapRoutes(router)
}