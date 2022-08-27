package pointers

import "testing"

func TestWallet(t *testing.T) {

  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(0)}
    wallet.Deposit(Bitcoin(10))

    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("withdrawal", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(20)}
    err := wallet.Withdraw(Bitcoin(10))

    assertNoError(t, err)
    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("withdraw insufficient funds", func(t *testing.T) {
    wallet := Wallet{balance: Bitcoin(2)}
    err := wallet.Withdraw(Bitcoin(3))

    assertError(t, err, ErrInsufficientFunds)
    assertBalance(t, wallet, Bitcoin(2))
  })
}


func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
  t.Helper()
  got := w.Balance()
  
  if got != want {
    t.Errorf("got %s, want %s", got, want)
  }
}

func assertNoError(t *testing.T, got error) {
  t.Helper()
  if got != nil {
    t.Errorf("got %s instead of no error", got)
  }
}

func assertError(t *testing.T, got, want error) {
  t.Helper()
  if got == nil {
    t.Error("didn't get an error, but wanted one")
  }

  if got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}
