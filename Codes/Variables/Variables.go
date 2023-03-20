package main

import (
	"fmt"
)

// We can declare variables on the package level
var gloabal int = 10
var gloabal2 = 10

// Important thing is that short declaration operator can’t be used at package level it must be used inside of the function.
// gloabal3 := 10

func main() {

	// Basic variable declarations
	var a string = "Basic variable declaration"
	var b = "We can also omit the type of varaible. In this case , it will be inferred"
	c := "We don't even need to use var keyword. We can declare variables with :=(short declaration) operator."

	// We can declare multiple variables at the same times
	var d, e int = 10, 20
	var f, g = 10, 20
	h, i := 10, 20

	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g, h, i)

}
