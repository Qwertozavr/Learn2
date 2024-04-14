package main

import (
	"fmt"
)

func Ex12() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex12")
	fmt.Println("---------------------------------------")

	data := []string{"cat", "cat", "dog", "cat", "tree"}

	grouped_data := make(map[string]int)

	for _, val := range data {
		grouped_data[val] += 1
	}

	fmt.Println(grouped_data)

	fmt.Println("=======================================\n")
}
