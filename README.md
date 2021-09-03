# Basic Example for using Fiber and Stripe-Go for Fintech

#### Libraries
- **Fiber** - https://github.com/gofiber/fiber
As REST API framework

- **Stripe Go** - https://github.com/stripe/stripe-go/
The official SDK from stripe to interact with their API

- **Air** - https://github.com/cosmtrek/air
Hot reloading library for development ease.

## How to run this project
Update Stripe secret key in `server.go` file and execute the following commands in project's root directory. 

- `go install`
- `go mod vendor`
- `go run server.go` 

## Available Endpoints
- **POST**: `api/customer` - Create a new customer. Params:
  - `email`: Email of the customer
  - `stripeCreditCardToken`: Default card token to be used for this customer
- **POST**: `api/payments` - Create a payment_intent for a customer. Params:
  - `stripeCustomerID`: Customer ID in stripe
  - `amount`: Amount in USD
- **GET**: `api/payments/:stripe_customer_id` - Get all payments made by a customer.