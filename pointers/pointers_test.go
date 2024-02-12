package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got != nil && got.Error() != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("got %s want no error", got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		assertBalance(t, wallet, BitCoin(10.0))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20.0}
		err := wallet.Withdraw(BitCoin(15.0))
		assertNoError(t, err)
		assertBalance(t, wallet, BitCoin(5.0))
	})

	t.Run("Withdraw with no enough balance", func(t *testing.T) {
		wallet := Wallet{balance: 20.0}
		got := wallet.Withdraw(BitCoin(25.0))
		assertError(t, got, Error_not_enough_balance.Error())
	})

}
