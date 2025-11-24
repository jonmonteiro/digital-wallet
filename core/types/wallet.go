package types

type WalletStore interface {
	CreateWallet(wallet Wallet) error
	AddBalanceWithBankSlip(walletID int, amount float64) error
	GetWalletByID(id int) (*Wallet, error)
	GetWalletsByUserID(userID int) ([]*Wallet, error)
	GetWalletBYCardNumber(cardNumber string) (*Wallet, error)
	UpdateCardNumber(userID int, walletID int, newCardNumber string) error
}

type Wallet struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	CardNumber string  `json:"card_number"`
	Balance    float64 `json:"balance"`
}

type WalletPayload struct {
	CardNumber string `json:"card_number" validate:"required,len=16,numeric"`
}
