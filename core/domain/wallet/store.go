package wallet

import (
	"database/sql"

	"github.com/jmonteiro/picpay-like/core/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) types.WalletStore {
	return &Store{db: db}
}

func (s *Store) CreateWallet(wallet types.Wallet) error {
	_, err := s.db.Exec("INSERT INTO wallets (user_id, card_number, balance) VALUES ($1, $2, $3)", wallet.UserID, wallet.CardNumber, wallet.Balance)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetWalletByID(id int) (*types.Wallet, error) {
	row := s.db.QueryRow(`SELECT id, user_id, card_number, balance FROM wallets WHERE id=$1`, id)

	var w types.Wallet
	err := row.Scan(&w.ID, &w.UserID, &w.CardNumber, &w.Balance)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (s *Store) GetWalletsByUserID(userID int) ([]*types.Wallet, error) {
	rows, err := s.db.Query(`SELECT id, user_id, card_number, balance FROM wallets WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []*types.Wallet
	for rows.Next() {
		var w types.Wallet
		err := rows.Scan(&w.ID, &w.UserID, &w.CardNumber, &w.Balance)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, &w)
	}

	return wallets, nil
}

func (s *Store) GetWalletBYCardNumber(cardNumber string) (*types.Wallet, error) {
	row := s.db.QueryRow(`SELECT id, user_id, card_number, balance FROM wallets WHERE card_number=$1`, cardNumber)

	var w types.Wallet
	err := row.Scan(&w.ID, &w.UserID, &w.CardNumber, &w.Balance)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (s *Store) UpdateCardNumber(userID int, walletID int, newCardNumber string) error {
	_, err := s.db.Exec("UPDATE wallets SET card_number=$1 WHERE id=$2 AND user_id=$3", newCardNumber, walletID, userID)
	return err
}
