package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {

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
		//go func(u string) {
		//	results[u] = wc(u)
		//}(url)

		// 我们可以通过使用 channels 协调我们的 goroutines 来解决这个数据竞争。
		// channels 是一个 Go 数据结构，可以同时接收和发送值。这些操作以及细节允许不同进程之间的通信。
		go func(u string) {
			// 现在，当我们迭代 urls 时，不是直接写入 map，而是使用 send statement 将每个调用 wc 的 result 结构体发送到 resultChannel。
			// 这使用 <- 操作符，channel 放在左边，值放在右边：
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i :=0; i < len(urls); i++ {
		// 我们在内部使用 receive expression，它将从通道接收到的值分配给变量。
		// 这也使用 <- 操作符，但现在两个操作数颠倒过来：现在 channel 在右边，我们指定的变量在左边：
		result := <-resultChannel
		// 然后我们使用接收到的 result 更新 map。
		results[result.string] = result.bool
	}
	// 通过将结果发送到通道，我们可以控制每次写入 results map 的时间，确保每次写入一个结果。虽然 wc 的每个调用都发送给结果通道，
	// 但是它们在其自己的进程内并行发生，因为我们将结果通道中的值与接收表达式一起逐个处理一个结果。
	// 我们已经将想要加快速度的那部分代码并行化，同时确保不能并发的部分仍然是线性处理。我们使用 channel 在多个进程间通信。

	return results
}

// go test
// go test -race
