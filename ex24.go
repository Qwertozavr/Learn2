package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func newPoint() Point {
	return Point{0, 0}
}

func distance(point1, point2 *Point) float64 {
	return math.Sqrt(float64((point1.x-point2.x)*(point1.x-point2.x) + (point1.y-point2.y)*(point1.y-point2.y)))
}

func Ex24() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex24")
	fmt.Println("---------------------------------------")

	point1 := newPoint()
	point2 := newPoint()

	point1.x = 1
	point1.y = 1
	point2.x = 4
	point2.y = 5

	result := distance(&point1, &point2)
	fmt.Printf("Distance: %0.2f\n", result)

	fmt.Println("=======================================\n")
}
