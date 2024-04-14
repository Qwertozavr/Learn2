package main

import (
	"fmt"
	"sort"
)

func Ex16() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex16")
	fmt.Println("---------------------------------------")

	data := []int{10, 9, 5, 6, 8, 7, 3, 4, 2, 1}

	sort.Slice(data, func(i, j int) bool { return data[j] > data[i] })
	fmt.Printf("Sorted array: %v\n", data)

	fmt.Println("=======================================\n")
}
