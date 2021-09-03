package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/token"
)

// Create a new customer in stripe
func CreateCustomer (ctx *fiber.Ctx) error {
	type CustomerDetails struct {
		Email string `json:"email" required:"true"`
		StripeCreditCardToken string `json:"stripeCreditCardToken"`
	}

	var input CustomerDetails
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// Making create call on stripe
	params := &stripe.CustomerParams{
		Email: stripe.String(input.Email),
		Token: stripe.String(input.StripeCreditCardToken),
	}

	cust, err := customer.New(params)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"stripeCustomerID": cust.ID,
	})
}

// Create a testing card token to be used by other endpoints
func CreateToken(ctx *fiber.Ctx) error {
	tokenParams := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String("4242424242424242"),
			ExpMonth: stripe.String("12"),
			ExpYear: stripe.String("2025"),
			CVC: stripe.String("123"),
		},
	}
	tokenMeta, err := token.New(tokenParams)
	if err != nil {
		return err
	}

	return ctx.JSON(tokenMeta)
}