package main

import "testing"

func TestSearch(t *testing.T) {

	dict := Dictionary{"key": "hello world"}

	t.Run("known key", func(t *testing.T) {

		expect := "hello world"

		assertSearch(t, dict, "key", expect)
	})

	t.Run("unknown key", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		key := "hello"
		val := "world"

		err := dict.Add(key, val)

		assertError(t, nil, err)
		assertSearch(t, dict, key, val)
	})

	// 如果值已存在，map 不会抛出错误。相反，它们将继续并使用新提供的值覆盖该值。
	// 这在实践中很方便，但会导致我们的函数名称不准确。Add 不应修改现有值。
	// 它应该只在我们的字典中添加新单词。
	t.Run("exists key", func(t *testing.T) {
		key := "hello"
		val := "world"
		dict := Dictionary{key: val}

		err := dict.Add(key, "new value")

		assertError(t, ErrKeyExists, err)
		assertSearch(t, dict, key, val)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("exists key", func(t *testing.T) {
		key := "hello"
		val := "world"
		new_value := "new value"

		dict := Dictionary{key: val}

		err := dict.Update(key, new_value)

		assertError(t, err, nil)
		assertSearch(t, dict, key, new_value)
	})

	t.Run("new key", func(t *testing.T) {

		dict := Dictionary{}

		err := dict.Update("hello", "world")

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func assertError(t *testing.T, expect error, actual error) {

	t.Helper()

	if expect != actual {
		t.Errorf("expect error '%s', actual '%s'", expect, actual)
	}
}

func assertSearch(t *testing.T,dictionary Dictionary, key string, expect string) {
	t.Helper()

	actual, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added key:", err)
	}

	if expect != actual {
		t.Errorf("expect '%s', actual '%s' given", expect, actual)
	}
}
