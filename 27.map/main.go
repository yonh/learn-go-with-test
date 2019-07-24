package main

// 我们将错误声明为常量，这需要我们创建自己的 DictionaryErr 类型来实现 error 接口。
// 你可以在 Dave Cheney 的这篇优秀文章中了解更多相关的细节。简而言之，它使错误更具可重用性和不可变性。(https://dave.cheney.net/2016/04/07/constant-errors)
const (
	ErrNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrKeyExists       = DictionaryErr("cannot add value because it already exists")
	ErrKeyDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

// 声明 map 的方式有点儿类似于数组。不同之处是，它以 map 关键字开头，需要两种类型。第一个是键的类型，写在 [] 中。第二个是值的类型，跟在 [] 之后。
// 键的类型很特别，它只能是一个可比较的类型，因为如果不能判断两个键是否相等，我们就无法确保我们得到的是正确的值。可比类型在语言规范中有详细解释。
// 另一方面，值的类型可以是任意类型，它甚至可以是另一个 map。
type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {

	// 我们使用了一个 map 查找的有趣特性。它可以返回两个值。第二个值是一个布尔值，表示是否成功找到 key。
	value, found := d[key]
	if !found {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) {
	// Go 的 map 有一个内置函数 delete。它需要两个参数。第一个是这个 map，第二个是要删除的键。
	delete(d, key)

	// delete 函数不返回任何内容，我们基于相同的概念构建 Delete 方法。由于删除一个不存在的值是没有影响的，与我们的 Update 和 Create 方法不同，我们不需要用错误复杂化 API。
}

// Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。
// 这意味着它拥有对底层数据结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map，
// 你可以在这里阅读有关 hash tables 的更多信息(https://en.wikipedia.org/wiki/Hash_table)。
func (d Dictionary) Add(key string, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrKeyExists
	default:
		return err
	}
}
