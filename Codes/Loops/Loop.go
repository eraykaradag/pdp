package main

import "fmt"

func main() {

	// Basic loop structure
	for k := 1; k <= 10; k++ {
		fmt.Print(k, "-")
	}
	fmt.Println()

	// Versions of basic loop structure
	for k := 1; ; k++ {
		fmt.Print(k, "-")
		if k == 10 {
			break
		}
	}
	fmt.Println()

	k := 1
	for ; ; k++ {
		fmt.Print(k, "-")
		if k == 10 {
			break
		}
	}
	fmt.Println()

	k = 1
	for k <= 10 {
		fmt.Print(k, "-")
		k++
	}
	fmt.Println()

	k = 1
	for {
		fmt.Print(k, "-")
		if k == 10 {
			break
		}
		k++
	}
	fmt.Println()

	// The for statement supports one additional form that uses the keyword range to iterate over an expression that evaluates to an array, slice, map, string, or channel

	// Creating an instance of a map data type.
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}

	// Example 1
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}
	fmt.Println()

	// Example 2
	for key := range strDict {
		fmt.Print(key)
	}
	fmt.Println()

	// Example 3
	for _, value := range strDict {
		fmt.Print(value)
	}
	fmt.Println()

}
