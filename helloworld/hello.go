package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello -> returns Hello world
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := getPrefix(lang)
	return fmt.Sprintf("%s%s", prefix, name)
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
