package main

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}
// 声明 map 的方式有点儿类似于数组。不同之处是，它以 map 关键字开头，需要两种类型。第一个是键的类型，写在 [] 中。第二个是值的类型，跟在 [] 之后。
// 键的类型很特别，它只能是一个可比较的类型，因为如果不能判断两个键是否相等，我们就无法确保我们得到的是正确的值。可比类型在语言规范中有详细解释。
// 另一方面，值的类型可以是任意类型，它甚至可以是另一个 map。
type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	return d[key]
}