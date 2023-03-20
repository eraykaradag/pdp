package main

import "fmt"

func main() {

	x := 60
	fmt.Println("X = ", x)
	// Basic Condition if - else statements
	if x > 50 {
		fmt.Println("Greater than 50")
	} else {
		fmt.Println("Less than 50")
	}

	// Basic if-else if-else statement
	if x > 75 {
		fmt.Println("Greater than 75")
	} else if x > 50 {
		fmt.Println("Greater than 50 but less than 75")
	} else {
		fmt.Println("Less than 50")
	}

	// If statement initialization
	x, y := 3, 4

	if resultM := x * y; resultM > 20 {
		fmt.Println("Result is greater than 20")
	} else if resultD := y / x; resultM > 0 {
		fmt.Printf("Multiplication is %d \nDivition is %d", resultM, resultD)
	} else {
		fmt.Println("Result is less than 0")
	}

}
