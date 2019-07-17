package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return helloPrefix
	}
}

func main() {
	fmt.Printf(Hello("",""))
}