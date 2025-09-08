package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func main() {
	// Для примера используется прямоугольный треугольник и пифагоровы тройки: 3, 4, 5.
	point1 := NewPoint(0, 4.0)
	point2 := NewPoint(3.0, 0)

	fmt.Println(point1.Distance(point2)) // 5
}
