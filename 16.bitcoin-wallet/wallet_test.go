package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	actual := wallet.Balance()
	expect := Bitcoin(10)

	if expect != actual {
		t.Errorf("expect %d actual %d", expect, actual)
	}
}
