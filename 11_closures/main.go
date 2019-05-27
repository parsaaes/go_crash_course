package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	sum := adder()
	fmt.Println(sum(1), sum(1), sum(1))
}
