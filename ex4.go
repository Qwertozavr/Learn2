package main

import (
	"fmt"
	"sync"
	"time"
)

func Ex4() {
	fmt.Println("\n=======================================")
	fmt.Println("                  Ex4                  ")
	fmt.Println("---------------------------------------")

	ch := make(chan int)
	defer close(ch)
	var N_workers int
	var wg sync.WaitGroup

	fmt.Println("Type amount of workers:")
	fmt.Scan(&N_workers)

	for i := 1; i <= N_workers; i++ {
		wg.Add(1)
		go worker(i, &wg, ch)
	}

	go func() {
		for {
			var data int
			time.Sleep(time.Millisecond * 100)
			fmt.Print("Enter data: ")
			fmt.Scanf("%d", &data)
			ch <- data
		}
	}()

	wg.Wait()
	fmt.Println("All workers stopped")

	fmt.Println("=======================================\n")
}

func worker(id int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Worker %d received data: %d\n", id, data)
	}
	fmt.Printf("Worker %d stopped\n", id)
}
