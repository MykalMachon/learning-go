package wallet

import (
	"errors"
	"fmt"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#wrapping-up
// key take-aways
// pointers: if you need to mutate state, use a pointer. If you don't want a copy, pass by reference with &
// errors: handle them gracefully, declare them and standardize them in your packages
// new types from existing ones: useful for adding functionality to existing types through methods. add domain specific context and interfaces.

// generic errors
var ErrInsufficientFunds = errors.New("insufficent funds to make withrdrawl")

// doing this just to make it descriptive?
// seems like it's obscuring the type for no real reason but alright
// edit: seems like this makes it easy to add methods
// to existing types; that makes sense actually.
type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// using func (w Wallet) takes a copy of wallet and works on that.
// using func (w *Wallet) uses a pointer to the called wallet.

func (w *Wallet) Deposit(amount Bitcoin) {
	// test the address: the & symbol gets the memory address of a var
	// fmt.Printf("address of wallet in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	// go uses automatic dereferncing, so we don't have to do it here
	// the code would be return (*w).balance but it's done automatically for us
	return w.balance
}
