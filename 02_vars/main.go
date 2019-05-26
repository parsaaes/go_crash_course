package main

import "fmt"

var globalVar = 1

func main() {
	var a string = "ABC"
	var b = 123
	c := "XYZ"
	d := 1.32
	e, f := 34.221, "PQR"
	const isTrue = true
	fmt.Println(a, b, globalVar, c, d, e, f)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", isTrue)
}
