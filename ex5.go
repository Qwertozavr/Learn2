package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Ex5() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex5                  ")
	fmt.Println("---------------------------------------")
	ch := make(chan int)
	exit := make(chan int)
	defer close(ch)
	defer close(exit)
	var work_time int64

	fmt.Println("Type time duration:")
	fmt.Scan(&work_time)

	time_start := time.Now().Unix()

	go take_ex5(ch, exit)

	for {
		ch <- rand.Intn(100)
		time.Sleep(3 * time.Second)
		if time.Now().Unix()-time_start >= work_time {
			exit <- 1
			break
		}
	}

	time.Sleep(1 * time.Second)
	fmt.Println("=======================================\n")
}

func take_ex5(ch chan int, exit chan int) {
	time.Sleep(100 * time.Millisecond)
	for {
		select {
		case <-ch:
			fmt.Println(<-ch)
		case <-exit:
			return
		}
	}
}
