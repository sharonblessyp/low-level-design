package adapter

// client
type CheckoutService struct {
	gateway PaymentGateway
}

func NewCheckoutService(paymentGateway PaymentGateway) *CheckoutService {
	return &CheckoutService{gateway: paymentGateway}
}

func (c *CheckoutService) Checkout(orderID, amount int) {
	c.gateway.ProcessPayment(orderID, amount)
}
