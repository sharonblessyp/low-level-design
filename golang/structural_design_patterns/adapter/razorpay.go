package adapter

import "fmt"

type RazorPay struct {
}

func (r *RazorPay) ProcessPayment(orderId, amount int) {
	fmt.Printf("processing payment for order id: %d, amount: %d \n", orderId, amount)
}

func (r *RazorPay) GenerateInvoice() {
	fmt.Println("Generated invoice")
}
