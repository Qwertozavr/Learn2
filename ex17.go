package main

import (
	"fmt"
	"sort"
)

func Ex17() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex17")
	fmt.Println("---------------------------------------")

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	index := sort.Search(len(data), func(i int) bool { return data[i] >= 5 })

	if index == -1 {
		fmt.Printf("7 not found\n")
	} else {
		fmt.Printf("Index of 5 is %d\n", index)
	}

	fmt.Println("=======================================\n")
}
