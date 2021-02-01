package domain

type IWalletRepository interface {
	store(wallet Wallet)
	getAll() []Wallet
}

type MockWalletRepository struct {
	wallet []Wallet
}

func CreateMockWalletRepository() *MockWalletRepository {
	return &MockWalletRepository{
		wallet: []Wallet{},
	}
}

func (w *MockWalletRepository) store(wallet Wallet) {

}

func (w *MockWalletRepository) getAll() []Wallet {
	return w.wallet
}
