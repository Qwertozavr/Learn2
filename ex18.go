package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
}

func Ex18() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex18")
	fmt.Println("---------------------------------------")

	c1 := Counter{count: 0}
	mu := sync.Mutex{}

	go func(c1 *Counter, mu *sync.Mutex) {
		for i := 0; i < 5; i++ {
			mu.Lock()
			c1.count++
			mu.Unlock()
		}
	}(&c1, &mu)

	time.Sleep(time.Millisecond * 1)

	for i := 0; i < 5; i++ {
		mu.Lock()
		c1.count++
		mu.Unlock()
	}

	fmt.Printf("Result count is %d\n", c1.count)

	fmt.Println("=======================================\n")
}
