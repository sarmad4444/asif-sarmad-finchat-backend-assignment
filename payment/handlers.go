package payments

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func CreatePayment(ctx *fiber.Ctx) error {
	type PaymentDetails struct {
		StripeCustomerID string `json:"stripeCustomerID"`
		Amount int64 `json:"amount"`
	}

	var input PaymentDetails
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// Making create call to stripe
	params := &stripe.PaymentIntentParams{
		Customer: stripe.String(input.StripeCustomerID),
		Amount: stripe.Int64(input.Amount * 100), // Multiplying by 100 because stripe considers it as cents. Docs: https://stripe.com/docs/currencies#zero-decimal
		CaptureMethod: stripe.String("automatic"),
		Currency: stripe.String("USD"),
	}

	payment, err := paymentintent.New(params)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"paymentIntentID": payment.ID,
	})
}

// Fetch all payments made by a particular customer
func CustomerPayments(ctx *fiber.Ctx) error {
	stripeCustomerID := ctx.Params("customer_id")

	if stripeCustomerID == "" {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	params := &stripe.PaymentIntentListParams{
		Customer: stripe.String(stripeCustomerID),
	}

	// Fetching Payments
	paymentsList := paymentintent.List(params)


	// Parsing response to required format
	type paymentResponse struct {
		ID string `json:"id"`
		Amount int64 `json:"amount"`
	}

	results := [] paymentResponse{}

	for paymentsList.Next() {

		intent := paymentsList.PaymentIntent()

		parsedPayment := paymentResponse{
			ID: intent.ID,
			Amount: intent.Amount,
		}

		results = append(results, parsedPayment)
	}

	return ctx.JSON(fiber.Map{
		"data": results,
	})
}