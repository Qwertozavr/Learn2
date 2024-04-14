package main

import (
	"fmt"
	"time"
)

func sleep_ex25(sec int) {
	start := time.Now()

	for {
		if time.Since(start).Seconds() > float64(sec) {
			fmt.Printf("\nTime is out!\n")
			break
		}
	}
}

func Ex25() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex25")
	fmt.Println("---------------------------------------")

	fmt.Printf("Type sleep duration(seconds): ")
	var seconds int
	fmt.Scan(&seconds)

	sleep_ex25(seconds)

	fmt.Println("=======================================\n")
}
