package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO!")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
