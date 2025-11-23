package user

import (
	"fmt"

	"github.com/jmonteiro/picpay-like/core/config"
	middelware "github.com/jmonteiro/picpay-like/core/middleware/auth"
	"github.com/jmonteiro/picpay-like/core/types"
)

type UserService struct {
	store       types.UserStore
	walletStore types.WalletStore
}

func NewUserService(store types.UserStore, walletStore types.WalletStore) *UserService {
	return &UserService{
		store:       store,
		walletStore: walletStore,
	}
}

func (s *UserService) RegisterUser(payload types.RegisterUserPayload) error {
	_, err := s.store.GetUserByEmail(payload.Email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", payload.Email)
	}

	hashedPassword, err := middelware.HashPassword(payload.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user := types.User{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	err = s.store.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Buscar o usuário recém-criado para obter o ID
	createdUser, err := s.store.GetUserByEmail(payload.Email)
	if err != nil {
		return fmt.Errorf("user created but failed to retrieve ID: %w", err)
	}

	// Criar wallet automaticamente com saldo inicial zero
	wallet := types.Wallet{
		UserID:  createdUser.ID,
		Balance: 0.0,
	}

	err = s.walletStore.CreateWallet(wallet)
	if err != nil {
		// Mesmo se falhar a criação da wallet, o usuário foi criado
		// Poderia fazer rollback aqui se necessário
		return fmt.Errorf("user created but failed to create wallet: %w", err)
	}

	return nil
}

func (s *UserService) LoginUser(payload types.LoginUserPayload) (string, error) {
	u, err := s.store.GetUserByEmail(payload.Email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	if !middelware.ComparePasswords(u.Password, []byte(payload.Password)) {
		return "", fmt.Errorf("invalid email or password")
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := middelware.CreateJWT(secret, int(u.ID))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	return token, nil
}

func (s *UserService) GetUsers() ([]*types.User, error) {
	users, err := s.store.GetUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil
}

func (s *UserService) GetUserByID(id int) (*types.User, error) {
	user, err := s.store.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.store.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s *UserService) UpdateUser(id int, payload types.RegisterUserPayload) error {
	_, err := s.store.GetUserByID(id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	hashedPassword, err := middelware.HashPassword(payload.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	payload = types.RegisterUserPayload{
		Email:    payload.Email,
		Password: hashedPassword,
	}

	err = s.store.UpdateUser(payload, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
