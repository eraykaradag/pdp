package main

import "fmt"

func Function(number int) {
	number += 10
}

func FunctionPointer(number *int) {
	*number += 10
}

func main() {
	n := 10

	Function(n)
	fmt.Println(n)

	FunctionPointer(&n)
	fmt.Println(n)

}
