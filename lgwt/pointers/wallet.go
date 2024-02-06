package wallet

// doing this just to make it descriptive?
// seems like it's obscuring the type for no real reason but alright

// edit: seems like this makes it easy to add methods
// to existing types; that makes sense actually.
type Bitcoin int

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

func (w *Wallet) Balance() Bitcoin {
	// go uses automatic dereferncing, so we don't have to do it here
	// the code would be return (*w).balance but it's done automatically for us
	return w.balance
}
