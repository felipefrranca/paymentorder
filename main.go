package main

import (
	"database/sql"
	"fmt"

	"github.com/felipefrranca/paymentorder/internal/infra/database"
	"github.com/felipefrranca/paymentorder/internal/usecases"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	paymentOrderRepository := database.NewPaymentOrderRepository(db)
	calculateFinalPrice := usecases.NewCalculateFinalPrice(paymentOrderRepository)

	request := usecases.PaymentOrderRequest{
		Id:     uuid.New().String(),
		Amount: 100,
		Tax:    10,
	}
	response, err := calculateFinalPrice.Execute(request)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(response.Id)

	getById := usecases.NewGetById(paymentOrderRepository)
	getByIdResponse, err := getById.Execute(usecases.GetByIdRequest{Id: response.Id})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(getByIdResponse)
}
