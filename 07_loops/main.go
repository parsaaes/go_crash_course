package main

import "fmt"

func main() {
	i := 2
	for i > 0 {
		fmt.Println(i)
		i--
	}
	// FizzBuzz
	fmt.Println("FizzBuzz:")
	for i := 1; i < 30; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
