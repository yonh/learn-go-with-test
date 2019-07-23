package main

import "fmt"

// Go 允许从现有的类型创建新的类型。语法是 type MyName OriginalType
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
