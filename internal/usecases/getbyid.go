package usecases

import "github.com/felipefrranca/paymentorder/internal/entities"

type GetByIdRequest struct {
	Id string
}

type GetByIdResponse struct {
	Id         string
	Amount     float64
	Tax        float64
	FinalPrice float64
}

type GetById struct {
	PaymentOrderRepository entities.IPaymentOrderRepository
}

func NewGetById(paymentOrderRepository entities.IPaymentOrderRepository) *GetById {
	return &GetById{
		PaymentOrderRepository: paymentOrderRepository,
	}
}

func (c GetById) Execute(request GetByIdRequest) (*GetByIdResponse, error) {
	paymentOrder, err := c.PaymentOrderRepository.Get(request.Id)
	if err != nil {
		return nil, err
	}
	return &GetByIdResponse{
		Id:         paymentOrder.Id,
		Amount:     paymentOrder.Amount,
		Tax:        paymentOrder.Tax,
		FinalPrice: paymentOrder.FinalPrice,
	}, nil
}
