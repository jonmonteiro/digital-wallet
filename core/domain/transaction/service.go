package transaction

import (
	"fmt"
	"github.com/jmonteiro/picpay-like/core/types"
)

type TransactionService struct {
	store       types.TransactionStore
	walletStore types.WalletStore
}

func NewTransactionService(store types.TransactionStore, walletStore types.WalletStore) *TransactionService {
	return &TransactionService{
		store:       store,
		walletStore: walletStore,
	}
}

func (s *TransactionService) MakeTransaction(fromUserID int, payload types.TransactionPayload) error {
	fromWallets, err := s.walletStore.GetWalletsByUserID(fromUserID)
	if err != nil {
		return err
	}
	if len(fromWallets) == 0 {
		return fmt.Errorf("user has no wallet")
	}
	fromWallet := fromWallets[0] 

	toWallet, err := s.walletStore.GetWalletByCardNumber(payload.TransactionKey)
	if err != nil {
		return err
	}

	err = s.walletStore.TransferBalance(fromWallet.ID, toWallet.ID, payload.Amount)
	if err != nil {
		return err
	}

	transaction := types.Transaction{
		FromWalletID: fromWallet.ID,
		ToWalletID:   toWallet.ID,
		Amount:       payload.Amount,
	}
	_, err = s.store.MakeTransaction(transaction)
	return err
}
