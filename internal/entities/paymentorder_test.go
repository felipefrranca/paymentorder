package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_Id_Is_Blank(t *testing.T) {
	paymentOrder := PaymentOrder{}
	assert.Error(t, paymentOrder.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Amount_Is_Blank(t *testing.T) {
	paymentOrder := PaymentOrder{Id: "123"}
	assert.Error(t, paymentOrder.Validate(), "amount must be greater than zero")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	paymentOrder := PaymentOrder{Id: "123", Amount: 10}
	assert.Error(t, paymentOrder.Validate(), "tax must be greater than or equal to zero")
}

func Test_If_Final_Price_Is_Correct(t *testing.T) {
	paymentOrder := PaymentOrder{Id: "123", Amount: 10, Tax: 1}
	assert.NoError(t, paymentOrder.Validate(), "No error expected")
	paymentOrder.CalculateFinalPrice()
	assert.Equal(t, 11.0, paymentOrder.FinalPrice, "Final price is not correct")
}
