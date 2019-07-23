package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, expect Bitcoin) {
		actual := wallet.Balance()

		if expect != actual {
			t.Errorf("expect %s actual %s", expect, actual)
		}
	}


	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		expect := Bitcoin(10)
		assertBalance(t, wallet, expect)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		expect := Bitcoin(10)
		assertBalance(t, wallet, expect)
	})


}
