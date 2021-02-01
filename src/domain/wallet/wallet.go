package domain

import "errors"

type Wallet struct {
	owner Owner
	money Money
}

type WalletActions interface {
	Transfer(amount Money, receiver Wallet)
}

func CreateWallet(owner Owner, money Money) (*Wallet, error) {
	return &Wallet{
		owner: owner,
		money: money,
	}, nil
}

func (w *Wallet) Transfer(amount Money, receiver Wallet) {
	if w.money.usd-amount.usd < 0 {
		panic(errors.New("Not enough money in this wallet to perform this transaction"))
	}

	w.money.usd -= amount.usd
	receiver.receive(amount)

}

func (w *Wallet) receive(amount Money) {
	w.money.usd += amount.usd
}
