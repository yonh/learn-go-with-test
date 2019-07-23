package main

import "fmt"

// Go 允许从现有的类型创建新的类型。语法是 type MyName OriginalType
type Bitcoin int
//
type Stringer interface {
	String() string
}

// 这个接口是在 fmt 包中定义的。当使用 %s 打印格式化的字符串时，你可以定义此类型的打印方式。
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}


type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

