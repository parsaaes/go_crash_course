package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	width float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (square Square) area() float64 {
	return square.width * square.width
}

func getArea(shape Shape) float64 {
	return shape.area()
}
func main() {
	circle := Circle{1}
	square := Square{2}
	fmt.Println(getArea(circle))
	fmt.Println(getArea(square))
}
