package main

import "fmt"

func main() {
	x, y := 10, 20
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "Red"
	if color == "Green" {
		fmt.Println(color + " is Green :|")
	} else if color == "Red" {
		fmt.Println("Color is Red")
	}

	switch color {
	case "Red":
		fmt.Println("Red")
	case "Blue":
		fmt.Println("Blue")
	default:
		fmt.Println("This is default")
	}
}
