package main

import "fmt"

func main() {
	defer fmt.Println("Defer here")
	fmt.Println(" defer")
	defer fmt.Println("LAst here")
	myDef()

}

func myDef() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
