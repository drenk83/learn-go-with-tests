package main

import "fmt"

const (
	english = "English"
	spanish = "Spanish"
	french  = "French"
)

var helloPrefixes = map[string]string{
	english: "Hello, ",
	french: "Bonjour, ",
	spanish: "Hola, ",
}

func Hello(name, language string) string {
	prefix := helloPrefixes[english]

	if name == "" {
		name = "world"
	}

	if p, ok := helloPrefixes[language]; ok {
		prefix = p
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
