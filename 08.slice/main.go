package main

// 数组有一个有趣的属性，它的大小也属于类型的一部分，如果你尝试将 [4]int 作为 [5]int 类型的参数传入函数，
// 是不能通过编译的。它们是不同的类型，就像尝试将 string 当做 int 类型的参数传入函数一样。
func Sum(numbers []int) int {
	sum := 0
	//for i := 0; i < len(numbers); i++ {
	//	sum += numbers[i]
	//}
	// 我们可以使用 range 语法来让函数变得更加整洁。
	// range 会迭代数组，每次迭代都会返回数组元素的索引和值。我们选择使用 _ 空白标志符 来忽略索引。
	for _, number := range numbers {
		sum += number
	}

	return sum
}
