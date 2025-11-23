package wallet

import (
	"fmt"
	"github.com/jmonteiro/picpay-like/core/types"
)

type WalletService struct {
	store types.WalletStore
}

func NewWalletService(store types.WalletStore) *WalletService {
	return &WalletService{
		store: store,
	}
}

func (s *WalletService) CreateWallet(wallet types.Wallet) error {
	_, err := s.store.GetWalletBYCardNumber(wallet.CardNumber)
	if err == nil {
		return  fmt.Errorf("wallet with card number %s already exists", wallet.CardNumber)
	}
	return s.store.CreateWallet(wallet)
}

func (s *WalletService) GetWalletByID(id int) (*types.Wallet, error) {
	return s.store.GetWalletByID(id)
}

func (s *WalletService) GetWalletsByUserID(userID int) ([]*types.Wallet, error) {
	return s.store.GetWalletsByUserID(userID)
}

func (s *WalletService) UpdateCardNumber(userID int, walletID int, newCardNumber string) error {
	return s.store.UpdateCardNumber(userID, walletID, newCardNumber)
}