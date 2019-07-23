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

	assertError := func(t *testing.T, err error, expect string) {
		if err == nil {
			// t.Fatal。如果它被调用，它将停止测试。这是因为我们不希望对返回的错误进行更多断言。如果没有这个，测试将继续进行下一步，并且因为一个空指针而引起 panic。
			t.Fatal("wanted an error but didnt get one")
		}

		if err.Error() != expect {
			t.Errorf("got '%s', want '%s'", err.Error(), expect)
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

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		// 如果你尝试取出超过你余额的比特币，我们想让 Withdraw 返回一个错误，而余额应该保持不变。
		expect := Bitcoin(10)
		assertBalance(t, wallet, expect)

		assertError(t, err, "cannot withdraw, insufficient funds")
	})


}
