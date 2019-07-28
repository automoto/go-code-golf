package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height float64
	base float64
}

func main() {
	newSquare := square{ 20 }
	newTriangle := triangle{ height: 10, base: 3}

	printArea(newSquare)
	printArea(newTriangle)
}

func (s square) getArea() float64{
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64{
	return .5 * t.base * t.height
}

func printArea(s shape){
	fmt.Println("Area is: ", s.getArea())
}
