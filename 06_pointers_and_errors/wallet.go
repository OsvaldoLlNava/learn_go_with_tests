package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(cantidad Bitcoin) {
	w.balance += cantidad
}

func (w *Wallet) Withdraw(cantidad Bitcoin) error {

	if cantidad > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= cantidad
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
