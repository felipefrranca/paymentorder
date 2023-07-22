package usecases

import "github.com/felipefrranca/paymentorder/internal/entities"

type PaymentOrderRequest struct {
	Id     string
	Amount float64
	Tax    float64
}

type PaymentOrderResponse struct {
	Id         string
	Amount     float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	PaymentOrderRepository entities.IPaymentOrderRepository
}

func NewCalculateFinalPrice(paymentOrderRepository entities.IPaymentOrderRepository) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		PaymentOrderRepository: paymentOrderRepository,
	}
}

func (c CalculateFinalPrice) Execute(request PaymentOrderRequest) (*PaymentOrderResponse, error) {
	paymentOrder, err := entities.NewPaymentOrder(request.Id, request.Amount, request.Tax)
	if err != nil {
		return nil, err
	}
	err = paymentOrder.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.PaymentOrderRepository.Save(paymentOrder)
	if err != nil {
		return nil, err
	}
	return &PaymentOrderResponse{
		Id:         paymentOrder.Id,
		Amount:     paymentOrder.Amount,
		Tax:        paymentOrder.Tax,
		FinalPrice: paymentOrder.FinalPrice,
	}, nil
}
