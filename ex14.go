package main

import (
	"fmt"
)

func Ex14() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex14")
	fmt.Println("---------------------------------------")

	var A int
	var B string
	var C bool
	var D chan int

	slc := []interface{}{A, B, C, D}

	for _, value := range slc {
		switch value.(type) {
		case int:
			fmt.Printf("int\n")
		case string:
			fmt.Printf("string\n")
		case bool:
			fmt.Printf("bool\n")
		case chan int:
			fmt.Printf("chan int\n")
		default:
			fmt.Printf("unknown\n")
		}
	}

	fmt.Println("=======================================\n")
}
