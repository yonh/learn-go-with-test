package main

import "testing"

func TestSearch(t *testing.T) {

	dict := Dictionary{"key": "hello world"}

	t.Run("known key", func(t *testing.T) {
		actual, _ := dict.Search("key")
		expect := "hello world"

		assertSearch(t, expect, actual)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func assertError(t *testing.T, expect error, actual error) {
	// 引用类型引入了 maps 可以是 nil 值。如果你尝试使用一个 nil 的 map，你会得到一个 nil 指针异常，这将导致程序终止运行。
	// 由于 nil 指针异常，你永远不应该初始化一个空的 map 变量： var m map[string]string
	// 相反，你可以像我们上面那样初始化空 map，或使用 make 关键字创建 map：
	// dictionary = map[string]string{}
	// dictionary = make(map[string]string)
	// 这两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得 nil 指针异常。
	t.Helper()

	if expect != actual {
		t.Errorf("expect error '%s', actual '%s'", expect, actual)
	}
}

func assertSearch(t *testing.T, expect string, actual string) {
	t.Helper()

	if expect != actual {
		t.Errorf("expect '%s', actual '%s' given, '%s'", expect, actual, "key")
	}
}
