package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {

	buffer := bytes.Buffer{}
	Greet(&buffer, "Carrie")

	actual := buffer.String()
	expect := "Hello, Carrie"

	if expect != actual {
		t.Errorf("expect '%s', actual '%s'", expect, actual)
	}

}