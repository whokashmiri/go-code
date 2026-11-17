package main

import (
	"fmt"
	// "net/http"
	"sync"
)

var wg sync.WaitGroup

// func greeter(s string) {
// 	for i := 0; i < 4; i++ {

// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println("before time")
// 		fmt.Println("Data camp", s)
// 	}
// }
// func main() {
// 	go greeter("ehl")
// 	greeter("bomf")
// }

func main() {
	fmt.Println("Second Lecture Goroutines and Parallelism")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	mutR := &sync.RWMutex{}

	var scr = []int{0}
	wg.Add(3)
	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One r")
		mut.Lock()
		scr = append(scr, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(w *sync.WaitGroup) {
		fmt.Println("Two r")
		scr = append(scr, 2)
		wg.Done()
	}(wg)
	go func(w *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three r")
		mutR.RLock()
		scr = append(scr, 3)
		mutR.RUnlock()
		wg.Done()
	}(wg, mutR)

	wg.Wait()
	fmt.Println(scr)
	// websiteList := []string{
	// 	"https://go.dev",
	// 	"https://ai.com",
	// 	"https://fb.com",
	// }

	// for _, web := range websiteList {
	// 	go getStatusCode(web)
	// 	wg.Add(1)

	// }
	// wg.Wait()
}

// func getStatusCode(endpoint string) {
// 	defer wg.Done()
// 	res, err := http.Get(endpoint)

// 	if err != nil {
// 		fmt.Println("error in END")
// 	} else {
// 		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
// 	}
// }
