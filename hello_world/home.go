package main

import "fmt"
import "rsc.io/quote"

/* Or:
import (
	"fmt"
	"rsc.io/quote"
)
*/

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
