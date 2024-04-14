package main

import (
	"fmt"
	"time"
)

func Ex3() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex3                  ")
	fmt.Println("---------------------------------------")
	data := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	defer close(ch)
	sum := 0

	for _, val := range data {
		go square_ex3(val, ch)
	}

	for i := 0; i < len(data); i++ {
		sum += <-ch
	}
	fmt.Println("Sum = ", sum)

	time.Sleep(1 * time.Second)
	fmt.Println("=======================================\n")
}

func square_ex3(x int, ch chan int) {
	ch <- x * x
}
