package main 

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, " 
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(args ...string) string {
	name := args[0]
	prefix := englishHelloPrefix

	if name == "" {
		name = "world"
	}

	if len(args) > 1 {
		prefix = choosePrefix(args[1])
	}

	return prefix + name
}

func choosePrefix(lang string) (prefix string) {
	switch lang {
	case spanish: prefix = spanishHelloPrefix
	case french: prefix = frenchHelloPrefix
	default: prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world"))
}