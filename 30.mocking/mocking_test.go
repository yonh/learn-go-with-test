package main

import (
	"bytes"
	"testing"
)

func TestGountdown(t *testing.T) {

	buffer := &bytes.Buffer{}

	Countdown(buffer)

	actual := buffer.String()
	expect := "3"

	if expect != actual {
		t.Errorf("expect '%s', actual '%s'", expect, actual)
	}

}
