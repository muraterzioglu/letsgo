package main

import (
	"fmt"
	firstModule "letsgo/first_module"
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

	msgAge := firstModule.HelloAge("Murat", "18")
	fmt.Println(msgAge)
}
