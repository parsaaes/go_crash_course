package main

import "fmt"

func main() {
	grades := make(map[string]float64)
	grades["Ali"] = 20
	grades["Mahmood"] = 16.75
	grades["Sina"] = 18.5
	grades2 := map[string]float64{"Gholi": 10, "Pooya": 13.5}
	grades2["Mehdi"] = 18.5
	fmt.Println(grades)
	fmt.Println("Ali's grade:", grades["Ali"])
	delete(grades, "Mahmood")
	fmt.Println(grades)
	fmt.Println(grades2)
}
