package adapter

import "fmt"

// Adaptee
type PayPalAPI struct {
}

// method with incompatible interface
func (p *PayPalAPI) makePayment(orderId, amount int) {
	fmt.Printf("processing payment for order id: %d, amount: %d", orderId, amount)
}
