package main

import (
	domain "domain/wallet"
	"fmt"
)

func main() {
	// idk why it doesnt work when in another package
	var walletRepository domain.IWalletRepository = domain.CreateMockWalletRepository()
	fmt.Println(walletRepository.getAll())

}
