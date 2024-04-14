package main

import "fmt"

func Ex8() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex8                  ")
	fmt.Println("---------------------------------------")

	var input int64
	var id int
	var zer_or_one int

	fmt.Println("Type number:")
	fmt.Scan(&input)

	fmt.Println("Type bit number:")
	fmt.Scan(&id)

	fmt.Println("Type operation (0-set to 0, 1-set to 1):")
	fmt.Scan(&zer_or_one)
	if zer_or_one == 1 {
		output := input | (1 << id)
		fmt.Println("Number was: ", input, " now ", output)
	} else {
		output := input &^ (1 << id)
		fmt.Println("Number was: ", input, " now ", output)
	}

	fmt.Println("=======================================\n")
}
