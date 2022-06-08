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

	// Usage of random and time
	nameWithRandomGreeting("Ekrem")

}

func name(name string) {
	msg := greeter.JustName(name)
	fmt.Println(msg)
}

func nameWithErrors(name string, age string) {
	msgAge, msgAgeError := greeter.WithAge(name, age)
	if msgAgeError != nil {
		// Disabled because it's just an example
		// log.Fatal(msgAgeError)
		log.Print(msgAgeError)
	}
	fmt.Println(msgAge)
}

func nameWithRandomGreeting(name string) {
	msgName, msgNameError := greeter.WithRandom(name)
	if msgNameError != nil {
		log.Fatal(msgNameError)
	}
	fmt.Println(msgName)
}
