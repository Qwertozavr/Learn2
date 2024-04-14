package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	//"time"
)

func Ex7() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex7                  ")
	fmt.Println("---------------------------------------")
	ch := make(chan int)
	defer close(ch)
	data := map[int]int{}
	mu := sync.Mutex{}

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go map_write_ex7(data, &mu, ch, ctx, &wg)
	}

	//time.Sleep(time.Millisecond * 100)

	for i := 0; i < 15; i++ {
		ch <- i
	}

	cancel()

	wg.Wait()
	fmt.Println("=======================================\n")
}

func map_write_ex7(data map[int]int, mu *sync.Mutex, ch chan int, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case input := <-ch:
			mu.Lock()
			data[input] = rand.Intn(100)
			mu.Unlock()
		case <-ctx.Done():
			fmt.Printf("Goroutine has finished\n")
			wg.Done()
			return
		}
	}
}
