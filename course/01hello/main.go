package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main")

	presentTime := time.Now()
	asc := time.ANSIC
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Printf("%s", asc)

}
