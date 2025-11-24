package types

import "time"

type TransactionStore interface {
	MakeTransaction(tx Transaction) (int, error)
	GetTransactionByID(id int) (*Transaction, error)
	ListUserTransactions(userID int) ([]*Transaction, error)
}

type Transaction struct {
	ID           int       `json:"id"`
	FromWalletID int       `json:"from_wallet_id"`
	ToWalletID   int       `json:"to_wallet_id"`
	Amount       float64   `json:"amount"`
	CreatedAt    time.Time `json:"created_at"`
}

type TransactionPayload struct {
	TransactionKey string  `json:"transaction_key" validate:"required,len=16,numeric"`
	Amount         float64 `json:"amount" validate:"required,gt=0"`
}
