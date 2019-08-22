package main

import "fmt"
const englishHelloPrefix = "Hello, "
const thaiHelloPrefix = "สวัสดี "
const frenchHelloPrefix = "Bonjour, "
const thai = "thai"
const french = "french"
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
		case french:
			prefix = frenchHelloPrefix
		case thai:
			prefix = thaiHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "english"))
}