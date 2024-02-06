package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	// test the address. the & symbol gets the pointer or mem address of a var
	// fmt.Printf("address of wallet in test is %p \n", &wallet.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("%#v got %d, wanted %d", wallet, got, want)
	}
}
