package main

import (
	"fmt"
)

func Ex13() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex13")
	fmt.Println("---------------------------------------")

	x := 10
	y := 20

	fmt.Printf("X = %d, Y = %d\n", x, y)

	x = x + y
	y = x - y
	x = x - y

	fmt.Printf("X = %d, Y = %d\n", x, y)

	x, y = y, x

	fmt.Printf("X = %d, Y = %d\n", x, y)

	fmt.Println("=======================================\n")
}
