package main

import (
	"fmt"
	"time"
)

func Ex9() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex9                  ")
	fmt.Println("---------------------------------------")

	ch1 := make(chan int)
	ch2 := make(chan int)
	exit := make(chan int)
	defer close(ch1)
	defer close(ch2)
	defer close(exit)

	data := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	go recive_and_send_ex9(data, ch1)

	go recive_and_send_two_ex9(ch1, ch2, exit)

	for i := 0; i < 10; i++ {
		fmt.Println("Received and doubled: ", <-ch2)
	}
	time.Sleep(100 * time.Millisecond)
	exit <- 0

	fmt.Println("=======================================\n")
}

func recive_and_send_ex9(data [10]int, ch1 chan int) {
	for _, val := range data {
		ch1 <- val
	}
}

func recive_and_send_two_ex9(ch1, ch2, exit chan int) {
	for {
		select {
		case recive := <-ch1:
			ch2 <- (recive * 2)
		case <-exit:
			return
		}
	}
}
