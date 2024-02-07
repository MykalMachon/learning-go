package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("%#v got %s want %s", wallet, got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assetError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Expected error but did not recieve one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %q but wanted %q", got, want)
		}
	}

	t.Run("Deposit BTC", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw BTC", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw BTW with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))

		assetError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})
}
