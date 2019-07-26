package main

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url:= range urls {

		// 开启 goroutine 的唯一方法就是将 go 放在函数调用前面
		// 匿名函数有许多有用的特性，其中之一它们可以在声明的同时执行 —— 这就是匿名函数末尾的 () 实现的


		// 这里会导致results返回的结果只有一个结果
		// - 为什么只有一个结果？
		// 这里的问题是变量 url 被重复用于 for 循环的每次迭代
		// —— 每次都会从 urls 获取新值。但是我们的每个 goroutine 都是 url 变量的引用
		// —— 它们没有自己的独立副本。所以他们都会写入在迭代结束时的 url —— 最后一个 url。
		// 这就是为什么我们得到的结果是最后一个 url。
		//go func() {
		//	results[url] = wc(url)
		//}()


		// 不幸运的话，这里还是会得到错误结果 fatal error: concurrent map writes,
		// 这是由于两个 goroutines 完全同时写入 results map。Go 的 Maps 不喜欢多个事物试图一次性写入，所以就导致了 fatal error。
		// 这是一种 race condition（竞争条件），当软件的输出取决于事件发生的时间和顺序时，因为我们无法控制，bug 就会出现。
		// 因为我们无法准确控制每个 goroutine 写入结果 map 的时间，两个 goroutines 同一时间写入时程序将非常脆弱。
		// Go 可以帮助我们通过其内置的 race detector 来发现竞争条件。要启用此功能，请使用 race 标志运行测试：go test -race。
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(1 * time.Second)

	return results
}
// go test
// go test -race