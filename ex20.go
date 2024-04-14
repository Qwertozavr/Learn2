package main

import (
	"fmt"
	"strings"
)

func Ex20() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex20")
	fmt.Println("---------------------------------------")

	str := "snow dog sun"

	result := strings.Split(str, " ")

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	out := strings.Join(result, " ")

	fmt.Printf("Result: %s\n", out)

	fmt.Println("=======================================\n")
}
