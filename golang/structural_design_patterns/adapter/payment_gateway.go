package adapter

type PaymentGateway interface {
	ProcessPayment(orderid, amount int)
	GenerateInvoice()
}
