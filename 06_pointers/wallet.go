package _6_pointers

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
