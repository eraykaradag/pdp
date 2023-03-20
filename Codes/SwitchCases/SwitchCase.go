package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	// fmt.Println(today.Day())
	var day int = today.Day() / 5

	switch day {
	case 0:
		fmt.Println("1/6 of Month")
	case 1:
		fmt.Println("2/6 of Month")
	case 2:
		fmt.Println("3/6 of Month")
	case 3:
		fmt.Println("4/6 of Month")
	case 4:
		fmt.Println("5/6 of Month")
	case 5:
		fmt.Println("6/6 of Month")
	default:
		fmt.Println("There is a problem")
	}
	fmt.Println()

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Clean your house.")
	case 6, 7, 8, 9, 10:
		fmt.Println("Buy some food.")
	case 11, 12, 13, 14:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
	fmt.Println()

	index := 2
	switch index {
	case 1:
		fmt.Println("Clean your house.")
		fallthrough
	case 2:
		fmt.Println("Buy some fruit.")
		fallthrough
	case 3:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 4:
		fmt.Println("Buy some food.")
		fallthrough
	case 5:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
	fmt.Println()

	//

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some fruit.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
	fmt.Println()

}
