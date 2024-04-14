package main

import (
	"fmt"
	"runtime"
	"time"
)

func Ex6() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex6                 ")
	fmt.Println("---------------------------------------")
	exit := make(chan int)

	go goroutine_1_ex6(exit)
	time.Sleep(time.Microsecond * 100)
	exit <- 0

	go func() {
		time.Sleep(time.Microsecond * 100)
		defer fmt.Println("Goroutine 2 exiting")

		runtime.Goexit()
	}()

	go func() {
		time.Sleep(time.Microsecond * 200)
		defer fmt.Println("Goroutine 3 exiting")

		panic("Произошла ошибка")
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("=======================================\n")
}

func goroutine_1_ex6(exit chan int) {
	select {
	case <-exit:
		fmt.Println("Goroutine 1 exiting")
		return
	}
}
