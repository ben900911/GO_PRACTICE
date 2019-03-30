package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		height: 5,
		base:   5,
	}

	fmt.Println("tri 5,5:", printArea(tri))

	squ := square{
		sideLength: 10,
	}
	fmt.Println("squ 10,:", printArea(squ))
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(input shape) float64 {
	return input.getArea()
}
