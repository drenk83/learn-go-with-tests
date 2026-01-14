package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet.Balance(), Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		w := Wallet{startBalance}
		err := w.Withdraw(Bitcoin(100))

		assertBalance(t, w.Balance(), startBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if err != want {
		t.Errorf("got: %s, want: %s", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
