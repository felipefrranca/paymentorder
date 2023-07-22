package entities

type IPaymentOrderRepository interface {
	Save(paymentOrder *PaymentOrder) error
	Get(id string) (*PaymentOrder, error)
}
