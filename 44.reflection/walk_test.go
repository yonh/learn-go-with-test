package main

import "testing"

func TestWalk(t *testing.T) {
	expect := "Carrie"
	var actual []string

	x := struct {
		Name string
	}{expect}

	walk(x, func(input string) {
		actual = append(actual, input)
	})

	if len(actual)!= 1 {
		t.Errorf("wrong number of function calls, expect %d, actual %d", 1, len(actual))
	}

	if actual[0] != expect {
		t.Errorf("expect '%s', actual '%s'", expect, actual[0])
	}

}
