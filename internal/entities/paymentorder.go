package entities

import "errors"

type PaymentOrder struct {
	Id         string
	Amount     float64
	Tax        float64
	FinalPrice float64
}

func NewPaymentOrder(id string, amount float64, tax float64) (*PaymentOrder, error) {
	paymentOrder := &PaymentOrder{
		Id:     id,
		Amount: amount,
		Tax:    tax,
	}
	err := paymentOrder.Validate()
	if err != nil {
		return nil, err
	}
	return paymentOrder, nil
}

func (p *PaymentOrder) Validate() error {
	if p.Id == "" {
		return errors.New("id is required")
	}
	if p.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if p.Tax <= 0 {
		return errors.New("tax must be greater than or equal to zero")
	}
	return nil
}

func (p *PaymentOrder) CalculateFinalPrice() error {
	p.FinalPrice = p.Amount + p.Tax
	err := p.Validate()
	if err != nil {
		return err
	}
	return nil
}
