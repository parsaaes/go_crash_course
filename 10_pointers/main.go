package main

import "fmt"

func double(a *int) {
	*a = 2 * *a
}
func main() {
	a := 10
	b := &a
	fmt.Println(a, *b)
	*b = 4
	fmt.Println(a, *b)
	double(&a)
	fmt.Println(a, *b)
}
