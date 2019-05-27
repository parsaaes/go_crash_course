package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 8, 9, 2, 3, 21}
	for i, item := range arr {
		fmt.Printf("arr[%d]= %d\n", i, item)
	}
}
