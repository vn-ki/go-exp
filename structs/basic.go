package structs

import (
	"errors"
)

// Wallet is your wallet. It contains your money which is int64.
// It also has a maximum amount it can contain.
type Wallet struct {
	money    int64
	MaxMoney int64
}

// Balance returns your balance.
// This should've been a value reciever.
// 	func (w Wallet) Balance() int64 {
// 		return w.money
// 	}
// But I think I heard somewhere that godoc suggest to use pointer receiver
// for consistency
func (w *Wallet) Balance() int64 {
	return w.money
}

// AddMoney adds money to your wallet
func (w *Wallet) AddMoney(money int64) error {
	if money += w.money; money > w.MaxMoney { // How is this readable code?
		return errors.New("Cannot add more than MaxMoney to Wallet")
	}

	w.money = money

	return nil
}

// TakeMoney takes money from your wallet
func (w *Wallet) TakeMoney(money int64) error {
	if money = w.money - money; money < 0 { // How is this readable code?
		return errors.New("You are broke")
	}

	w.money = money

	return nil
}
