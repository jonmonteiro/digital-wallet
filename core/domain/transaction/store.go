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

func (s *Store) MakeTransaction(tx types.Transaction) (int, error) {
	row := s.db.QueryRow("INSERT INTO transactions (from_wallet_id, to_wallet_id, amount, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id", tx.FromWalletID, tx.ToWalletID, tx.Amount)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Store) GetTransactionByID(id int) (*types.Transaction, error) {
	panic("unimplemented")
}

func (s *Store) ListUserTransactions(userID int) ([]*types.Transaction, error) {
	panic("unimplemented")
}

