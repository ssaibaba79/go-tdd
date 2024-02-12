package pointers

import (
	"errors"
	"fmt"
)

type BitCoin float64

func (b BitCoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

var Error_not_enough_balance = errors.New("not enough balance")

type Wallet struct {
	balance BitCoin
}

func (w Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Deposit(deposit BitCoin) {
	w.balance += deposit
}

func (w *Wallet) Withdraw(withdraw BitCoin) error {
	if withdraw > w.balance {
		return Error_not_enough_balance
	}
	w.balance -= withdraw
	return nil
}
