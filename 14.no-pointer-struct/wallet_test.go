package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	fmt.Println("address of balance in test is", &wallet.balance)

	actual := wallet.Balance()
	expect := 10

	if expect != actual {
		t.Errorf("expect %d actual %d", expect, actual)
	}
}
