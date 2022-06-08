package main

import (
	"fmt"
	greeter "letsgo/greetings"
	"log"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Print name with parameters
	name("Murat")

	// Multiple parameters and error handling
	nameWithErrors("", "18")

}

func name(name string) {
	msg := greeter.JustName(name)
	fmt.Println(msg)
}

func nameWithErrors(name string, age string) {
	msgAge, msgAgeError := greeter.WithAge(name, age)
	if msgAgeError != nil {
		log.Fatal(msgAgeError)
	}
	fmt.Println(msgAge)
}
