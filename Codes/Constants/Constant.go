package main

import "fmt"

// We can declare constants by const keyword
const PRODUCT string = "Canada"
const PRICE = 500

// We must assign the value at the time of the constant declaration
// const PI

// Constants declaration can to be grouped together into blocks for greater readability and code quality
const (
	HEALTH = 100
	SPEED  = 30
	WEIGHT = 80
)

// We don't have to name constants with uppercase letters but it's a good-practice

func main() {
	// Constant declaration can be done in local area
	const AGE = 40
	const NAME = "Name"

	fmt.Printf("%s %d %d", PRODUCT, HEALTH, AGE)
}
