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

func SumAll(numbers ...[]int) []int {

	// 这里有一种创建切片的新方式。make 可以在创建切片的时候指定我们需要的长度和容量。
	//sums = make([]int, lengthOfNumbers)

	// 在这个实现中，我们不用担心切片元素会超过容量。我们开始使用空切片（在函数签名中定义），在每次计算完切片的总和后将结果添加到切片中。
	var sums []int

	for _, nums := range numbers {
		sums = append(sums, Sum(nums))
	}

	return sums
}
