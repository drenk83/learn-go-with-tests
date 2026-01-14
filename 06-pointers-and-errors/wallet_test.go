package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got: %s, want: %s", got, want)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		w := Wallet{Bitcoin(20)}
		err := w.Withdraw(Bitcoin(100))

		if w.Balance() != Bitcoin(20) {
			t.Errorf("balance has been changed")
		}

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
