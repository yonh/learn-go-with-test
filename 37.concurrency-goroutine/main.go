package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url:= range urls {

		// 开启 goroutine 的唯一方法就是将 go 放在函数调用前面
		// 匿名函数有许多有用的特性，其中之一它们可以在声明的同时执行 —— 这就是匿名函数末尾的 () 实现的
		go func() {
			results[url] = wc(url)
		}()

		// 上面匿名函数的主体和之前循环体中的完全一样。唯一的区别是循环的每次迭代都会启动一个新的 goroutine，
		// 与当前进程（WebsiteChecker 函数）同时发生，每个循环都会将结果添加到 results map 中。
	}

	return results
}