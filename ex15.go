package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {

	var str strings.Builder

	for i := 0; i < size; i++ {
		str.WriteString(".")
	}

	return str.String()
}

func someFunc() string {
	justString := createHugeString(1024)[:100]
	return justString
}

func Ex15() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex15")
	fmt.Println("---------------------------------------")

	fmt.Println(someFunc())

	fmt.Println("=======================================\n")
}
