package transaction

import (
	"database/sql"
	"github.com/jmonteiro/picpay-like/core/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) types.TransactionStore {
	return &Store{db: db}
}

func (s *Store) MakeTransaction(tx types.Transaction) (int64, error) {
	panic("unimplemented")
}

func (s *Store) GetTransactionByID(id int64) (*types.Transaction, error) {
	panic("unimplemented")
}

func (s *Store) ListUserTransactions(userID int64) ([]*types.Transaction, error) {
	panic("unimplemented")
}

