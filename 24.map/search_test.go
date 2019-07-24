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

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	key := "hello"
	val := "world"

	dict.Add(key, val)

	assertAdd(t, dict, key, val)
}

func assertAdd(t *testing.T, dictionary Dictionary, key string, value string) {
	t.Helper()

	actual, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added key:", err)
	}

	if actual != value {
		t.Errorf("expect '%s', actual '%s' given", value, actual)
	}
}


func assertError(t *testing.T, expect error, actual error) {

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
