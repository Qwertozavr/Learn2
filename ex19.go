package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Ex19() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex19")
	fmt.Println("---------------------------------------")

	str := "главрыба"
	str = reverse(str)
	fmt.Printf("First reverse: %s\n", str)

	var str2 strings.Builder
	runes := []rune(str)

	for i := len(runes) - 1; i > 0; i-- {
		str2.WriteRune(runes[i])
	}
	fmt.Printf("Second reverse: %s\n", str2.String())

	fmt.Println("=======================================\n")
}
