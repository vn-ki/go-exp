package structs

import (
	"errors"
	"fmt"
)

// Wallet is your wallet. It contains your money which is int64.
// It also has a maximum amount it can contain.
type Wallet struct {
	Money    int64
	MaxMoney int64
}

func (w Wallet) String() string {
	return fmt.Sprintf("<Wallet Money=Rs. %d Max=Rs. %d>", w.Money, w.MaxMoney)
}

// Balance returns your balance.
// This should've been a value reciever.
// 	func (w Wallet) Balance() int64 {
// 		return w.money
// 	}
// But I think I heard somewhere that godoc suggest to use pointer receiver
// for consistency
func (w *Wallet) Balance() int64 {
	return w.Money
}

// AddMoney adds money to your wallet
func (w *Wallet) AddMoney(money int64) error {
	if money += w.Money; money > w.MaxMoney { // How is this readable code?
		return errors.New("Cannot add more than MaxMoney to Wallet")
	}

	w.Money = money

	return nil
}

// TakeMoney takes money from your wallet
func (w *Wallet) TakeMoney(money int64) error {
	if money = w.Money - money; money < 0 { // How is this readable code?
		return errors.New("You are broke")
	}

	w.Money = money

	return nil
}
