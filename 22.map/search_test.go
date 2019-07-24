package main

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"key":"hello world"}

	actual := dict.Search("key")
	expect := "hello world"

	assertSearch(t, expect, actual)
}

func assertSearch(t *testing.T, expect string, actual string) {
	t.Helper()

	if expect != actual {
		t.Errorf("expect '%s' actual '%s' given, '%s'", expect, actual, "key")
	}
}