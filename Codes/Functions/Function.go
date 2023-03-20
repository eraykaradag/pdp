package main

import (
	"fmt"
)

// Basic function declaration
func Greeting(name string) {
	fmt.Println("Hello", name)
}

// Specify the return type
func Addition(x int, y int) int {
	return x + y
}

// And if we give a name to the return variable type
// It will create a new variable and at the end of the function
// it will return that variabele if I don't return anything
func Multiplication(x int, y int) (Area int) {
	Area = x * y
	return
}

// we can return more than one value
func Subtraction(x int, y int) (int, string) {
	return x - y, " Subtraction"
}

// we can get another function as a parameter
func Statistic(f1 func(int, int) int, f2 func(int, int) (int, string)) {
	var a = f1(5, 3)
	var b, operation = f2(5, 3)

	fmt.Print(a, b, operation)
}

// Anonymous functions can accept inputs and return outputs, just as standard functions do.
var (
	area = func(x int, y int) int {
		return x * y
	}
)

func main() {

	// Greeting("Murat")

	// fmt.Println(Addition(5, 10))

	// fmt.Println(Multiplication(5, 10))

	// fmt.Println(Subtraction(20, 14))

	// Statistic(Addition, Subtraction)

	// fmt.Println(area(3, 4))

	// Here we declare an anonymous function
	// Send 50 and 5 as a parameter
	// Expect to return 10 as a result of operation inside curly brackets
	// var result int = func(x int, y int) int { return x / y }(50, 5)
	// fmt.Println(result)

}
