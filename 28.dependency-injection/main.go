package main

import (
	"bytes"
	"fmt"
)

func Greet(write *bytes.Buffer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}