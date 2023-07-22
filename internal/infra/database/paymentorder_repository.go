package database

import (
	"database/sql"

	"github.com/felipefrranca/paymentorder/internal/entities"
)

type PaymentOrderRepository struct {
	Db *sql.DB
}

func NewPaymentOrderRepository(db *sql.DB) *PaymentOrderRepository {
	return &PaymentOrderRepository{
		Db: db,
	}
}

func (r *PaymentOrderRepository) Save(paymentOrder *entities.PaymentOrder) error {
	_, err := r.Db.Exec("INSERT INTO paymentorder (id, amount, tax, final_price) VALUES ($1, $2, $3, $4)", paymentOrder.Id, paymentOrder.Amount, paymentOrder.Tax, paymentOrder.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *PaymentOrderRepository) Get(id string) (*entities.PaymentOrder, error) {
	row := r.Db.QueryRow("SELECT id, amount, tax, final_price FROM paymentorder WHERE id = $1", id)
	paymentOrder := entities.PaymentOrder{}
	err := row.Scan(&paymentOrder.Id, &paymentOrder.Amount, &paymentOrder.Tax, &paymentOrder.FinalPrice)
	if err != nil {
		return nil, err
	}
	return &paymentOrder, nil
}
