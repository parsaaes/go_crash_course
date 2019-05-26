package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name + "!"
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Parsa"))
	fmt.Println(sum(1, 2))
}
