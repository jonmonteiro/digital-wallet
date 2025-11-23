package types

import "time"

type TransactionStore interface {
	CreateTransaction(tx Transaction) (int64, error)
	GetTransactionByID(id int64) (*Transaction, error)
	ListUserTransactions(userID int64) ([]*Transaction, error)
}

type Transaction struct {
	ID              int64     `json:"id"`
	FromUserID      int64     `json:"from_user_id"`
	ToUserID        int64     `json:"to_user_id"`
	Amount          float64   `json:"amount"`
	Status          string    `json:"status"` // pending, approved, rejected
	CreatedAt       time.Time `json:"created_at"`
	CompletedAt     *time.Time `json:"completed_at,omitempty"`
}

type CreateTransactionPayload struct {
	FromUserID int64   `json:"from_user_id" validate:"required"`
	ToUserID   int64   `json:"to_user_id" validate:"required"`
	Amount     float64 `json:"amount" validate:"required,gt=0"`
}