package main

import (
	"fmt"
	"math"
)

func main() {
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}
	fmt.Println("Rectangle:", getArea(r))
	fmt.Println("Circle:", getArea(c))
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}
