package adapter

/*
1. payment gateway interface
2. implement methods for types
3. checkout service which decides the type
4. client using the service
*/

func main() {
	// Razorpay obj creation
	razorPay := &RazorPay{}
	checkout := NewCheckoutService(razorPay)
	checkout.Checkout(123, 2345)

	// process payment through paypal adapter
	payPalAdapter := NewPayPal()
	checkout = NewCheckoutService(payPalAdapter)
	checkout.Checkout(124, 2345)
}
