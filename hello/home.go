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

	// Usage of math and for loop
	greetingPerName(true)

}

func name(name string) {
	msg, err := greeter.JustName(name)
	if err != nil {
		log.Print(err)
	}
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
	fmt.Println("")
}

func greetingPerName(forEach bool) {
	// Names
	names := []string{
		"Gladys", "Samantha", "Darrin",
	}
	messages, err := greeter.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of
	// messages to the console.
	if forEach != true {
		fmt.Println(messages)
	} else {
		for key, msg := range messages {
			fmt.Println(key + ": " + msg)
		}
	}
}
