package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Wallet deposit money", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("Wallet withdraw money", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)

		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
		assertNoError(t, err)
	})

	t.Run("Wallet withdraw insufficient  funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)

	})
}
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
