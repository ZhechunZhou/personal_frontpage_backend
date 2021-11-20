package common

import "github.com/stripe/stripe-go/v72"
const StoreServiceApi = "https://fakestoreapi.com"

func SetUpStrip()  {
	stripe.Key = "sk_key"
}


