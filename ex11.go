package main

import (
	"fmt"
)

func Ex11() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex11")
	fmt.Println("---------------------------------------")

	data1 := []int{1, 2, 4, 8, 16}
	data2 := []int{-1, 2, 4, 10}

	flags := make(map[int]bool)

	for _, val := range data1 {
		flags[val] = true
	}

	fmt.Printf("Data 1 = %v\n", data1)
	fmt.Printf("Data 2 = %v\n", data2)

	for _, val := range data2 {
		if flags[val] {
			fmt.Printf("\n%d is in first data.\n", val)
		}
	}

	fmt.Println("=======================================\n")
}
