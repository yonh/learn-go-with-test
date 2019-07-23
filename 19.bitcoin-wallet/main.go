package main

import (
	"errors"
	"fmt"
)

// 指针
// 当你传值给函数或方法时，Go 会复制这些值。因此，如果你写的函数需要更改状态，你就需要用指针指向你想要更改的值
// nil
// 当函数返回一个的指针，你需要确保检查过它是否为 nil，否则你可能会抛出一个执行异常，编译器在这里不能帮到你

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

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance >= amount {
		w.balance -= amount
		return nil
	}

	return errors.New("cannot withdraw, insufficient funds")

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

