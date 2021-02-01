package infrastructure

import (
	wallet "domain/wallet"
)

type MockWalletRepository struct {
	wallet []wallet.Wallet
}

func CreateMockWalletRepository() *MockWalletRepository {
	return &MockWalletRepository{
		wallet: []wallet.Wallet{},
	}
}

func (w *MockWalletRepository) store(wallet wallet.Wallet) {

}
func (w *MockWalletRepository) getAll() []wallet.Wallet {
	return w.wallet
}
