package main

import (
	"fmt"
	firstModule "letsgo/first_module"
	"log"
	"rsc.io/quote"
)

/* Or:
import "fmt"
import "rsc.io/quote"
*/

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	msg := firstModule.Hello("Murat")
	fmt.Println(msg)

	msgAge, msgAgeError := firstModule.HelloAge("", "18")
	if msgAgeError != nil {
		log.Fatal(msgAgeError)
	}

	fmt.Println(msgAge)
}
