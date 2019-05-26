package main

import "fmt"

func main() {
	var fruitArr [5]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fruitArr[2] = "Banana"
	numArr := [3]int{1, 2, 3}
	numSlice := []int{1, 2, 3}
	fmt.Println(fruitArr, numArr, numSlice)
	fmt.Println(len(numSlice))
	fmt.Println(numSlice[1:2])
}
