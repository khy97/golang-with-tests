package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin - type
type Bitcoin float64

// Wallet - Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit -  Deposits certain amount to Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds - Error for insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw -  Withdraws certain amount to Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance - Returns the balance of Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
