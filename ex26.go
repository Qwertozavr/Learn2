package main

import (
	"fmt"
)

func unic(str string) (string, error) {
	runes := []rune(str)
	dictionary := make(map[rune]int)

	for _, val := range runes {
		if dictionary[val] == 0 {
			dictionary[val] = 1
		} else {
			return "", fmt.Errorf("Lettter %s is not unic in %s", string(val), str)
		}
	}
	return "all latters is unic", nil
}

func Ex26() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex26")
	fmt.Println("---------------------------------------")

	fmt.Println("First answer")
	str := "abcdefghijklmnopqrstuvwxyz"
	result, err := unic(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Orig sring: %s\nResult: %s\n", str, result)

	fmt.Println("\nSecond answer")
	str = "abbccdef"
	result, err = unic(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nOrig sring: %s\nResult: %s", str, result)

	fmt.Println("=======================================\n")
}
