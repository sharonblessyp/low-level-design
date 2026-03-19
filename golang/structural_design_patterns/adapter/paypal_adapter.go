package adapter

import "fmt"

// Adapter
type PayPal struct {
	payPalAPI *PayPalAPI
}

// composite behavior
func NewPayPal() *PayPal {
	return &PayPal{payPalAPI: &PayPalAPI{}}
}

func (p *PayPal) ProcessPayment(orderId, amount int) {
	p.payPalAPI.makePayment(orderId, amount)
	return
}

func (p *PayPal) GenerateInvoice() {
	fmt.Println("Generated invoice")
}
