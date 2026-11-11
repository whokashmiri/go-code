package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[1] = "we"
	fruitList[2] = "re"

	fmt.Println(fruitList)
	fmt.Printf("%d" ,len(fruitList))

	var vegList [4]string
	vegList[1] = "we"
	vegList[2] = "re"

	fmt.Println(vegList)
	fmt.Printf("%d" ,len(vegList))
	
}