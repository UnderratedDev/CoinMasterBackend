package main

import (
	"github.com/stripe/stripe-go"
)

func main () {

	// Set your secret key: remember to change this to your live secret key in production
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = "sk_test_ks5dXRtdRdmb6Z2wy9AAnz3s"

	// Token is created using Checkout or Elements!
	// Get the payment token ID submitted by the form:
	token := r.FormValue("stripeToken")

	// Create a Customer:
	customerParams := &stripe.CustomerParams {
	  // Email: "paying.user@example.com",
	  Email : "rajyudhvir29@gmail.com",
	}
	customerParams.SetSource(token)
	customer, err := customer.New(customerParams)

	// Charge the Customer instead of the card:
	chargeParams := &stripe.ChargeParams{
	  Amount: 1000,
	  Currency: "usd",
	  Customer: customer.id,
	}
	charge, err := charge.New(chargeParams)

	// YOUR CODE: Save the customer ID and other info in a database for later.

	// YOUR CODE (LATER): When it's time to charge the customer again, retrieve the customer ID.
	chargeParams := &stripe.ChargeParams{
	  Amount: 1500, // $15.00 this time
	  Currency: "usd",
	  Customer: customerId,
	}
	charge, err := charge.New(chargeParams)

}