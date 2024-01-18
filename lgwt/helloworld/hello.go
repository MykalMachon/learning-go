package helloworld

import "fmt"

// setup a bunch of const's
const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// the (prefix string)
// defines a return varialbe called prefix of type string
func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getGreetingPrefix(language) + name + "!"
}

func main() {
	fmt.Println(Hello("Mykal", ""))
}
