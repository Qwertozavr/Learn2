package main

import (
	"fmt"
	"sync"
	"time"
)

func Ex2() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex2                  ")
	fmt.Println("---------------------------------------")
	data := [5]int{2, 4, 6, 8, 10}
	go square(data[0])
	go square(data[1])
	go square(data[2])
	go square(data[3])
	go square(data[4])
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------------------------")

	var wg sync.WaitGroup
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(i int) {
			c := data[i]
			defer wg.Done()
			fmt.Println(c, " * ", c, " = ", c*c)
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("=======================================\n")
}

func square(x int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(x, " * ", x, " = ", x*x)
}
