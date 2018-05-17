package main

import "fmt"

func main() {
	fmt.Println(Hello("Myself", ""))
}

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// Named return. a prefix variable is created can be used inside the function.
//MUY IMPORTANTE.
//En Go, las funciones privadas empiezan en minúsculas, y las publicas en mayúsculas.
func greetingPrefix (language string) (prefix string){
    switch language {
    case spanish:
        prefix=spanishHelloPrefix
    case french:
        prefix=frenchHelloPrefix
    default:
        prefix=helloPrefix
    }
    //return will send anything inside prefix variable.
    return
}
