package main

import (
	"fmt"
	"time"
)

func greeter(s string) {
	for i := 0; i < 4; i++ {

		time.Sleep(3 * time.Millisecond)
		fmt.Println("before time")
		fmt.Println("Data camp", s)
	}
}
func main() {
	go greeter("ehl")
	greeter("bomf")
}
